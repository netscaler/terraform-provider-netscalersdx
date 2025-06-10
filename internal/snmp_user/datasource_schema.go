package snmp_user

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func snmpUserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a SNMP User.",
		Attributes: map[string]schema.Attribute{
			"auth_password": schema.StringAttribute{
				Computed:            true,
				Description:         "Authentication Password of SNMP User. Minimum length =  8 Maximum length =  32",
				MarkdownDescription: "Authentication Password of SNMP User. Minimum length =  8 Maximum length =  32",
			},
			"auth_protocol": schema.Int64Attribute{
				Computed:            true,
				Description:         "Authentication Protocol of SNMP User. Values: 0:noValue, 1: MD5, 2: SHA1. Maximum value =  ",
				MarkdownDescription: "Authentication Protocol of SNMP User. Values: 0:noValue, 1: MD5, 2: SHA1. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Name of SNMP User. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP User. Minimum length =  1 Maximum length =  32",
			},
			"privacy_password": schema.StringAttribute{
				Computed:            true,
				Description:         "Privacy Password of SNMP User. Minimum length =  8 Maximum length =  32",
				MarkdownDescription: "Privacy Password of SNMP User. Minimum length =  8 Maximum length =  32",
			},
			"privacy_protocol": schema.Int64Attribute{
				Computed:            true,
				Description:         "Privacy Protocol of SNMP User. Values: 0:noValue, 1: DES, 2: AES. Maximum value =  ",
				MarkdownDescription: "Privacy Protocol of SNMP User. Values: 0:noValue, 1: DES, 2: AES. Maximum value =  ",
			},
			"security_level": schema.Int64Attribute{
				Computed:            true,
				Description:         "Security Level of SNMP User. Values: 0: noAuthNoPriv, 1: authNoPriv, 2: authPriv. Maximum value =  ",
				MarkdownDescription: "Security Level of SNMP User. Values: 0: noAuthNoPriv, 1: authNoPriv, 2: authPriv. Maximum value =  ",
			},
			"view_name": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP View Name attached to the SNMP User. Maximum length =  32",
				MarkdownDescription: "SNMP View Name attached to the SNMP User. Maximum length =  32",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the same as the name value.",
			},
		},
	}
}
