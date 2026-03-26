package nslaslicense_offline

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"terraform-provider-netscalersdx/internal/service"
	lasutils "terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*nslaslicenseOfflineResource)(nil)
var _ resource.ResourceWithConfigure = (*nslaslicenseOfflineResource)(nil)
var _ resource.ResourceWithImportState = (*nslaslicenseOfflineResource)(nil)

func NslaslicenseOfflineResource() resource.Resource {
	return &nslaslicenseOfflineResource{}
}

type nslaslicenseOfflineResource struct {
	client *service.NitroClient
}

// LASSecretsModel describes the LAS secrets JSON structure
type LASSecretsModel struct {
	Ccid        string `json:"ccid"`
	Client      string `json:"client"`
	Password    string `json:"password"`
	LasEndpoint string `json:"las_endpoint"`
	CcEndpoint  string `json:"cc_endpoint"`
}

func (r *nslaslicenseOfflineResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *nslaslicenseOfflineResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslaslicense_offline"
}

// Configure configures the client resource.
func (r *nslaslicenseOfflineResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *nslaslicenseOfflineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nslaslicenseOfflineResourceSchema(ctx)
}

func (r *nslaslicenseOfflineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of nslaslicense_offline Resource")

	var data nslaslicenseOfflineModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate provider configuration
	if r.client == nil {
		resp.Diagnostics.AddError(
			"Provider Not Configured",
			"The provider must be configured with ccid, client, password, las_endpoint, and cc_endpoint",
		)
		return
	}

	// Get provider configuration from NitroClient
	endpoint := r.client.Host()
	username := r.client.Username()
	password := r.client.Password()

	if endpoint == "" || username == "" || password == "" {
		resp.Diagnostics.AddError(
			"Provider Configuration Incomplete",
			"All provider fields (endpoint, username, password) must be configured",
		)
		return
	}

	// Extract device IP from provider endpoint
	deviceIP := extractIPFromEndpoint(endpoint)

	// Validate username
	if username != "nsroot" {
		resp.Diagnostics.AddError(
			"Invalid Username",
			"Username must be 'nsroot'",
		)
		return
	}

	// Read LAS secrets from file
	lasSecretsPath := data.LASSecretsJson.ValueString()
	lasSecretsData, err := os.ReadFile(lasSecretsPath)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Read LAS Secrets File",
			fmt.Sprintf("Cannot read file %s: %s", lasSecretsPath, err.Error()),
		)
		return
	}

	var lasSecrets LASSecretsModel
	if err := json.Unmarshal(lasSecretsData, &lasSecrets); err != nil {
		resp.Diagnostics.AddError(
			"Failed to Parse LAS Secrets JSON",
			fmt.Sprintf("Invalid JSON in file %s: %s", lasSecretsPath, err.Error()),
		)
		return
	}

	// Validate LAS secrets
	if lasSecrets.Ccid == "" || lasSecrets.Client == "" || lasSecrets.Password == "" ||
		lasSecrets.LasEndpoint == "" || lasSecrets.CcEndpoint == "" {
		resp.Diagnostics.AddError(
			"Incomplete LAS Secrets",
			"LAS secrets JSON must contain: ccid, client, password, las_endpoint, cc_endpoint",
		)
		return
	}

	// // Validate PEM code for SDX - must start with CNS_M
	// requestPEM := data.RequestPEM.ValueString()
	// if !strings.HasPrefix(requestPEM, "CNS_M") {
	// 	resp.Diagnostics.AddError(
	// 		"Invalid PEM Code for SDX",
	// 		fmt.Sprintf("SDX PEM codes must start with 'CNS_M' (e.g., CNS_M8920_SERVER, CNS_M15120_SERVER). Got: %s", requestPEM),
	// 	)
	// 	return
	// }

	// // Validate edition for SDX based on PEM code
	// requestED := data.RequestED.ValueString()
	// if strings.Contains(requestPEM, "CNS_M15") {
	// 	// M15xxx models support "50G" edition
	// 	if requestED != "50G" && requestED != "Premium" {
	// 		resp.Diagnostics.AddError(
	// 			"Invalid Edition for SDX M15xxx",
	// 			fmt.Sprintf("SDX M15xxx models support '50G' or 'Premium' editions. Got: %s", requestED),
	// 		)
	// 		return
	// 	}
	// } else if strings.Contains(requestPEM, "CNS_M26") {
	// 	// M26xxx models support "50S", "100G", or "Premium" editions
	// 	if requestED != "50S" && requestED != "100G" && requestED != "Premium" {
	// 		resp.Diagnostics.AddError(
	// 			"Invalid Edition for SDX M26xxx",
	// 			fmt.Sprintf("SDX M26xxx models support '50S', '100G', or 'Premium' editions. Got: %s", requestED),
	// 		)
	// 		return
	// 	}
	// } else {
	// 	// Standard SDX models only support "Premium" edition
	// 	if requestED != "Premium" {
	// 		resp.Diagnostics.AddError(
	// 			"Invalid Edition for SDX",
	// 			fmt.Sprintf("Standard SDX models only support 'Premium' edition. For special bandwidth models, use M15xxx (50G) or M26xxx (50S/100G). Got: %s", requestED),
	// 		)
	// 		return
	// 	}
	// }

	// Read and validate entitlement_name from config
	entitlementName := data.EntitlementName.ValueString()
	validSDXPrefixes := []string{
		"SDX 89", "SDX 91", "SDX 92", "SDX 14",
		"SDX 15", "SDX 16", "SDX 17", "SDX 26",
	}
	validPrefix := false
	for _, prefix := range validSDXPrefixes {
		if strings.HasPrefix(entitlementName, prefix) {
			validPrefix = true
			break
		}
	}
	if !validPrefix {
		resp.Diagnostics.AddError(
			"Invalid Entitlement Name",
			fmt.Sprintf("entitlement_name must start with a valid SDX model prefix (%s). Got: %s", strings.Join(validSDXPrefixes, ", "), entitlementName),
		)
		return
	}

	product := "SDX" // This resource is for SDX only

	tflog.Info(ctx, "Starting offline LAS license generation for SDX", map[string]interface{}{
		"device_ip": deviceIP,
	})

	hostname := "sdx" // Use default hostname for SDX

	// Step 1: Get SDX version and determine API selection
	release, build, err := lasutils.GetMPSVersion(ctx, deviceIP, username, password)
	if err != nil {
		tflog.Warn(ctx, "Failed to get SDX version, using default API", map[string]interface{}{"error": err.Error()})
	} else {
		// Set version and build in state
		data.Version = types.StringValue(release)
		data.Build = types.StringValue(build)
	}

	useNewAPI := false
	if release != "" && build != "" {
		useNewAPI = lasutils.DetermineNewAPINeeded(product, release, build)
		tflog.Info(ctx, "API selection for SDX", map[string]interface{}{
			"release":   release,
			"build":     build,
			"useNewAPI": useNewAPI,
		})
	}

	// Step 2: Generate offline request package
	filename, packageData, err := lasutils.GetOfflineRequestPackage(ctx, product, deviceIP, hostname, username, password, useNewAPI)
	if err != nil {
		resp.Diagnostics.AddError(
			"Request Package Generation Failed",
			fmt.Sprintf("Failed to generate offline request package: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Generated request package", map[string]interface{}{"filename": filename})

	// Step 3: Extract LSGUID from package
	lsguid, err := lasutils.ExtractLSGUIDFromPackage(ctx, product, packageData)
	if err != nil {
		resp.Diagnostics.AddError(
			"LSGUID Extraction Failed",
			fmt.Sprintf("Failed to extract LSGUID from request package: %s", err.Error()),
		)
		return
	}
	data.LSGUID = types.StringValue(lsguid)

	tflog.Info(ctx, "Extracted LSGUID", map[string]interface{}{"lsguid": lsguid})

	// // Step 4: Determine entitlement name for SDX
	// entitlementName, err := lasutils.GetEntitlementNameForFixedBW(product, data.RequestPEM.ValueString(), data.RequestED.ValueString(), false)
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Invalid Entitlement Configuration",
	// 		fmt.Sprintf("Failed to determine entitlement name: %s", err.Error()),
	// 	)
	// 	return
	// }
	// tflog.Info(ctx, "Determined entitlement", map[string]interface{}{"entitlementName": entitlementName})

	// Step 5: Initialize LAS Token Generator
	ltg := lasutils.NewLASTokenGenerator(
		"netscalerfixedbw",
		lsguid,
		lasSecrets.Ccid,
		lasSecrets.Client,
		lasSecrets.Password,
		lasSecrets.LasEndpoint,
		lasSecrets.CcEndpoint,
	)

	// Step 6: Generate bearer token
	_, err = ltg.GenerateBearerToken(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Bearer Token Generation Failed",
			fmt.Sprintf("Failed to generate bearer token: %s", err.Error()),
		)
		return
	}

	// Step 7: Validate entitlement_name against LAS customer entitlements
	platform := strings.ReplaceAll(entitlementName[:6], " ", "_")
	entitlementResp, entErr := ltg.GetCustomerEntitlements(ctx, platform)
	if entErr != nil {
		resp.Diagnostics.AddError(
			"Failed to Fetch Customer Entitlements",
			fmt.Sprintf("Failed to fetch entitlements from LAS for platform '%s': %s", platform, entErr.Error()),
		)
		return
	}

	validEntitlement := false
	var validEntitlementNames []string
	if entitlements, ok := entitlementResp["entitlements"].([]interface{}); ok {
		for _, e := range entitlements {
			if obj, ok := e.(map[string]interface{}); ok {
				if name, ok := obj["type"].(string); ok {
					validEntitlementNames = append(validEntitlementNames, name)
					if name == entitlementName {
						validEntitlement = true
					}
				}
			}
		}
	}
	if !validEntitlement {
		resp.Diagnostics.AddError(
			"Invalid Entitlement Name",
			fmt.Sprintf("entitlement_name '%s' is not a valid entitlement for platform '%s'. Valid entitlements: %s", entitlementName, platform, strings.Join(validEntitlementNames, ", ")),
		)
		return
	}

	tflog.Info(ctx, "Entitlement name validated against LAS", map[string]interface{}{
		"entitlementName": entitlementName,
		"platform":        platform,
	})

	// Step 8: Get fingerprint and deregister if exists
	fingerprint, err := ltg.GetFingerprintForLSGUID(ctx)
	if err != nil {
		tflog.Warn(ctx, "Failed to get fingerprint", map[string]interface{}{"error": err.Error()})
		fingerprint = ""
	}

	tflog.Info(ctx, "Fingerprint lookup complete", map[string]interface{}{"fingerprint": fingerprint})

	// Step 9: Import offline activation request
	var importToken string
	if data.RestrictedMode.ValueBool() {
		lsid, pubkey, err := lasutils.ExtractLSIDAndPubKeyFromPackage(ctx, product, packageData)
		if err != nil {
			resp.Diagnostics.AddError(
				"LSID/PubKey Extraction Failed",
				fmt.Sprintf("Failed to extract lsid and pubkey from request package: %s", err.Error()),
			)
			return
		}
		importToken, err = ltg.ImportRestrictedOfflineActivationRequest(ctx, lsid, pubkey)
		if err != nil {
			resp.Diagnostics.AddError(
				"Import Restricted Request Failed",
				fmt.Sprintf("Failed to import restricted offline activation request: %s", err.Error()),
			)
			return
		}
	} else {
		importToken, err = ltg.ImportOfflineActivationRequest(ctx, packageData, fingerprint)
		if err != nil {
			resp.Diagnostics.AddError(
				"Import Request Failed",
				fmt.Sprintf("Failed to import offline activation request: %s", err.Error()),
			)
			return
		}
	}

	tflog.Info(ctx, "Import successful", map[string]interface{}{"importToken": importToken})

	// Step 10: Generate offline activation
	activationResp, err := ltg.GenerateOfflineActivation(ctx, importToken, entitlementName)
	if err != nil {
		resp.Diagnostics.AddError(
			"Activation Generation Failed",
			fmt.Sprintf("Failed to generate offline activation: %s", err.Error()),
		)
		return
	}

	activationID, ok := activationResp["newactivationid"].(string)
	if !ok {
		resp.Diagnostics.AddError(
			"Invalid Activation Response",
			"Failed to extract newactivationid from activation response",
		)
		return
	}

	activationFingerprint, ok := activationResp["lsfingerprint"].(string)
	if !ok {
		activationFingerprint = fingerprint
	}

	tflog.Info(ctx, "Activation generated", map[string]interface{}{"activationID": activationID})

	// Step 11: Export offline activation response (license blob)
	licenseBlob, err := ltg.ExportOfflineActivationResponse(ctx, activationID, activationFingerprint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Export Failed",
			fmt.Sprintf("Failed to export license blob: %s", err.Error()),
		)
		return
	}

	// Step 12: Save license blob to local file
	blobPath := fmt.Sprintf("/tmp/offline_token_%s_%s_activation.blob.tgz", deviceIP, hostname)
	if err := os.WriteFile(blobPath, licenseBlob, 0644); err != nil {
		resp.Diagnostics.AddError(
			"File Save Failed",
			fmt.Sprintf("Failed to save license blob: %s", err.Error()),
		)
		return
	}
	data.LicenseBlob = types.StringValue(blobPath)

	tflog.Info(ctx, "License blob saved", map[string]interface{}{"path": blobPath})

	// Step 13: Apply license blob to SDX device
	err = lasutils.ApplyLicenseBlobADM(ctx, deviceIP, username, password, licenseBlob)

	if err != nil {
		resp.Diagnostics.AddError(
			"License Application Failed",
			fmt.Sprintf("Failed to apply license blob to device: %s", err.Error()),
		)
		return
	}

	// Set computed values
	data.Id = types.StringValue(deviceIP)
	data.Status = types.StringValue("applied")
	data.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	tflog.Info(ctx, "License applied successfully", map[string]interface{}{"device_ip": deviceIP})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// extractIPFromEndpoint extracts IP address from endpoint URL
