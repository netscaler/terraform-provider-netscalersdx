package snmp_view

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func snmpViewDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a SNMP View.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Name of SNMP view. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP view. Minimum length =  1 Maximum length =  32",
			},
			"subtree": schema.StringAttribute{
				Computed:            true,
				Description:         "Subtree associated with the SNMP view.A particular branch (subtree) of the MIB tree that you want to associate with this view.You must specify the subtree as an SNMP OID. Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "Subtree associated with the SNMP view.A particular branch (subtree) of the MIB tree that you want to associate with this view.You must specify the subtree as an SNMP OID. Minimum length =  1 Maximum length =  256",
			},
			"type": schema.BoolAttribute{
				Computed:            true,
				Description:         "Include or Exclude the associated subtree . Values. true:Include, false: Exclude.",
				MarkdownDescription: "Include or Exclude the associated subtree . Values. true:Include, false: Exclude.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the same as the name value.",
			},
		},
	}
}
