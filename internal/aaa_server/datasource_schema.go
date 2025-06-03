package aaa_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func aaaServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Retrieve an AAA server by its ID.",
		Attributes: map[string]schema.Attribute{
			"external_servers": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"external_server_name": schema.StringAttribute{
							Required:            true,
							Description:         "Name of external server. Minimum length =  1 Maximum length =  128",
							MarkdownDescription: "Name of external server. Minimum length =  1 Maximum length =  128",
						},
						"external_server_type": schema.StringAttribute{
							Required:            true,
							Description:         "Type of external server. Supported types 1.RADIUS 2.LDAP 3.TACACS 4.KEYSTONE. Minimum length =  1 Maximum length =  32",
							MarkdownDescription: "Type of external server. Supported types 1.RADIUS 2.LDAP 3.TACACS 4.KEYSTONE. Minimum length =  1 Maximum length =  32",
						},
						"priority": schema.Int64Attribute{
							Required:            true,
							Description:         "Priority of external server. Minimum value =  2 Maximum value =  ",
							MarkdownDescription: "Priority of external server. Minimum value =  2 Maximum value =  ",
						},
					},
				},
				Optional:            true,
				Computed:            false,
				Description:         "List of external servers.",
				MarkdownDescription: "List of external servers.",
			},
			"fallback_local_authentication": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable local fallback authentication.",
				MarkdownDescription: "Enable local fallback authentication.",
			},
			"log_ext_group_info": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Log external group info.",
				MarkdownDescription: "Log external group info.",
			},
			"primary_server_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of primary server name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of primary server name. Minimum length =  1 Maximum length =  128",
			},
			"primary_server_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Type of primary server. Supported types 1. LOCAL 2.RADIUS 3.LDAP 4.TACACS 5.KEYSTONE. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Type of primary server. Supported types 1. LOCAL 2.RADIUS 3.LDAP 4.TACACS 5.KEYSTONE. Minimum length =  1 Maximum length =  32",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}
