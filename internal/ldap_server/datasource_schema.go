package ldap_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ldapServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing Ldap Server",
		Attributes: map[string]schema.Attribute{
			"auth_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the LDAP server.",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the LDAP server.",
			},
			"authentication": schema.BoolAttribute{
				Computed:            true,
				Description:         "Set to false for public key Authentication.",
				MarkdownDescription: "Set to false for public key Authentication.",
			},
			"base_dn": schema.StringAttribute{
				Computed:            true,
				Description:         "The base or node where the ldapsearch should start. Maximum length =  128",
				MarkdownDescription: "The base or node where the ldapsearch should start. Maximum length =  128",
			},
			"bind_dn": schema.StringAttribute{
				Computed:            true,
				Description:         "The full distinguished name used to bind to the LDAP server. Maximum length =  128",
				MarkdownDescription: "The full distinguished name used to bind to the LDAP server. Maximum length =  128",
			},
			"bind_passwd": schema.StringAttribute{
				Computed:            true,
				Description:         "The password used to bind to the LDAP server. Maximum length =  128",
				MarkdownDescription: "The password used to bind to the LDAP server. Maximum length =  128",
			},
			"change_password": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable change of the user.",
				MarkdownDescription: "Enable change of the user.",
			},
			"default_authentication_group": schema.StringAttribute{
				Computed:            true,
				Description:         "This is the default group. Maximum length =  64",
				MarkdownDescription: "This is the default group. Maximum length =  64",
			},
			"follow_referrals": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable following LDAP referrals received from LDAP server.",
				MarkdownDescription: "Enable following LDAP referrals received from LDAP server.",
			},
			"group_attr_name": schema.StringAttribute{
				Computed:            true,
				Description:         "The Attribute name for group extraction from the LDAP server. Maximum length =  32",
				MarkdownDescription: "The Attribute name for group extraction from the LDAP server. Maximum length =  32",
			},
			"group_name_identifier": schema.StringAttribute{
				Computed:            true,
				Description:         "Name that uniquely identifies a group in LDAP server. Maximum length =  32",
				MarkdownDescription: "Name that uniquely identifies a group in LDAP server. Maximum length =  32",
			},
			"group_search_attribute": schema.StringAttribute{
				Computed:            true,
				Description:         "LDAP group search attribute. Used to determine to which groups a group belongs. Maximum length =  32",
				MarkdownDescription: "LDAP group search attribute. Used to determine to which groups a group belongs. Maximum length =  32",
			},
			"group_search_filter": schema.StringAttribute{
				Computed:            true,
				Description:         "String to be combined with the default LDAP group search string to form the search value. Maximum length =  128",
				MarkdownDescription: "String to be combined with the default LDAP group search string to form the search value. Maximum length =  128",
			},
			"group_search_subattribute": schema.StringAttribute{
				Computed:            true,
				Description:         "LDAP group search subattribute. Used to determine to which groups a group belongs.. Maximum length =  32",
				MarkdownDescription: "LDAP group search subattribute. Used to determine to which groups a group belongs.. Maximum length =  32",
			},
			"ip_address": schema.StringAttribute{
				Computed:            true,
				Description:         "The IP address of the LDAP server.. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "The IP address of the LDAP server.. Minimum length =  1 Maximum length =  64",
			},
			"ldap_host_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Host Name on the certificate from LDAP Server. Maximum length =  128",
				MarkdownDescription: "Host Name on the certificate from LDAP Server. Maximum length =  128",
			},
			"login_name": schema.StringAttribute{
				Computed:            true,
				Description:         "The name attribute used by the system to query the external LDAP server. Maximum length =  32",
				MarkdownDescription: "The name attribute used by the system to query the external LDAP server. Maximum length =  32",
			},
			"max_ldap_referrals": schema.Int64Attribute{
				Computed:            true,
				Description:         "Maximum number of ldap referrals to follow.",
				MarkdownDescription: "Maximum number of ldap referrals to follow.",
			},
			"max_nesting_level": schema.Int64Attribute{
				Computed:            true,
				Description:         "Number of levels at which group extraction is allowed. Maximum value =  ",
				MarkdownDescription: "Number of levels at which group extraction is allowed. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Name of LDAP server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of LDAP server. Minimum length =  1 Maximum length =  128",
			},
			"nested_group_extraction": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable Nested Group Extraction.",
				MarkdownDescription: "Enable Nested Group Extraction.",
			},
			"port": schema.Int64Attribute{
				Computed:            true,
				Description:         "The port number on which the LDAP server is running. Maximum value =  ",
				MarkdownDescription: "The port number on which the LDAP server is running. Maximum value =  ",
			},
			"search_filter": schema.StringAttribute{
				Computed:            true,
				Description:         "The String to be combined with the default LDAP user search string to form the value. Maximum length =  256",
				MarkdownDescription: "The String to be combined with the default LDAP user search string to form the value. Maximum length =  256",
			},
			"sec_type": schema.StringAttribute{
				Computed:            true,
				Description:         "The communication type between the system and the LDAP server.",
				MarkdownDescription: "The communication type between the system and the LDAP server.",
			},
			"ssh_public_key": schema.StringAttribute{
				Computed:            true,
				Description:         "SSH public key attribute holds the public keys of the user. Maximum length =  64",
				MarkdownDescription: "SSH public key attribute holds the public keys of the user. Maximum length =  64",
			},
			"subattribute_name": schema.StringAttribute{
				Computed:            true,
				Description:         "The Sub-Attribute name for group extraction from LDAP server. Maximum length =  32",
				MarkdownDescription: "The Sub-Attribute name for group extraction from LDAP server. Maximum length =  32",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				Description:         "The type of LDAP server. Minimum length =  2 Maximum length =  64",
				MarkdownDescription: "The type of LDAP server. Minimum length =  2 Maximum length =  64",
			},
			"validate_ldap_server_certs": schema.BoolAttribute{
				Computed:            true,
				Description:         "Validate LDAP Server Certificate.",
				MarkdownDescription: "Validate LDAP Server Certificate.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource",
			},
		},
	}
}