func extractIPFromEndpoint(endpoint string) string {
	// Remove protocol prefix if present
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")

	// Remove port if present
	if idx := strings.Index(endpoint, ":"); idx != -1 {
		endpoint = endpoint[:idx]
	}

	// Remove trailing slash
	endpoint = strings.TrimSuffix(endpoint, "/")

	return endpoint
}

func (r *nslaslicenseOfflineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of nslaslicense_offline Resource with Id: %s", resId))

	var data nslaslicenseOfflineModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// For offline license, we can verify if the license blob file still exists
	if !data.LicenseBlob.IsNull() {
		blobPath := data.LicenseBlob.ValueString()
		if _, err := os.Stat(blobPath); os.IsNotExist(err) {
			tflog.Warn(ctx, "License blob file no longer exists", map[string]interface{}{"path": blobPath})
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *nslaslicenseOfflineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data nslaslicenseOfflineModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NO-OP: Offline licenses cannot be updated in place
	// Any changes to key attributes require resource replacement
	tflog.Info(ctx, "Update operation is a no-op for offline license resource", map[string]interface{}{
		"device_ip": data.Id.ValueString(),
	})

	// Save data into Terraform state as-is
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *nslaslicenseOfflineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data nslaslicenseOfflineModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NO-OP: Offline licenses are not removed from the device on destroy
	// The license remains active on the device; only Terraform state is removed
	tflog.Info(ctx, "Delete operation is a no-op for offline license resource", map[string]interface{}{
		"device_ip": data.Id.ValueString(),
	})
}
