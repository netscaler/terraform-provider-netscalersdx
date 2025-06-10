package snmp_trap

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func snmpTrapDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a SNMP Trap.",
		Attributes: map[string]schema.Attribute{
			"community": schema.StringAttribute{
				Computed:            true,
				Description:         "Community Name. Maximum length =  32",
				MarkdownDescription: "Community Name. Maximum length =  32",
			},
			"dest_port": schema.Int64Attribute{
				Computed:            true,
				Description:         "Destination Port. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Destination Port. Minimum value =  1 Maximum value =  ",
			},
			"dest_server": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Trap Destination Server Address.",
				MarkdownDescription: "Trap Destination Server Address.",
			},
			"user_name": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Name of SNMP Trap User. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP Trap User. Minimum length =  1 Maximum length =  32",
			},
			"version": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP version. Maximum length =  2",
				MarkdownDescription: "SNMP version. Maximum length =  2",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the same as the dest_server value.",
			},
		},
	}
}
