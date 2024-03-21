package snmp_manager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpManagerResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP Agent Manager configuration resource.",
		Attributes: map[string]schema.Attribute{
			"community": schema.StringAttribute{
				Required:            true,
				Description:         "Community Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Community Name. Minimum length =  1 Maximum length =  128",
			},
			"ip_address": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Manager IPAddress.",
				MarkdownDescription: "Manager IPAddress.",
			},
			"netmask": schema.StringAttribute{
				Optional:            true,
				Description:         "Netmask. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Netmask. Minimum length =  1 Maximum length =  64",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpManagerModel struct {
	Community types.String `tfsdk:"community"`
	IpAddress types.String `tfsdk:"ip_address"`
	Netmask   types.String `tfsdk:"netmask"`
	Id        types.String `tfsdk:"id"`
}

func snmpManagerGetThePayloadFromtheConfig(ctx context.Context, data *snmpManagerModel) snmpManagerReq {
	tflog.Debug(ctx, "In snmpManagerGetThePayloadFromtheConfig Function")
	snmpManagerReqPayload := snmpManagerReq{
		Community: data.Community.ValueString(),
		IpAddress: data.IpAddress.ValueString(),
		Netmask:   data.Netmask.ValueString(),
	}
	return snmpManagerReqPayload
}
func snmpManagerSetAttrFromGet(ctx context.Context, data *snmpManagerModel, getResponseData map[string]interface{}) *snmpManagerModel {
	tflog.Debug(ctx, "In snmpManagerSetAttrFromGet Function")
	// if !data.Community.IsNull() {
	// 	data.Community = types.StringValue(getResponseData["community"].(string))
	// }
	if !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !data.Netmask.IsNull() {
		data.Netmask = types.StringValue(getResponseData["netmask"].(string))
	}
	return data
}

type snmpManagerReq struct {
	Community string `json:"community,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}
