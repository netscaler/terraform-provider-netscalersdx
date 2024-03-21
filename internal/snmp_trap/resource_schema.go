package snmp_trap

import (
	"context"
	"strconv"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpTrapResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP Trap Destinations resource.",
		Attributes: map[string]schema.Attribute{
			"community": schema.StringAttribute{
				Optional:            true,
				Description:         "Community Name. Maximum length =  32",
				MarkdownDescription: "Community Name. Maximum length =  32",
			},
			"dest_port": schema.Int64Attribute{
				Optional:            true,
				Description:         "Destination Port. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Destination Port. Minimum value =  1 Maximum value =  ",
			},
			"dest_server": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Trap Destination Server Address.",
				MarkdownDescription: "Trap Destination Server Address.",
			},
			"user_name": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Name of SNMP Trap User. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP Trap User. Minimum length =  1 Maximum length =  32",
			},
			"version": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP version. Maximum length =  2",
				MarkdownDescription: "SNMP version. Maximum length =  2",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpTrapModel struct {
	Community  types.String `tfsdk:"community"`
	DestPort   types.Int64  `tfsdk:"dest_port"`
	DestServer types.String `tfsdk:"dest_server"`
	UserName   types.List   `tfsdk:"user_name"`
	Version    types.String `tfsdk:"version"`
	Id         types.String `tfsdk:"id"`
}

func snmpTrapGetThePayloadFromtheConfig(ctx context.Context, data *snmpTrapModel) snmpTrapReq {
	tflog.Debug(ctx, "In snmpTrapGetThePayloadFromtheConfig Function")
	snmpTrapReqPayload := snmpTrapReq{
		Community:  data.Community.ValueString(),
		DestPort:   data.DestPort.ValueInt64(),
		DestServer: data.DestServer.ValueString(),
		UserName:   utils.TypeListToUnmarshalStringList(data.UserName),
		Version:    data.Version.ValueString(),
	}
	return snmpTrapReqPayload
}
func snmpTrapSetAttrFromGet(ctx context.Context, data *snmpTrapModel, getResponseData map[string]interface{}) *snmpTrapModel {
	tflog.Debug(ctx, "In snmpTrapSetAttrFromGet Function")
	if !data.Community.IsNull() && getResponseData["version"] == "v2" {
		data.Community = types.StringValue(getResponseData["community"].(string))
	}
	if !data.DestPort.IsNull() {
		val, _ := strconv.Atoi(getResponseData["dest_port"].(string))
		data.DestPort = types.Int64Value(int64(val))
	}
	if !data.DestServer.IsNull() {
		data.DestServer = types.StringValue(getResponseData["dest_server"].(string))
	}
	if !data.UserName.IsNull() && getResponseData["version"] == "v3" {
		data.UserName = utils.StringListToTypeList(utils.ToStringList(getResponseData["user_name"].([]interface{})))
	}
	if !data.Version.IsNull() {
		data.Version = types.StringValue(getResponseData["version"].(string))
	}
	return data
}

type snmpTrapReq struct {
	Community  string   `json:"community,omitempty"`
	DestPort   int64    `json:"dest_port,omitempty"`
	DestServer string   `json:"dest_server,omitempty"`
	UserName   []string `json:"user_name,omitempty"`
	Version    string   `json:"version,omitempty"`
}
