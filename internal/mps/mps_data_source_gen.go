package mps

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func mpsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing Mps entity.",
		Attributes: map[string]schema.Attribute{
			"config_motd": schema.BoolAttribute{
				Computed:            true,
				Description:         "Configure Message of the Day (contents of motd file), this needs to be set true if user wants to configure message if the day.",
				MarkdownDescription: "Configure Message of the Day (contents of motd file), this needs to be set true if user wants to configure message if the day.",
			},
			"deployment_type": schema.StringAttribute{
				Computed:            true,
				Description:         "Indicates the type of deployment of NetScaler ADM: standalone or ha or scaleout..",
				MarkdownDescription: "Indicates the type of deployment of NetScaler ADM: standalone or ha or scaleout..",
			},
			"hist_mig_inprog": schema.BoolAttribute{
				Computed:            true,
				Description:         "Indicates whether the historical tables database migration is in progress or not..",
				MarkdownDescription: "Indicates whether the historical tables database migration is in progress or not..",
			},
			"is_cloud": schema.BoolAttribute{
				Computed:            true,
				Description:         "True if its a cloud deployment.",
				MarkdownDescription: "True if its a cloud deployment.",
			},
			"is_container": schema.BoolAttribute{
				Computed:            true,
				Description:         "True if its a container deployment.",
				MarkdownDescription: "True if its a container deployment.",
			},
			"is_member_of_default_group": schema.BoolAttribute{
				Computed:            true,
				Description:         "Is Member Of Default Group.",
				MarkdownDescription: "Is Member Of Default Group.",
			},
			"is_passive": schema.BoolAttribute{
				Computed:            true,
				Description:         "Indicates the node's state: ACTIVE or PASSIVE in a HA deployment..",
				MarkdownDescription: "Indicates the node's state: ACTIVE or PASSIVE in a HA deployment..",
			},
			"is_thirdparty_vm_supported": schema.BoolAttribute{
				Computed:            true,
				Description:         "True, if third party VM is supported.",
				MarkdownDescription: "True, if third party VM is supported.",
			},
			"maps_apikey": schema.StringAttribute{
				Computed:            true,
				Description:         "Maps API Key.",
				MarkdownDescription: "Maps API Key.",
			},
			"motd": schema.StringAttribute{
				Computed:            true,
				Description:         "Message of the Day (contents of motd file) that can be shown on UI after successful login. Maximum length =  255",
				MarkdownDescription: "Message of the Day (contents of motd file) that can be shown on UI after successful login. Maximum length =  255",
			},
			"privilege_feature": schema.StringAttribute{
				Computed:            true,
				Description:         "privilege_feature.",
				MarkdownDescription: "privilege_feature.",
			},
			"upgrade_agent_version": schema.StringAttribute{
				Computed:            true,
				Description:         "Indicates the next version the agent needs to upgrade to..",
				MarkdownDescription: "Indicates the next version the agent needs to upgrade to..",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource.",
			},
		},
	}
}
