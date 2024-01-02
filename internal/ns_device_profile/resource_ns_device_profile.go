package ns_device_profile

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &nsDeviceProfileResource{}
	_ resource.ResourceWithConfigure = &nsDeviceProfileResource{}
)

// nsDeviceProfileResource defines the resource implementation.
type nsDeviceProfileResource struct {
	client *service.NitroClient
}

func NsDeviceProfileResource() resource.Resource {
	return &nsDeviceProfileResource{}
}

// Metadata returns the resource type name.
func (r *nsDeviceProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ns_device_profile"
}

// Configure configures the client resource.
func (r *nsDeviceProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *nsDeviceProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Configuration for Device profile for NetScaler ADC(MPX/VPX/CPX/Gateway) instances resource.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Profile Name. Minimum length =  1 Maximum length =  128",
			},
			"svm_ns_comm": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Communication protocol (http or https) with Instances. Minimum length =  1 Maximum length =  10",
			},
			"use_global_setting_for_communication_with_ns": schema.BoolAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "True, if the communication with Instance needs to be global and not device specific.",
			},
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Id is system generated key for all the device profiles.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Profile Type, This must be with in specified supported instance types: blx,sdvanvw,ns,nssdx,cbwanopt,cpx. Minimum length =  1 Maximum length =  128",
			},
			"ns_profile_name": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Profile Name, This is one of the already created NetScaler ADC profiles.",
			},
			"password": schema.StringAttribute{
				Required:  true,
				Computed:  false,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Instance credentials.Password for this profile. Minimum length =  1 Maximum length =  127",
			},
			"snmpsecuritylevel": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 security level for this profile. Possible values: ['NoAuthNoPriv', 'AuthNoPriv', 'AuthPriv']",
			},
			"username": schema.StringAttribute{
				Required: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Instance credentials.Username provided in the profile will be used to contact the instance. Minimum length =  1 Maximum length =  127",
			},
			"snmpauthprotocol": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 auth protocol for this profile.",
			},
			"ssh_port": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SSH port to connect to the device.",
			},
			"snmpprivprotocol": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 priv protocol for this profile.",
			},
			"host_password": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host Password for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"snmpversion": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP version for this profile.",
			},
			"passphrase": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Passphrase with which private key is encrypted.",
			},
			"host_username": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host User Name for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"snmpsecurityname": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 security name for this profile. Maximum length =  31",
			},
			"ssl_private_key": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SSL Private Key for key based authentication.",
			},
			"ssl_cert": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SSL Certificate for certificate based authentication.",
			},
			"http_port": schema.Int64Attribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "HTTP port to connect to the device.",
			},
			"snmpcommunity": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP community for this profile. Maximum length =  31",
			},
			"https_port": schema.Int64Attribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "HTTPS port to connect to the device.",
			},
			"max_wait_time_reboot": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Max waiting time to reboot NetScaler ADC.",
			},
			"snmpprivpassword": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 priv password for this profile. Minimum length =  8 Maximum length =  31",
			},
			"cb_profile_name": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Profile Name, This is one of the already created NetScaler SD-WAN profiles.",
			},
			"snmpauthpassword": schema.StringAttribute{
				Optional: true,
				Computed: false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP v3 auth password for this profile. Minimum length =  8 Maximum length =  31",
			},
		},
	}
}

