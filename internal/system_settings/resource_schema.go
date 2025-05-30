package system_settings

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func systemSettingsResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for System Settings resource.",
		Attributes: map[string]schema.Attribute{
			"authorize_deviceapiproxy": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Authorize the DeviceAPIProxy request.",
				MarkdownDescription: "Authorize the DeviceAPIProxy request.",
			},
			"basicauth": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Allow Basic Authentication Protocol.",
				MarkdownDescription: "Allow Basic Authentication Protocol.",
			},
			"disable_agent_old_password_input": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Disable old password input requirement while changing ADM agent password.",
				MarkdownDescription: "Disable old password input requirement while changing ADM agent password.",
			},
			"disk_utilization_threshold": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Disk utilization threshold after which data processing it stopped.",
				MarkdownDescription: "Disk utilization threshold after which data processing it stopped.",
			},
			"enable_apiproxy_credentials": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable API Proxy Credentials.",
				MarkdownDescription: "Enable API Proxy Credentials.",
			},
			"enable_certificate_download": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable Certificate Download.",
				MarkdownDescription: "Enable Certificate Download.",
			},
			"enable_cuxip": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
				MarkdownDescription: "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
			},
			"enable_delete_interface_on_adc": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Flag to enable/disable deleting interface from ADCs on SDX.",
				MarkdownDescription: "Flag to enable/disable deleting interface from ADCs on SDX.",
			},
			"enable_nsrecover_login": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "This setting enalbes nsrecover login for SVM.",
				MarkdownDescription: "This setting enalbes nsrecover login for SVM.",
			},
			"enable_session_timeout": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enables session timeout feature.",
				MarkdownDescription: "Enables session timeout feature.",
			},
			"enable_shell_access": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable Shell access for non-nsroot User(s).",
				MarkdownDescription: "Enable Shell access for non-nsroot User(s).",
			},
			"is_metering_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable Metering for NetScaler VPX on SDX.",
				MarkdownDescription: "Enable Metering for NetScaler VPX on SDX.",
			},
			"keep_adc_image_count": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Count for number of NetScaler images to be saved in Agent.",
				MarkdownDescription: "Count for number of NetScaler images to be saved in Agent.",
			},
			"keep_alive_ping_interval": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Agent web socket keep alive ping interval for the system.",
				MarkdownDescription: "Agent web socket keep alive ping interval for the system.",
			},
			"prompt_creds_for_stylebooks": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Prompt Credentials for Stylebooks.",
				MarkdownDescription: "Prompt Credentials for Stylebooks.",
			},
			"secure_access_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Secure Access only.",
				MarkdownDescription: "Secure Access only.",
			},
			"session_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Session timeout for the system.",
				MarkdownDescription: "Session timeout for the system.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Session timeout unit for the system. Possible Values: [ Minutes, Hours ]",
				MarkdownDescription: "Session timeout unit for the system. Possible Values: [ Minutes, Hours ]",
			},
			"svm_ns_comm": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Communication with Instances. Minimum length =  1 Maximum length =  10. Possible Values: [ http, https ]",
				MarkdownDescription: "Communication with Instances. Minimum length =  1 Maximum length =  10. Possible Values: [ http, https ]",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type systemSettingsModel struct {
	AuthorizeDeviceapiproxy      types.Bool   `tfsdk:"authorize_deviceapiproxy"`
	Basicauth                    types.Bool   `tfsdk:"basicauth"`
	DisableAgentOldPasswordInput types.Bool   `tfsdk:"disable_agent_old_password_input"`
	DiskUtilizationThreshold     types.Int64  `tfsdk:"disk_utilization_threshold"`
	EnableApiproxyCredentials    types.Bool   `tfsdk:"enable_apiproxy_credentials"`
	EnableCertificateDownload    types.Bool   `tfsdk:"enable_certificate_download"`
	EnableCuxip                  types.Bool   `tfsdk:"enable_cuxip"`
	EnableDeleteInterfaceOnAdc   types.Bool   `tfsdk:"enable_delete_interface_on_adc"`
	EnableNsrecoverLogin         types.Bool   `tfsdk:"enable_nsrecover_login"`
	EnableSessionTimeout         types.Bool   `tfsdk:"enable_session_timeout"`
	EnableShellAccess            types.Bool   `tfsdk:"enable_shell_access"`
	IsMeteringEnabled            types.Bool   `tfsdk:"is_metering_enabled"`
	KeepAdcImageCount            types.Int64  `tfsdk:"keep_adc_image_count"`
	KeepAlivePingInterval        types.Int64  `tfsdk:"keep_alive_ping_interval"`
	PromptCredsForStylebooks     types.Bool   `tfsdk:"prompt_creds_for_stylebooks"`
	SecureAccessOnly             types.Bool   `tfsdk:"secure_access_only"`
	SessionTimeout               types.Int64  `tfsdk:"session_timeout"`
	SessionTimeoutUnit           types.String `tfsdk:"session_timeout_unit"`
	SvmNsComm                    types.String `tfsdk:"svm_ns_comm"`
	Id                           types.String `tfsdk:"id"`
}

