package tacacs_server

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tacacsServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for TACACS Server configuration resource.",
		Attributes: map[string]schema.Attribute{
			"accounting": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable accounting in the tacacs server.",
				MarkdownDescription: "Enable accounting in the tacacs server.",
			},
			"auth_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
			},
			"group_attr_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The Attribute name for group extraction from the ACS server. If not passed, then groups will not be extracted. No other harm. Maximum length =  64",
				MarkdownDescription: "The Attribute name for group extraction from the ACS server. If not passed, then groups will not be extracted. No other harm. Maximum length =  64",
			},
			"ip_address": schema.StringAttribute{
				Required:            true,
				Description:         "IP Address of TACACS server. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IP Address of TACACS server. Minimum length =  1 Maximum length =  64",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of TACACS server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of TACACS server. Minimum length =  1 Maximum length =  128",
			},
			"port": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "port number of TACACS server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "port number of TACACS server. Minimum value =  1 Maximum value =  ",
			},
			"tacacs_key": schema.StringAttribute{
				Required:            true,
				Description:         "Key shared between the TACACS+ server and clients. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Key shared between the TACACS+ server and clients. Minimum length =  1 Maximum length =  64",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type tacacsServerModel struct {
	Accounting    types.Bool   `tfsdk:"accounting"`
	AuthTimeout   types.Int64  `tfsdk:"auth_timeout"`
	GroupAttrName types.String `tfsdk:"group_attr_name"`
	IpAddress     types.String `tfsdk:"ip_address"`
	Name          types.String `tfsdk:"name"`
	Port          types.Int64  `tfsdk:"port"`
	TacacsKey     types.String `tfsdk:"tacacs_key"`
	Id            types.String `tfsdk:"id"`
}

func tacacsServerModelSetAttrFromGet(ctx context.Context, data *tacacsServerModel, getResponseData map[string]interface{}) *tacacsServerModel {

	data.GroupAttrName = types.StringValue(getResponseData["group_attr_name"].(string))
	data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.Port = types.Int64Value(utils.StringToInt(getResponseData["port"].(string)))
	data.AuthTimeout = types.Int64Value(utils.StringToInt(getResponseData["auth_timeout"].(string)))
	data.Accounting = types.BoolValue(utils.StringToBool(getResponseData["accounting"].(string)))

	return data
}

func tacacsServerGetThePayloadFromtheConfig(ctx context.Context, data *tacacsServerModel) tacacsServerReq {
	tflog.Debug(ctx, "In tacacsServerGetThePayloadFromtheConfig Function")
	tacacsServerReqPayload := tacacsServerReq{
		GroupAttrName: data.GroupAttrName.ValueString(),
		IpAddress:     data.IpAddress.ValueString(),
		Name:          data.Name.ValueString(),
		TacacsKey:     data.TacacsKey.ValueString(),
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		tacacsServerReqPayload.Port = data.Port.ValueInt64Pointer()
	}
	if !data.AuthTimeout.IsNull() && !data.AuthTimeout.IsUnknown() {
		tacacsServerReqPayload.AuthTimeout = data.AuthTimeout.ValueInt64Pointer()
	}
	if !data.Accounting.IsNull() && !data.Accounting.IsUnknown() {
		tacacsServerReqPayload.Accounting = data.Accounting.ValueBoolPointer()
	}

	return tacacsServerReqPayload
}

type tacacsServerReq struct {
	Accounting    *bool  `json:"accounting,omitempty"`
	AuthTimeout   *int64 `json:"auth_timeout,omitempty"`
	GroupAttrName string `json:"group_attr_name,omitempty"`
	IpAddress     string `json:"ip_address,omitempty"`
	Name          string `json:"name,omitempty"`
	Port          *int64 `json:"port,omitempty"`
	TacacsKey     string `json:"tacacs_key,omitempty"`
}
