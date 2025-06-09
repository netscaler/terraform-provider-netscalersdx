package mps_feature

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func mpsFeatureDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing Mps feature.",
		Attributes: map[string]schema.Attribute{
			"admin_toggle": schema.Int64Attribute{
				Computed:            true,
				Description:         "This is Admin controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend.. Maximum value =  ",
				MarkdownDescription: "This is Admin controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend.. Maximum value =  ",
			},
			"built_in": schema.BoolAttribute{
				Computed:            true,
				Description:         "This is Ops controllable and will not be visible to the Admin to control. If true: Ops controllable feature, false: Admin controllable feature..",
				MarkdownDescription: "This is Ops controllable and will not be visible to the Admin to control. If true: Ops controllable feature, false: Admin controllable feature..",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				Description:         "Feature Description..",
				MarkdownDescription: "Feature Description..",
			},
			"feature_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Feature Name.",
				MarkdownDescription: "Feature Name.",
			},
			"ops_toggle": schema.Int64Attribute{
				Computed:            true,
				Description:         "This is Ops controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend. Ops controlled takes higher precedence than Admin Controlled.. Maximum value =  ",
				MarkdownDescription: "This is Ops controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend. Ops controlled takes higher precedence than Admin Controlled.. Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource is feature_name",
			},
		},
	}
}
