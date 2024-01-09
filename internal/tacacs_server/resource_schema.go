package tacacs_server

import (
	"context"
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
				Description:         "Enable accounting in the tacacs server.",
				MarkdownDescription: "Enable accounting in the tacacs server.",
			},
			"auth_timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
			},
			"group_attr_name": schema.StringAttribute{
				Optional:            true,
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

func tacacsServerGetThePayloadFromtheConfig(ctx context.Context, data *tacacsServerModel) tacacsServerReq {
	tflog.Debug(ctx, "In tacacsServerGetThePayloadFromtheConfig Function")
	tacacsServerReqPayload := tacacsServerReq{
		Accounting:    data.Accounting.ValueBool(),
		AuthTimeout:   data.AuthTimeout.ValueInt64(),
		GroupAttrName: data.GroupAttrName.ValueString(),
		IpAddress:     data.IpAddress.ValueString(),
		Name:          data.Name.ValueString(),
		Port:          data.Port.ValueInt64(),
		TacacsKey:     data.TacacsKey.ValueString(),
	}
	return tacacsServerReqPayload
}

type tacacsServerReq struct {
	Accounting    bool   `json:"accounting,omitempty"`
	AuthTimeout   int64  `json:"auth_timeout,omitempty"`
	GroupAttrName string `json:"group_attr_name,omitempty"`
	IpAddress     string `json:"ip_address,omitempty"`
	Name          string `json:"name,omitempty"`
	Port          int64  `json:"port,omitempty"`
	TacacsKey     string `json:"tacacs_key,omitempty"`
}
