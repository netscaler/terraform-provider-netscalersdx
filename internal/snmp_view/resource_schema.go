package snmp_view

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpViewResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP view resource.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Name of SNMP view. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP view. Minimum length =  1 Maximum length =  32",
			},
			"subtree": schema.StringAttribute{
				Required:            true,
				Description:         "Subtree associated with the SNMP view.A particular branch (subtree) of the MIB tree that you want to associate with this view.You must specify the subtree as an SNMP OID. Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "Subtree associated with the SNMP view.A particular branch (subtree) of the MIB tree that you want to associate with this view.You must specify the subtree as an SNMP OID. Minimum length =  1 Maximum length =  256",
			},
			"type": schema.BoolAttribute{
				Required:            true,
				Description:         "Include or Exclude the associated subtree . Values. true:Include, false: Exclude.",
				MarkdownDescription: "Include or Exclude the associated subtree . Values. true:Include, false: Exclude.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpViewModel struct {
	Name    types.String `tfsdk:"name"`
	Subtree types.String `tfsdk:"subtree"`
	Type    types.Bool   `tfsdk:"type"`
	Id      types.String `tfsdk:"id"`
}

func snmpViewGetThePayloadFromtheConfig(ctx context.Context, data *snmpViewModel) snmpViewReq {
	tflog.Debug(ctx, "In snmpViewGetThePayloadFromtheConfig Function")
	snmpViewReqPayload := snmpViewReq{
		Name:    data.Name.ValueString(),
		Subtree: data.Subtree.ValueString(),
		Type:    data.Type.ValueBoolPointer(),
	}
	return snmpViewReqPayload
}
func snmpViewSetAttrFromGet(ctx context.Context, data *snmpViewModel, getResponseData map[string]interface{}) *snmpViewModel {
	tflog.Debug(ctx, "In snmpViewSetAttrFromGet Function")
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.Subtree.IsNull() {
		data.Subtree = types.StringValue(getResponseData["subtree"].(string))
	}
	if !data.Type.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["type"].(string))
		data.Type = types.BoolValue(val)
	}
	return data
}

type snmpViewReq struct {
	Name    string `json:"name,omitempty"`
	Subtree string `json:"subtree,omitempty"`
	Type    *bool  `json:"type,omitempty"`
}
