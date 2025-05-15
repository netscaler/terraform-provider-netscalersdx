package aclrule

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func aclruleResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for ACL Rule resource.",
		Attributes: map[string]schema.Attribute{
			"action": schema.StringAttribute{
				Required:            true,
				Description:         "Action can be [Allow Block]. Minimum length =  4 Maximum length =  5",
				MarkdownDescription: "Action can be [Allow Block]. Minimum length =  4 Maximum length =  5",
			},
			"dst_port": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable external authentication.",
				MarkdownDescription: "Enable external authentication.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Rule Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Rule Name. Minimum length =  1 Maximum length =  128",
			},
			"priority": schema.Int64Attribute{
				Required:            true,
				Description:         "Priority. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Priority. Minimum value =  1 Maximum value =  ",
			},
			"protocol": schema.StringAttribute{
				Required:            true,
				Description:         "IP Protocol. The allowed values are [TCP UDP ICMP ANY]. Minimum length =  3 Maximum length =  4",
				MarkdownDescription: "IP Protocol. The allowed values are [TCP UDP ICMP ANY]. Minimum length =  3 Maximum length =  4",
			},
			"src_ip": schema.StringAttribute{
				Required:            true,
				Description:         "Source IP Address or Subnet. Minimum length =  3 Maximum length =  128",
				MarkdownDescription: "Source IP Address or Subnet. Minimum length =  3 Maximum length =  128",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type aclruleModel struct {
	Action   types.String `tfsdk:"action"`
	DstPort  types.String `tfsdk:"dst_port"`
	Name     types.String `tfsdk:"name"`
	Priority types.Int64  `tfsdk:"priority"`
	Protocol types.String `tfsdk:"protocol"`
	SrcIp    types.String `tfsdk:"src_ip"`
	Id       types.String `tfsdk:"id"`
}

func aclruleGetThePayloadFromtheConfig(ctx context.Context, data *aclruleModel) aclruleReq {
	tflog.Debug(ctx, "In aclruleGetThePayloadFromtheConfig Function")
	aclruleReqPayload := aclruleReq{
		Action:   data.Action.ValueString(),
		DstPort:  data.DstPort.ValueString(),
		Name:     data.Name.ValueString(),
		Priority: data.Priority.ValueInt64Pointer(),
		Protocol: data.Protocol.ValueString(),
		SrcIp:    data.SrcIp.ValueString(),
	}
	return aclruleReqPayload
}
func aclruleSetAttrFromGet(ctx context.Context, data *aclruleModel, getResponseData map[string]interface{}) *aclruleModel {
	tflog.Debug(ctx, "In aclruleSetAttrFromGet Function")
	data.Action = types.StringValue(getResponseData["action"].(string))
	data.DstPort = types.StringValue(getResponseData["dst_port"].(string))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.Priority = types.Int64Value(utils.StringToInt(getResponseData["priority"].(string)))
	data.Protocol = types.StringValue(getResponseData["protocol"].(string))
	data.SrcIp = types.StringValue(getResponseData["src_ip"].(string))
	return data
}

type aclruleReq struct {
	Action   string `json:"action,omitempty"`
	DstPort  string `json:"dst_port,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority *int64 `json:"priority,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	SrcIp    string `json:"src_ip,omitempty"`
}
