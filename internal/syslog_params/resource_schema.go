package syslog_params

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func syslogParamsResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Syslog Parameters resource.",
		Attributes: map[string]schema.Attribute{
			"date_format": schema.StringAttribute{
				Required: true,
				// We have below code insted of ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Format of date to be added in the syslog message.",
				MarkdownDescription: "Format of date to be added in the syslog message.",
			},
			"timezone": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Timezone to be used in the syslog message.",
				MarkdownDescription: "Timezone to be used in the syslog message.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type syslogParamsModel struct {
	DateFormat types.String `tfsdk:"date_format"`
	Timezone   types.String `tfsdk:"timezone"`
	Id         types.String `tfsdk:"id"`
}

func syslogParamsGetThePayloadFromtheConfig(ctx context.Context, data *syslogParamsModel) syslogParamsReq {
	tflog.Debug(ctx, "In syslogParamsGetThePayloadFromtheConfig Function")
	syslogParamsReqPayload := syslogParamsReq{
		DateFormat: data.DateFormat.ValueString(),
		Timezone:   data.Timezone.ValueString(),
	}
	return syslogParamsReqPayload
}
func syslogParamsSetAttrFromGet(ctx context.Context, data *syslogParamsModel, getResponseData map[string]interface{}) *syslogParamsModel {
	tflog.Debug(ctx, "In syslogParamsSetAttrFromGet Function")

	data.DateFormat = types.StringValue(getResponseData["date_format"].(string))
	data.Timezone = types.StringValue(getResponseData["timezone"].(string))

	return data
}

type syslogParamsReq struct {
	DateFormat string `json:"date_format,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}
