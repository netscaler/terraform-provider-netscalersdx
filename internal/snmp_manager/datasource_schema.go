package snmp_manager

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func snmpManagerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a SNMP Manager.",
		Attributes: map[string]schema.Attribute{
			"community": schema.StringAttribute{
				Computed:            true,
				Description:         "Community Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Community Name. Minimum length =  1 Maximum length =  128",
			},
			"ip_address": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Manager IPAddress.",
				MarkdownDescription: "Manager IPAddress.",
			},
			"netmask": schema.StringAttribute{
				Computed:            true,
				Description:         "Netmask. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Netmask. Minimum length =  1 Maximum length =  64",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource, it is same as the ip_address.",
			},
		},
	}
}
