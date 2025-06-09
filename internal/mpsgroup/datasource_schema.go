package mpsgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mpsgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing mpsgroup",
		Attributes: map[string]schema.Attribute{
			"allow_application_only": schema.BoolAttribute{
				Computed:            true,
				Description:         "Checks if only application centic page is needed.",
				MarkdownDescription: "Checks if only application centic page is needed.",
			},
			"application_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "All Application names that are part of this group.This includes selected appnames as well as applications which are result of defined regex.",
				MarkdownDescription: "All Application names that are part of this group.This includes selected appnames as well as applications which are result of defined regex.",
			},
			"apply_all_bound_entities": schema.BoolAttribute{
				Computed:            true,
				Description:         "Apply for all bound entities (TRUE|FALSE).",
				MarkdownDescription: "Apply for all bound entities (TRUE|FALSE).",
			},
			"assign_all_apps": schema.BoolAttribute{
				Computed:            true,
				Description:         "Assign All Applications (YES|NO).",
				MarkdownDescription: "Assign All Applications (YES|NO).",
			},
			"assign_all_autoscale_groups": schema.BoolAttribute{
				Computed:            true,
				Description:         "Assign All Autoscale groups (YES|NO).",
				MarkdownDescription: "Assign All Autoscale groups (YES|NO).",
			},
			"assign_all_devices": schema.BoolAttribute{
				Computed:            true,
				Description:         "Assign All Instances (YES|NO).",
				MarkdownDescription: "Assign All Instances (YES|NO).",
			},
			"assign_all_selected_device_apps": schema.BoolAttribute{
				Computed:            true,
				Description:         "Assign All Application from selected instances (YES|NO).",
				MarkdownDescription: "Assign All Application from selected instances (YES|NO).",
			},
			"autoscale_groups_id": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Autoscale groups belong to this groupp.",
				MarkdownDescription: "Autoscale groups belong to this groupp.",
			},
			"bound_entity_selected": schema.Int64Attribute{
				Computed:            true,
				Description:         "Which bound entiy is selected VSERVER(0),SERVICE(1),SERVICEGROUP(2),SERVER(3).",
				MarkdownDescription: "Which bound entiy is selected VSERVER(0),SERVICE(1),SERVICEGROUP(2),SERVER(3).",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				Description:         "Description of Group. Minimum length =  1 Maximum length =  1024",
				MarkdownDescription: "Description of Group. Minimum length =  1 Maximum length =  1024",
			},
			"enable_session_timeout": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enables session timeout for group.",
				MarkdownDescription: "Enables session timeout for group.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Group Name. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Group Name. Minimum length =  1 Maximum length =  64",
			},
			"permission": schema.StringAttribute{
				Computed:            true,
				Description:         "Permission for the group (admin/read-only). Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Permission for the group (admin/read-only). Minimum length =  1 Maximum length =  128",
			},
			"role": schema.StringAttribute{
				Computed:            true,
				Description:         "Role (admin|nonadmin).",
				MarkdownDescription: "Role (admin|nonadmin).",
			},
			"roles": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Roles assigned to the group.",
				MarkdownDescription: "Roles assigned to the group.",
			},
			"select_individual_entity": schema.BoolAttribute{
				Computed:            true,
				Description:         "Select Individual Entity Type.",
				MarkdownDescription: "Select Individual Entity Type.",
			},
			"session_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "Session timeout for the Group.",
				MarkdownDescription: "Session timeout for the Group.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Computed:            true,
				Description:         "Session timeout unit for the Group.",
				MarkdownDescription: "Session timeout unit for the Group.",
			},
			"standalone_instances_id": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Stand alone instances belong to this groupp.",
				MarkdownDescription: "Stand alone instances belong to this groupp.",
			},
			"tenant_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Id of the tenant. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Id of the tenant. Minimum length =  1 Maximum length =  128",
			},
			"users": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Users belong to the group.",
				MarkdownDescription: "Users belong to the group.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource",
			},
		},
	}
}
