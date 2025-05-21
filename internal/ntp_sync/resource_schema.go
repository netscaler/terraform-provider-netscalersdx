package ntp_sync

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ntpSyncResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for NTP Sync configuration resource.",
		Attributes: map[string]schema.Attribute{
			"ntpd_status": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "ntpd status. true, to Enable NTP Sync. false, to Disable NTP Sync.",
				MarkdownDescription: "ntpd status. true, to Enable NTP Sync. false, to Disable NTP Sync.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type ntpSyncModel struct {
	NtpdStatus types.Bool   `tfsdk:"ntpd_status"`
	Id         types.String `tfsdk:"id"`
}

func ntpSyncGetThePayloadFromtheConfig(ctx context.Context, data *ntpSyncModel) ntpSyncReq {
	tflog.Debug(ctx, "In ntpSyncGetThePayloadFromtheConfig Function")
	ntpSyncReqPayload := ntpSyncReq{
		NtpdStatus: data.NtpdStatus.ValueBoolPointer(),
	}

	return ntpSyncReqPayload
}
func ntpSyncSetAttrFromGet(ctx context.Context, data *ntpSyncModel, getResponseData map[string]interface{}) *ntpSyncModel {
	tflog.Debug(ctx, "In ntpSyncSetAttrFromGet Function")

	data.NtpdStatus = types.BoolValue(utils.StringToBool(getResponseData["ntpd_status"].(string)))

	return data
}

type ntpSyncReq struct {
	NtpdStatus *bool `json:"ntpd_status,omitempty"`
}