// Create creates a new resources and adds it into the Terraform state.
func (r *nsDeviceProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "[DEBUG] In Create Method of nsDeviceProfileResource")

	var data *nsDeviceProfileResourceModel

	// Read Terraform plan data into the model( proVpxResourceModel )
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	nsDeviceProfileReq := nsDeviceProfileGetThePayloadFromtheConfig(ctx, data)

	endpoint := "ns_device_profile"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, nsDeviceProfileReq)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Resource",
			fmt.Sprintf("Error creating resource: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Set the ID of the created resource
	data.Id = types.StringValue(resID)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
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

func (r *nsDeviceProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Read Method of nsDeviceProfileResource with Id: %s", resId))

	var state *nsDeviceProfileResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "ns_device_profile"

	data, err := r.client.GetResource(endpoint, state.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ns_device_profile: %v from state because it is not present in the remote", state.Id.ValueString()))
		return
	}

	getResponseData := data[endpoint].([]interface{})[0].(map[string]interface{})

	if !state.Name.IsNull() || getResponseData["name"] != nil {
		state.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !state.SvmNsComm.IsNull() {
		state.SvmNsComm = types.StringValue(getResponseData["svm_ns_comm"].(string))
	}

	if !state.UseGlobalSettingForCommunicationWithNs.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["use_global_setting_for_communication_with_ns"].(string))
		state.UseGlobalSettingForCommunicationWithNs = types.BoolValue(val)
	}
	if !state.Type.IsNull() {
		state.Type = types.StringValue(getResponseData["type"].(string))
	}
	if !state.NsProfileName.IsNull() {
		state.NsProfileName = types.StringValue(getResponseData["ns_profile_name"].(string))
	}
	if !state.Snmpsecuritylevel.IsNull() {
		state.Snmpsecuritylevel = types.StringValue(getResponseData["snmpsecuritylevel"].(string))
	}
	if !state.Username.IsNull() {
		state.Username = types.StringValue(getResponseData["username"].(string))
	}
	if !state.Snmpauthprotocol.IsNull() {
		state.Snmpauthprotocol = types.StringValue(getResponseData["snmpauthprotocol"].(string))
	}
	if !state.SshPort.IsNull() {
		state.SshPort = types.StringValue(getResponseData["ssh_port"].(string))
	}
	if !state.Snmpprivprotocol.IsNull() {
		state.Snmpprivprotocol = types.StringValue(getResponseData["snmpprivprotocol"].(string))
	}
	if !state.HostPassword.IsNull() {
		state.HostPassword = types.StringValue(getResponseData["host_password"].(string))
	}
	if !state.Snmpversion.IsNull() {
		state.Snmpversion = types.StringValue(getResponseData["snmpversion"].(string))
	}
	if !state.Passphrase.IsNull() {
		state.Passphrase = types.StringValue(getResponseData["passphrase"].(string))
	}
	if !state.HostUsername.IsNull() {
		state.HostUsername = types.StringValue(getResponseData["host_username"].(string))
	}
	if !state.Snmpsecurityname.IsNull() {
		state.Snmpsecurityname = types.StringValue(getResponseData["snmpsecurityname"].(string))
	}
	if !state.SslPrivateKey.IsNull() {
		state.SslPrivateKey = types.StringValue(getResponseData["ssl_private_key"].(string))
	}
	if !state.SslCert.IsNull() {
		state.SslCert = types.StringValue(getResponseData["ssl_cert"].(string))
	}
	if !state.HttpPort.IsNull() {
		val, _ := strconv.Atoi(getResponseData["http_port"].(string))
		state.HttpPort = types.Int64Value(int64(val))
	}
	if !state.Snmpcommunity.IsNull() {
		state.Snmpcommunity = types.StringValue(getResponseData["snmpcommunity"].(string))
	}
	if !state.HttpsPort.IsNull() {
		val, _ := strconv.Atoi(getResponseData["https_port"].(string))
		state.HttpsPort = types.Int64Value(int64(val))
	}
	if !state.MaxWaitTimeReboot.IsNull() {
		state.MaxWaitTimeReboot = types.StringValue(getResponseData["max_wait_time_reboot"].(string))
	}
	if !state.Snmpprivpassword.IsNull() {
		state.Snmpprivpassword = types.StringValue(getResponseData["snmpprivpassword"].(string))
	}
	if !state.CbProfileName.IsNull() {
		state.CbProfileName = types.StringValue(getResponseData["cb_profile_name"].(string))
	}
	if !state.Snmpauthpassword.IsNull() {
		state.Snmpauthpassword = types.StringValue(getResponseData["snmpauthpassword"].(string))
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *nsDeviceProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *nsDeviceProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Delete Method of nsDeviceProfileResource with Id: %s", resId))

	var data *nsDeviceProfileResourceModel

	// Read terraform prior state data into the model
	diag := req.State.Get(ctx, &data)

	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "ns_device_profile"
	// Delete the resource
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Resource",
			fmt.Sprintf("Error deleting resource: %s", err.Error()),
		)
		return
	}
}
