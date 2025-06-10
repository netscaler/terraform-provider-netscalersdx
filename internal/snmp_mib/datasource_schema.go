package snmp_mib

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func snmpMibDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a SNMP Mib.",
		Attributes: map[string]schema.Attribute{
			"contact": schema.StringAttribute{
				Computed:            true,
				Description:         "Name of the administrator for appliance.. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Name of the administrator for appliance.. Minimum length =  1 Maximum length =  127",
			},
			"custom_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Custom identification number for appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Custom identification number for appliance. Minimum length =  1 Maximum length =  127",
			},
			"location": schema.StringAttribute{
				Computed:            true,
				Description:         "Physical location of appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Physical location of appliance. Minimum length =  1 Maximum length =  127",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name for appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Name for appliance. Minimum length =  1 Maximum length =  127",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource, it is same as name attribute.",
			},
		},
	}
}
