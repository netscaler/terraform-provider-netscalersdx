package ntp_param

import (
	"context"

	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ntpParamResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for NTP Parameters resource.",
		Attributes: map[string]schema.Attribute{
			"authentication": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Authentication Enabled.",
				MarkdownDescription: "Authentication Enabled.",
			},
			"automax_logsec": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description:         "Automax Interval (as power of 2 in seconds).",
				MarkdownDescription: "Automax Interval (as power of 2 in seconds).",
			},
			"revoke_logsec": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description:         "Revoke Interval (as power of 2 in seconds).",
				MarkdownDescription: "Revoke Interval (as power of 2 in seconds).",
			},
			"trusted_key_list": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				// We have below code insted of ForceNew
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description:         "List of Trusted Key Identifiers for Symmetric Key Cryptography. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "List of Trusted Key Identifiers for Symmetric Key Cryptography. Minimum value =  1 Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type ntpParamModel struct {
	Authentication types.Bool   `tfsdk:"authentication"`
	AutomaxLogsec  types.Int64  `tfsdk:"automax_logsec"`
	RevokeLogsec   types.Int64  `tfsdk:"revoke_logsec"`
	TrustedKeyList types.List   `tfsdk:"trusted_key_list"`
	Id             types.String `tfsdk:"id"`
}

func ntpParamGetThePayloadFromtheConfig(ctx context.Context, data *ntpParamModel) ntpParamReq {
	tflog.Debug(ctx, "In ntpParamGetThePayloadFromtheConfig Function")
	ntpParamReqPayload := ntpParamReq{
		Authentication: data.Authentication.ValueBool(),
		AutomaxLogsec:  data.AutomaxLogsec.ValueInt64(),
		RevokeLogsec:   data.RevokeLogsec.ValueInt64(),
		TrustedKeyList: utils.TypeListToStringList(data.TrustedKeyList),
	}
	return ntpParamReqPayload
}

type ntpParamReq struct {
	Authentication bool     `json:"authentication"`
	AutomaxLogsec  int64    `json:"automax_logsec,omitempty"`
	RevokeLogsec   int64    `json:"revoke_logsec,omitempty"`
	TrustedKeyList []string `json:"trusted_key_list,omitempty"`
}
