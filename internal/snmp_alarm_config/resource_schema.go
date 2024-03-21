package snmp_alarm_config

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpAlarmConfigResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP Alarm Configurations resource.",
		Attributes: map[string]schema.Attribute{
			"enable": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable Alarm.",
				MarkdownDescription: "Enable Alarm.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Alarm Name. Maximum length =  128",
				MarkdownDescription: "Alarm Name. Maximum length =  128",
			},
			"severity": schema.StringAttribute{
				Optional:            true,
				Description:         "Alarm severity. Supported values: Critical, Major, Minor, Warning, Informational . Maximum length =  128",
				MarkdownDescription: "Alarm severity. Supported values: Critical, Major, Minor, Warning, Informational . Maximum length =  128",
			},
			"threshold": schema.Int64Attribute{
				Optional:            true,
				Description:         "Threshold Value for the alarm.",
				MarkdownDescription: "Threshold Value for the alarm.",
			},
			"time": schema.Int64Attribute{
				Optional:            true,
				Description:         "Frequency of the alarm in minutes.",
				MarkdownDescription: "Frequency of the alarm in minutes.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpAlarmConfigModel struct {
	Enable    types.Bool   `tfsdk:"enable"`
	Name      types.String `tfsdk:"name"`
	Severity  types.String `tfsdk:"severity"`
	Threshold types.Int64  `tfsdk:"threshold"`
	Time      types.Int64  `tfsdk:"time"`
	Id        types.String `tfsdk:"id"`
}

func snmpAlarmConfigGetThePayloadFromtheConfig(ctx context.Context, data *snmpAlarmConfigModel) snmpAlarmConfigReq {
	tflog.Debug(ctx, "In snmpAlarmConfigGetThePayloadFromtheConfig Function")
	snmpAlarmConfigReqPayload := snmpAlarmConfigReq{
		Enable:    data.Enable.ValueBoolPointer(),
		Name:      data.Name.ValueString(),
		Severity:  data.Severity.ValueString(),
		Threshold: data.Threshold.ValueInt64Pointer(),
		Time:      data.Time.ValueInt64Pointer(),
	}
	return snmpAlarmConfigReqPayload
}
func snmpAlarmConfigSetAttrFromGet(ctx context.Context, data *snmpAlarmConfigModel, getResponseData map[string]interface{}) *snmpAlarmConfigModel {
	tflog.Debug(ctx, "In snmpAlarmConfigSetAttrFromGet Function")
	if !data.Enable.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable"].(string))
		data.Enable = types.BoolValue(val)
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.Severity.IsNull() {
		data.Severity = types.StringValue(getResponseData["severity"].(string))
	}
	if !data.Threshold.IsNull() {
		val, _ := strconv.Atoi(getResponseData["threshold"].(string))
		data.Threshold = types.Int64Value(int64(val))
	}
	if !data.Time.IsNull() {
		val, _ := strconv.Atoi(getResponseData["time"].(string))
		data.Time = types.Int64Value(int64(val))
	}
	return data
}

type snmpAlarmConfigReq struct {
	Enable    *bool  `json:"enable,omitempty"`
	Name      string `json:"name,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Threshold *int64 `json:"threshold,omitempty"`
	Time      *int64 `json:"time,omitempty"`
}
