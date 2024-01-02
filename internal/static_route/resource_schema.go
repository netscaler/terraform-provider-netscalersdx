package static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func staticRouteResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Static Route resource.",
		Attributes: map[string]schema.Attribute{
			"gateway": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Gateway for route addition.",
				MarkdownDescription: "Gateway for route addition.",
			},
			"netmask": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "netmask for route addition.",
				MarkdownDescription: "netmask for route addition.",
			},
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "network for route addition.",
				MarkdownDescription: "network for route addition.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type staticRouteModel struct {
	Gateway types.String `tfsdk:"gateway"`
	Netmask types.String `tfsdk:"netmask"`
	Network types.String `tfsdk:"network"`
	Id      types.String `tfsdk:"id"`
}

func staticRouteGetThePayloadFromtheConfig(ctx context.Context, data *staticRouteModel) staticRouteReq {
	tflog.Debug(ctx, "In staticRouteGetThePayloadFromtheConfig Function")
	staticRouteReqPayload := staticRouteReq{
		Gateway: data.Gateway.ValueString(),
		Netmask: data.Netmask.ValueString(),
		Network: data.Network.ValueString(),
	}
	return staticRouteReqPayload
}

type staticRouteReq struct {
	Gateway string `json:"gateway,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}