func systemSettingsGetThePayloadFromtheConfig(ctx context.Context, data *systemSettingsModel) systemSettingsReq {
	tflog.Debug(ctx, "In systemSettingsGetThePayloadFromtheConfig Function")
	systemSettingsReqPayload := systemSettingsReq{
		SessionTimeoutUnit: data.SessionTimeoutUnit.ValueString(),
		SvmNsComm:          data.SvmNsComm.ValueString(),
	}
	if !data.AuthorizeDeviceapiproxy.IsNull() && !data.AuthorizeDeviceapiproxy.IsUnknown() {
		systemSettingsReqPayload.AuthorizeDeviceapiproxy = data.AuthorizeDeviceapiproxy.ValueBoolPointer()
	}
	if !data.Basicauth.IsNull() && !data.Basicauth.IsUnknown() {
		systemSettingsReqPayload.Basicauth = data.Basicauth.ValueBoolPointer()
	}
	if !data.DisableAgentOldPasswordInput.IsNull() && !data.DisableAgentOldPasswordInput.IsUnknown() {
		systemSettingsReqPayload.DisableAgentOldPasswordInput = data.DisableAgentOldPasswordInput.ValueBoolPointer()
	}
	if !data.DiskUtilizationThreshold.IsNull() && !data.DiskUtilizationThreshold.IsUnknown() {
		systemSettingsReqPayload.DiskUtilizationThreshold = data.DiskUtilizationThreshold.ValueInt64Pointer()
	}
	if !data.EnableApiproxyCredentials.IsNull() && !data.EnableApiproxyCredentials.IsUnknown() {
		systemSettingsReqPayload.EnableApiproxyCredentials = data.EnableApiproxyCredentials.ValueBoolPointer()
	}
	if !data.EnableCertificateDownload.IsNull() && !data.EnableCertificateDownload.IsUnknown() {
		systemSettingsReqPayload.EnableCertificateDownload = data.EnableCertificateDownload.ValueBoolPointer()
	}
	if !data.EnableCuxip.IsNull() && !data.EnableCuxip.IsUnknown() {
		systemSettingsReqPayload.EnableCuxip = data.EnableCuxip.ValueBoolPointer()
	}
	if !data.EnableDeleteInterfaceOnAdc.IsNull() && !data.EnableDeleteInterfaceOnAdc.IsUnknown() {
		systemSettingsReqPayload.EnableDeleteInterfaceOnAdc = data.EnableDeleteInterfaceOnAdc.ValueBoolPointer()
	}
	if !data.EnableNsrecoverLogin.IsNull() && !data.EnableNsrecoverLogin.IsUnknown() {
		systemSettingsReqPayload.EnableNsrecoverLogin = data.EnableNsrecoverLogin.ValueBoolPointer()
	}
	if !data.EnableSessionTimeout.IsNull() && !data.EnableSessionTimeout.IsUnknown() {
		systemSettingsReqPayload.EnableSessionTimeout = data.EnableSessionTimeout.ValueBoolPointer()
	}
	if !data.EnableShellAccess.IsNull() && !data.EnableShellAccess.IsUnknown() {
		systemSettingsReqPayload.EnableShellAccess = data.EnableShellAccess.ValueBoolPointer()
	}
	if !data.IsMeteringEnabled.IsNull() && !data.IsMeteringEnabled.IsUnknown() {
		systemSettingsReqPayload.IsMeteringEnabled = data.IsMeteringEnabled.ValueBoolPointer()
	}
	if !data.KeepAdcImageCount.IsNull() && !data.KeepAdcImageCount.IsUnknown() {
		systemSettingsReqPayload.KeepAdcImageCount = data.KeepAdcImageCount.ValueInt64Pointer()
	}
	if !data.KeepAlivePingInterval.IsNull() && !data.KeepAlivePingInterval.IsUnknown() {
		systemSettingsReqPayload.KeepAlivePingInterval = data.KeepAlivePingInterval.ValueInt64Pointer()
	}
	if !data.PromptCredsForStylebooks.IsNull() && !data.PromptCredsForStylebooks.IsUnknown() {
		systemSettingsReqPayload.PromptCredsForStylebooks = data.PromptCredsForStylebooks.ValueBoolPointer()
	}
	if !data.SecureAccessOnly.IsNull() && !data.SecureAccessOnly.IsUnknown() {
		systemSettingsReqPayload.SecureAccessOnly = data.SecureAccessOnly.ValueBoolPointer()
	}
	if !data.SessionTimeout.IsNull() && !data.SessionTimeout.IsUnknown() {
		systemSettingsReqPayload.SessionTimeout = data.SessionTimeout.ValueInt64Pointer()
	}

	return systemSettingsReqPayload
}
func systemSettingsSetAttrFromGet(ctx context.Context, data *systemSettingsModel, getResponseData map[string]interface{}) *systemSettingsModel {
	tflog.Debug(ctx, "In systemSettingsSetAttrFromGet Function")

	data.AuthorizeDeviceapiproxy = types.BoolValue(utils.StringToBool(getResponseData["authorize_deviceapiproxy"].(string)))
	data.Basicauth = types.BoolValue(utils.StringToBool(getResponseData["basicauth"].(string)))
	data.DisableAgentOldPasswordInput = types.BoolValue(utils.StringToBool(getResponseData["disable_agent_old_password_input"].(string)))
	data.DiskUtilizationThreshold = types.Int64Value(utils.StringToInt(getResponseData["disk_utilization_threshold"].(string)))
	data.EnableApiproxyCredentials = types.BoolValue(utils.StringToBool(getResponseData["enable_apiproxy_credentials"].(string)))
	data.EnableCertificateDownload = types.BoolValue(utils.StringToBool(getResponseData["enable_certificate_download"].(string)))
	data.EnableCuxip = types.BoolValue(utils.StringToBool(getResponseData["enable_cuxip"].(string)))
	data.EnableDeleteInterfaceOnAdc = types.BoolValue(utils.StringToBool(getResponseData["enable_delete_interface_on_adc"].(string)))
	data.EnableNsrecoverLogin = types.BoolValue(utils.StringToBool(getResponseData["enable_nsrecover_login"].(string)))
	data.EnableSessionTimeout = types.BoolValue(utils.StringToBool(getResponseData["enable_session_timeout"].(string)))
	data.EnableShellAccess = types.BoolValue(utils.StringToBool(getResponseData["enable_shell_access"].(string)))
	data.IsMeteringEnabled = types.BoolValue(utils.StringToBool(getResponseData["is_metering_enabled"].(string)))
	data.KeepAdcImageCount = types.Int64Value(utils.StringToInt(getResponseData["keep_adc_image_count"].(string)))
	data.KeepAlivePingInterval = types.Int64Value(utils.StringToInt(getResponseData["keep_alive_ping_interval"].(string)))
	data.PromptCredsForStylebooks = types.BoolValue(utils.StringToBool(getResponseData["prompt_creds_for_stylebooks"].(string)))
	data.SecureAccessOnly = types.BoolValue(utils.StringToBool(getResponseData["secure_access_only"].(string)))
	data.SessionTimeout = types.Int64Value(utils.StringToInt(getResponseData["session_timeout"].(string)))
	data.SessionTimeoutUnit = types.StringValue(getResponseData["session_timeout_unit"].(string))
	data.SvmNsComm = types.StringValue(getResponseData["svm_ns_comm"].(string))

	return data
}

