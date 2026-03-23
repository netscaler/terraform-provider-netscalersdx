package nslaslicense_offline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func nslaslicenseOfflineResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "NetScaler SDX LAS Offline License resource. This resource generates and applies offline LAS licenses for NetScaler SDX appliances.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource identifier (device_ip from provider)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"las_secrets_json": schema.StringAttribute{
				MarkdownDescription: "Path to JSON file containing LAS credentials (ccid, client, password, las_endpoint, cc_endpoint)",
				Required:            true,
				Sensitive:           true,
			},
			"entitlement_name": schema.StringAttribute{
				MarkdownDescription: "Entitlement name for the license (e.g., 'SDX 9195 Premium')",
				Required:            true,
			},
			"lsguid": schema.StringAttribute{
				MarkdownDescription: "License Server GUID (computed from device)",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "SDX software version (computed)",
				Computed:            true,
			},
			"build": schema.StringAttribute{
				MarkdownDescription: "SDX build number (computed)",
				Computed:            true,
			},
			"license_blob_path": schema.StringAttribute{
				MarkdownDescription: "Path where license blob is saved locally",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "License application status",
				Computed:            true,
			},
			"last_updated": schema.StringAttribute{
				MarkdownDescription: "Timestamp of last update",
				Computed:            true,
			},
		},
	}
}

type nslaslicenseOfflineModel struct {
	Id              types.String `tfsdk:"id"`
	LASSecretsJson  types.String `tfsdk:"las_secrets_json"`
	EntitlementName types.String `tfsdk:"entitlement_name"`
	LSGUID          types.String `tfsdk:"lsguid"`
	Version         types.String `tfsdk:"version"`
	Build           types.String `tfsdk:"build"`
	LicenseBlob     types.String `tfsdk:"license_blob_path"`
	Status          types.String `tfsdk:"status"`
	LastUpdated     types.String `tfsdk:"last_updated"`
}
