package system_settings

import (
	"context"
	"strconv"

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
				Description:         "Authorize the DeviceAPIProxy request.",
				MarkdownDescription: "Authorize the DeviceAPIProxy request.",
			},
			"basicauth": schema.BoolAttribute{
				Optional:            true,
				Description:         "Allow Basic Authentication Protocol.",
				MarkdownDescription: "Allow Basic Authentication Protocol.",
			},
			"disable_agent_old_password_input": schema.BoolAttribute{
				Optional:            true,
				Description:         "Disable old password input requirement while changing ADM agent password.",
				MarkdownDescription: "Disable old password input requirement while changing ADM agent password.",
			},
			"disk_utilization_threshold": schema.Int64Attribute{
				Optional:            true,
				Description:         "Disk utilization threshold after which data processing it stopped.",
				MarkdownDescription: "Disk utilization threshold after which data processing it stopped.",
			},
			"enable_apiproxy_credentials": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable API Proxy Credentials.",
				MarkdownDescription: "Enable API Proxy Credentials.",
			},
			"enable_certificate_download": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable Certificate Download.",
				MarkdownDescription: "Enable Certificate Download.",
			},
			"enable_cuxip": schema.BoolAttribute{
				Optional:            true,
				Description:         "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
				MarkdownDescription: "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
			},
			"enable_delete_interface_on_adc": schema.BoolAttribute{
				Optional:            true,
				Description:         "Flag to enable/disable deleting interface from ADCs on SDX.",
				MarkdownDescription: "Flag to enable/disable deleting interface from ADCs on SDX.",
			},
			"enable_nsrecover_login": schema.BoolAttribute{
				Optional:            true,
				Description:         "This setting enalbes nsrecover login for SVM.",
				MarkdownDescription: "This setting enalbes nsrecover login for SVM.",
			},
			"enable_session_timeout": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enables session timeout feature.",
				MarkdownDescription: "Enables session timeout feature.",
			},
			"enable_shell_access": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable Shell access for non-nsroot User(s).",
				MarkdownDescription: "Enable Shell access for non-nsroot User(s).",
			},
			"is_metering_enabled": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable Metering for NetScaler VPX on SDX.",
				MarkdownDescription: "Enable Metering for NetScaler VPX on SDX.",
			},
			"keep_adc_image_count": schema.Int64Attribute{
				Optional:            true,
				Description:         "Count for number of NetScaler images to be saved in Agent.",
				MarkdownDescription: "Count for number of NetScaler images to be saved in Agent.",
			},
			"keep_alive_ping_interval": schema.Int64Attribute{
				Optional:            true,
				Description:         "Agent web socket keep alive ping interval for the system.",
				MarkdownDescription: "Agent web socket keep alive ping interval for the system.",
			},
			"prompt_creds_for_stylebooks": schema.BoolAttribute{
				Optional:            true,
				Description:         "Prompt Credentials for Stylebooks.",
				MarkdownDescription: "Prompt Credentials for Stylebooks.",
			},
			"secure_access_only": schema.BoolAttribute{
				Optional:            true,
				Description:         "Secure Access only.",
				MarkdownDescription: "Secure Access only.",
			},
			"session_timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "Session timeout for the system.",
				MarkdownDescription: "Session timeout for the system.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Optional:            true,
				Description:         "Session timeout unit for the system. Possible Values: [ Minutes, Hours ]",
				MarkdownDescription: "Session timeout unit for the system. Possible Values: [ Minutes, Hours ]",
			},
			"svm_ns_comm": schema.StringAttribute{
				Optional:            true,
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
		AuthorizeDeviceapiproxy:      data.AuthorizeDeviceapiproxy.ValueBoolPointer(),
		Basicauth:                    data.Basicauth.ValueBoolPointer(),
		DisableAgentOldPasswordInput: data.DisableAgentOldPasswordInput.ValueBoolPointer(),
		DiskUtilizationThreshold:     data.DiskUtilizationThreshold.ValueInt64Pointer(),
		EnableApiproxyCredentials:    data.EnableApiproxyCredentials.ValueBoolPointer(),
		EnableCertificateDownload:    data.EnableCertificateDownload.ValueBoolPointer(),
		EnableCuxip:                  data.EnableCuxip.ValueBoolPointer(),
		EnableDeleteInterfaceOnAdc:   data.EnableDeleteInterfaceOnAdc.ValueBoolPointer(),
		EnableNsrecoverLogin:         data.EnableNsrecoverLogin.ValueBoolPointer(),
		EnableSessionTimeout:         data.EnableSessionTimeout.ValueBoolPointer(),
		EnableShellAccess:            data.EnableShellAccess.ValueBoolPointer(),
		IsMeteringEnabled:            data.IsMeteringEnabled.ValueBoolPointer(),
		KeepAdcImageCount:            data.KeepAdcImageCount.ValueInt64Pointer(),
		KeepAlivePingInterval:        data.KeepAlivePingInterval.ValueInt64Pointer(),
		PromptCredsForStylebooks:     data.PromptCredsForStylebooks.ValueBoolPointer(),
		SecureAccessOnly:             data.SecureAccessOnly.ValueBoolPointer(),
		SessionTimeout:               data.SessionTimeout.ValueInt64Pointer(),
		SessionTimeoutUnit:           data.SessionTimeoutUnit.ValueString(),
		SvmNsComm:                    data.SvmNsComm.ValueString(),
	}
	return systemSettingsReqPayload
}
func systemSettingsSetAttrFromGet(ctx context.Context, data *systemSettingsModel, getResponseData map[string]interface{}) *systemSettingsModel {
	tflog.Debug(ctx, "In systemSettingsSetAttrFromGet Function")
	if !data.AuthorizeDeviceapiproxy.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["authorize_deviceapiproxy"].(string))
		data.AuthorizeDeviceapiproxy = types.BoolValue(val)
	}
	if !data.Basicauth.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["basicauth"].(string))
		data.Basicauth = types.BoolValue(val)
	}
	if !data.DisableAgentOldPasswordInput.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["disable_agent_old_password_input"].(string))
		data.DisableAgentOldPasswordInput = types.BoolValue(val)
	}
	if !data.DiskUtilizationThreshold.IsNull() {
		val, _ := strconv.Atoi(getResponseData["disk_utilization_threshold"].(string))
		data.DiskUtilizationThreshold = types.Int64Value(int64(val))
	}
	if !data.EnableApiproxyCredentials.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_apiproxy_credentials"].(string))
		data.EnableApiproxyCredentials = types.BoolValue(val)
	}
	if !data.EnableCertificateDownload.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_certificate_download"].(string))
		data.EnableCertificateDownload = types.BoolValue(val)
	}
	if !data.EnableCuxip.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_cuxip"].(string))
		data.EnableCuxip = types.BoolValue(val)
	}
	if !data.EnableDeleteInterfaceOnAdc.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_delete_interface_on_adc"].(string))
		data.EnableDeleteInterfaceOnAdc = types.BoolValue(val)
	}
	if !data.EnableNsrecoverLogin.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_nsrecover_login"].(string))
		data.EnableNsrecoverLogin = types.BoolValue(val)
	}
	if !data.EnableSessionTimeout.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_session_timeout"].(string))
		data.EnableSessionTimeout = types.BoolValue(val)
	}
	if !data.EnableShellAccess.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_shell_access"].(string))
		data.EnableShellAccess = types.BoolValue(val)
	}
	if !data.IsMeteringEnabled.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_metering_enabled"].(string))
		data.IsMeteringEnabled = types.BoolValue(val)
	}
	if !data.KeepAdcImageCount.IsNull() {
		val, _ := strconv.Atoi(getResponseData["keep_adc_image_count"].(string))
		data.KeepAdcImageCount = types.Int64Value(int64(val))
	}
	if !data.KeepAlivePingInterval.IsNull() {
		val, _ := strconv.Atoi(getResponseData["keep_alive_ping_interval"].(string))
		data.KeepAlivePingInterval = types.Int64Value(int64(val))
	}
	if !data.PromptCredsForStylebooks.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["prompt_creds_for_stylebooks"].(string))
		data.PromptCredsForStylebooks = types.BoolValue(val)
	}
	if !data.SecureAccessOnly.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["secure_access_only"].(string))
		data.SecureAccessOnly = types.BoolValue(val)
	}
	if !data.SessionTimeout.IsNull() {
		val, _ := strconv.Atoi(getResponseData["session_timeout"].(string))
		data.SessionTimeout = types.Int64Value(int64(val))
	}
	if !data.SessionTimeoutUnit.IsNull() {
		data.SessionTimeoutUnit = types.StringValue(getResponseData["session_timeout_unit"].(string))
	}
	if !data.SvmNsComm.IsNull() {
		data.SvmNsComm = types.StringValue(getResponseData["svm_ns_comm"].(string))
	}
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