type systemSettingsReq struct {
	AuthorizeDeviceapiproxy      *bool  `json:"authorize_deviceapiproxy,omitempty"`
	Basicauth                    *bool  `json:"basicauth,omitempty"`
	DisableAgentOldPasswordInput *bool  `json:"disable_agent_old_password_input,omitempty"`
	DiskUtilizationThreshold     *int64 `json:"disk_utilization_threshold,omitempty"`
	EnableApiproxyCredentials    *bool  `json:"enable_apiproxy_credentials,omitempty"`
	EnableCertificateDownload    *bool  `json:"enable_certificate_download,omitempty"`
	EnableCuxip                  *bool  `json:"enable_cuxip,omitempty"`
	EnableDeleteInterfaceOnAdc   *bool  `json:"enable_delete_interface_on_adc,omitempty"`
	EnableNsrecoverLogin         *bool  `json:"enable_nsrecover_login,omitempty"`
	EnableSessionTimeout         *bool  `json:"enable_session_timeout,omitempty"`
	EnableShellAccess            *bool  `json:"enable_shell_access,omitempty"`
	IsMeteringEnabled            *bool  `json:"is_metering_enabled,omitempty"`
	KeepAdcImageCount            *int64 `json:"keep_adc_image_count,omitempty"`
	KeepAlivePingInterval        *int64 `json:"keep_alive_ping_interval,omitempty"`
	PromptCredsForStylebooks     *bool  `json:"prompt_creds_for_stylebooks,omitempty"`
	SecureAccessOnly             *bool  `json:"secure_access_only,omitempty"`
	SessionTimeout               *int64 `json:"session_timeout,omitempty"`
	SessionTimeoutUnit           string `json:"session_timeout_unit,omitempty"`
	SvmNsComm                    string `json:"svm_ns_comm,omitempty"`
}
