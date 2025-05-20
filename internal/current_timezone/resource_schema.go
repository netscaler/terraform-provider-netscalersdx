package current_timezone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func currentTimezoneResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for Current timezone resource.",
		Attributes: map[string]schema.Attribute{
			"timezone": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Timezone.",
				MarkdownDescription: "Timezone.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type currentTimezoneModel struct {
	Timezone types.String `tfsdk:"timezone"`
	Id       types.String `tfsdk:"id"`
}

func currentTimezoneGetThePayloadFromtheConfig(ctx context.Context, data *currentTimezoneModel) currentTimezoneReq {
	tflog.Debug(ctx, "In currentTimezoneGetThePayloadFromtheConfig Function")
	currentTimezoneReqPayload := currentTimezoneReq{
		Timezone: data.Timezone.ValueString(),
	}
	return currentTimezoneReqPayload
}
func currentTimezoneSetAttrFromGet(ctx context.Context, data *currentTimezoneModel, getResponseData map[string]interface{}) *currentTimezoneModel {
	tflog.Debug(ctx, "In currentTimezoneSetAttrFromGet Function")

	data.Timezone = types.StringValue(getResponseData["timezone"].(string))

	return data
}

type currentTimezoneReq struct {
	Timezone string `json:"timezone,omitempty"`
}
