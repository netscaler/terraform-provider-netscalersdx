package device_profile

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*deviceProfileResource)(nil)
var _ resource.ResourceWithConfigure = (*deviceProfileResource)(nil)

func DeviceProfileResource() resource.Resource {
	return &deviceProfileResource{}
}

type deviceProfileResource struct {
	client *service.NitroClient
}

func (r *deviceProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_profile"
}

// Configure configures the client resource.
func (r *deviceProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *deviceProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = deviceProfileResourceSchema(ctx)
}

func (r *deviceProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of device_profile Resource")

	var data deviceProfileModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	deviceProfileReq := deviceProfileGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "device_profile"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, deviceProfileReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Example data value setting
	data.Id = types.StringValue(resID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	rreq := resource.ReadRequest{
		State:        resp.State,
		ProviderMeta: req.ProviderMeta,
	}
	rresp := resource.ReadResponse{
		State:       resp.State,
		Diagnostics: resp.Diagnostics,
	}

	r.Read(ctx, rreq, &rresp)

	*resp = resource.CreateResponse{
		State:       rresp.State,
		Diagnostics: rresp.Diagnostics,
	}

}

func (r *deviceProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of device_profile Resource with Id: %s", resId))

	var data deviceProfileModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "device_profile"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource device_profile: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.CbProfileName.IsNull() {
		data.CbProfileName = types.StringValue(getResponseData["cb_profile_name"].(string))
	}
	if !data.HttpPort.IsNull() {
		val, _ := strconv.Atoi(getResponseData["http_port"].(string))
		data.HttpPort = types.Int64Value(int64(val))
	}
	if !data.HttpsPort.IsNull() {
		val, _ := strconv.Atoi(getResponseData["https_port"].(string))
		data.HttpsPort = types.Int64Value(int64(val))
	}
	if !data.MaxWaitTimeReboot.IsNull() {
		data.MaxWaitTimeReboot = types.StringValue(getResponseData["max_wait_time_reboot"].(string))
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.NsProfileName.IsNull() {
		data.NsProfileName = types.StringValue(getResponseData["ns_profile_name"].(string))
	}
	if !data.Snmpauthprotocol.IsNull() {
		data.Snmpauthprotocol = types.StringValue(getResponseData["snmpauthprotocol"].(string))
	}
	if !data.Snmpcommunity.IsNull() {
		data.Snmpcommunity = types.StringValue(getResponseData["snmpcommunity"].(string))
	}
	if !data.Snmpprivprotocol.IsNull() {
		data.Snmpprivprotocol = types.StringValue(getResponseData["snmpprivprotocol"].(string))
	}
	if !data.Snmpsecuritylevel.IsNull() {
		data.Snmpsecuritylevel = types.StringValue(getResponseData["snmpsecuritylevel"].(string))
	}
	if !data.Snmpsecurityname.IsNull() {
		data.Snmpsecurityname = types.StringValue(getResponseData["snmpsecurityname"].(string))
	}
	if !data.Snmpversion.IsNull() {
		data.Snmpversion = types.StringValue(getResponseData["snmpversion"].(string))
	}
	if !data.SshPort.IsNull() {
		data.SshPort = types.StringValue(getResponseData["ssh_port"].(string))
	}
	if !data.SslCert.IsNull() {
		data.SslCert = types.StringValue(getResponseData["ssl_cert"].(string))
	}
	if !data.SslPrivateKey.IsNull() {
		data.SslPrivateKey = types.StringValue(getResponseData["ssl_private_key"].(string))
	}
	if !data.SvmNsComm.IsNull() {
		data.SvmNsComm = types.StringValue(getResponseData["svm_ns_comm"].(string))
	}
	if !data.Type.IsNull() {
		data.Type = types.StringValue(getResponseData["type"].(string))
	}
	if !data.UseGlobalSettingForCommunicationWithNs.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["use_global_setting_for_communication_with_ns"].(string))
		data.UseGlobalSettingForCommunicationWithNs = types.BoolValue(val)
	}
	if !data.Username.IsNull() {
		data.Username = types.StringValue(getResponseData["username"].(string))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *deviceProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of device_profile Resource")

	var data deviceProfileModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state deviceProfileModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "device_profile"
	requestPayload := deviceProfileGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, resourceId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Resource",
			fmt.Sprintf("Error updating resource: %s", err.Error()),
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of device_profile Resource")

	var data deviceProfileModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "device_profile"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
