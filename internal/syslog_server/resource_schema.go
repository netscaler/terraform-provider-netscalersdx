// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package syslog_server

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func syslogServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Syslog Server resource.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Syslog server name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Syslog server name. Minimum length =  1 Maximum length =  128",
			},
			"port": schema.Int64Attribute{
				Required:            true,
				Description:         "Syslog server port. Maximum value =  ",
				MarkdownDescription: "Syslog server port. Maximum value =  ",
			},
			"ip_address": schema.StringAttribute{
				Required:            true,
				Description:         "Syslog server IP address. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Syslog server IP address. Minimum length =  1 Maximum length =  64",
			},
			"log_level_all": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send logs of all levels to this syslog server.",
				MarkdownDescription: "Send logs of all levels to this syslog server.",
			},
			"log_level_critical": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send logs of level critical to this syslog server.",
				MarkdownDescription: "Send logs of level critical to this syslog server.",
			},
			"log_level_error": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send logs of level error to this syslog server.",
				MarkdownDescription: "Send logs of level error to this syslog server.",
			},
			"log_level_info": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send logs of level info to this syslog server.",
				MarkdownDescription: "Send logs of level info to this syslog server.",
			},
			"log_level_none": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send no logs to this syslog server.",
				MarkdownDescription: "Send no logs to this syslog server.",
			},
			"log_level_warning": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Send logs of level warning to this syslog server.",
				MarkdownDescription: "Send logs of level warning to this syslog server.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type syslogServerModel struct {
	IpAddress        types.String `tfsdk:"ip_address"`
	LogLevelAll      types.Bool   `tfsdk:"log_level_all"`
	LogLevelCritical types.Bool   `tfsdk:"log_level_critical"`
	LogLevelError    types.Bool   `tfsdk:"log_level_error"`
	LogLevelInfo     types.Bool   `tfsdk:"log_level_info"`
	LogLevelNone     types.Bool   `tfsdk:"log_level_none"`
	LogLevelWarning  types.Bool   `tfsdk:"log_level_warning"`
	Name             types.String `tfsdk:"name"`
	Port             types.Int64  `tfsdk:"port"`
	Id               types.String `tfsdk:"id"`
}

func syslogServerModelSetAttrFromGet(ctx context.Context, data *syslogServerModel, getResponseData map[string]interface{}) *syslogServerModel {
	tflog.Debug(ctx, "In syslogServerModelSetAttrFromGet Function")

	data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	data.LogLevelAll = types.BoolValue(utils.StringToBool(getResponseData["log_level_all"].(string)))
	data.LogLevelCritical = types.BoolValue(utils.StringToBool(getResponseData["log_level_critical"].(string)))
	data.LogLevelError = types.BoolValue(utils.StringToBool(getResponseData["log_level_error"].(string)))
	data.LogLevelInfo = types.BoolValue(utils.StringToBool(getResponseData["log_level_info"].(string)))
	data.LogLevelNone = types.BoolValue(utils.StringToBool(getResponseData["log_level_none"].(string)))
	data.LogLevelWarning = types.BoolValue(utils.StringToBool(getResponseData["log_level_warning"].(string)))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.Port = types.Int64Value(utils.StringToInt(getResponseData["port"].(string)))

	return data
}

func syslogServerGetThePayloadFromtheConfig(ctx context.Context, data *syslogServerModel) syslogServerReq {
	tflog.Debug(ctx, "In syslogServerGetThePayloadFromtheConfig Function")
	syslogServerReqPayload := syslogServerReq{
		IpAddress: data.IpAddress.ValueString(),
		Name:      data.Name.ValueString(),
	}
	if !data.LogLevelAll.IsNull() && !data.LogLevelAll.IsUnknown() {
		syslogServerReqPayload.LogLevelAll = data.LogLevelAll.ValueBoolPointer()
	}
	if !data.LogLevelCritical.IsNull() && !data.LogLevelCritical.IsUnknown() {
		syslogServerReqPayload.LogLevelCritical = data.LogLevelCritical.ValueBoolPointer()
	}
	if !data.LogLevelError.IsNull() && !data.LogLevelError.IsUnknown() {
		syslogServerReqPayload.LogLevelError = data.LogLevelError.ValueBoolPointer()
	}
	if !data.LogLevelInfo.IsNull() && !data.LogLevelInfo.IsUnknown() {
		syslogServerReqPayload.LogLevelInfo = data.LogLevelInfo.ValueBoolPointer()
	}
	if !data.LogLevelNone.IsNull() && !data.LogLevelNone.IsUnknown() {

		syslogServerReqPayload.LogLevelNone = data.LogLevelNone.ValueBoolPointer()
	}
	if !data.LogLevelWarning.IsNull() && !data.LogLevelWarning.IsUnknown() {
		syslogServerReqPayload.LogLevelWarning = data.LogLevelWarning.ValueBoolPointer()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		syslogServerReqPayload.Port = data.Port.ValueInt64Pointer()
	}

	return syslogServerReqPayload
}

type syslogServerReq struct {
	IpAddress        string `json:"ip_address,omitempty"`
	LogLevelAll      *bool  `json:"log_level_all,omitempty"`
	LogLevelCritical *bool  `json:"log_level_critical,omitempty"`
	LogLevelError    *bool  `json:"log_level_error,omitempty"`
	LogLevelInfo     *bool  `json:"log_level_info,omitempty"`
	LogLevelNone     *bool  `json:"log_level_none,omitempty"`
	LogLevelWarning  *bool  `json:"log_level_warning,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             *int64 `json:"port,omitempty"`
}
