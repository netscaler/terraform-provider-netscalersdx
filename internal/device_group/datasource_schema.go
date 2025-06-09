package device_group

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func deviceGroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing Device Group of NetScaler SDX instance",
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				Computed:            true,
				Description:         "Device group category. Will be default/upgrade.. Maximum length =  255",
				MarkdownDescription: "Device group category. Will be default/upgrade.. Maximum length =  255",
			},
			"criteria_condn": schema.StringAttribute{
				Computed:            true,
				Description:         "Tenant. Maximum length =  255",
				MarkdownDescription: "Tenant. Maximum length =  255",
			},
			"criteria_type": schema.StringAttribute{
				Computed:            true,
				Description:         "Device Group Criteria. Maximum length =  255",
				MarkdownDescription: "Device Group Criteria. Maximum length =  255",
			},
			"criteria_value": schema.StringAttribute{
				Computed:            true,
				Description:         "Criteria Value. Maximum length =  255",
				MarkdownDescription: "Criteria Value. Maximum length =  255",
			},
			"device_family": schema.StringAttribute{
				Computed:            true,
				Description:         "Device Family. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Device Family. Minimum length =  1 Maximum length =  64",
			},
			"disable_upgrade": schema.BoolAttribute{
				Computed:            true,
				Description:         "Setting to disable agent upgrades.",
				MarkdownDescription: "Setting to disable agent upgrades.",
			},
			"duration": schema.Int64Attribute{
				Computed:            true,
				Description:         "Duration of Maintenance window for groups of category upgrade. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Duration of Maintenance window for groups of category upgrade. Minimum value =  1 Maximum value =  ",
			},
			"lock_acquire_time": schema.StringAttribute{
				Computed:            true,
				Description:         "Upgrade Lock acquiring time. Maximum length =  255",
				MarkdownDescription: "Upgrade Lock acquiring time. Maximum length =  255",
			},
			"lock_acquiring_device": schema.StringAttribute{
				Computed:            true,
				Description:         "Upgrade Lock acquiring device. Maximum length =  255",
				MarkdownDescription: "Upgrade Lock acquiring device. Maximum length =  255",
			},
			"maintenance_window_start": schema.StringAttribute{
				Computed:            true,
				Description:         "Maintenance window start time for groups of category upgrade. Maximum length =  255",
				MarkdownDescription: "Maintenance window start time for groups of category upgrade. Maximum length =  255",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Device Group Name. Maximum length =  255",
				MarkdownDescription: "Device Group Name. Maximum length =  255",
			},
			"static_device_list": schema.StringAttribute{
				Computed:            true,
				Description:         "Devices in the group.",
				MarkdownDescription: "Devices in the group.",
			},
			"static_device_list_arr": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Static Device List.",
				MarkdownDescription: "Static Device List.",
			},
			"upgrade_lock": schema.BoolAttribute{
				Computed:            true,
				Description:         "Lock to be acquired before upgrading device group.",
				MarkdownDescription: "Lock to be acquired before upgrading device group.",
			},
			"upgrade_version": schema.StringAttribute{
				Computed:            true,
				Description:         "New Available upgrade version for devices. Maximum length =  255",
				MarkdownDescription: "New Available upgrade version for devices. Maximum length =  255",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Description: "The ID of this resource",
			},
		},
	}
}
