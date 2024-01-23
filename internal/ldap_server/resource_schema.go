package ldap_server

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ldapServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for LDAP Server resource.",
		Attributes: map[string]schema.Attribute{
			"auth_timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the LDAP server.",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the LDAP server.",
			},
			"authentication": schema.BoolAttribute{
				Optional:            true,
				Description:         "Set to false for public key Authentication.",
				MarkdownDescription: "Set to false for public key Authentication.",
			},
			"base_dn": schema.StringAttribute{
				Optional:            true,
				Description:         "The base or node where the ldapsearch should start. Maximum length =  128",
				MarkdownDescription: "The base or node where the ldapsearch should start. Maximum length =  128",
			},
			"bind_dn": schema.StringAttribute{
				Optional:            true,
				Description:         "The full distinguished name used to bind to the LDAP server. Maximum length =  128",
				MarkdownDescription: "The full distinguished name used to bind to the LDAP server. Maximum length =  128",
			},
			"bind_passwd": schema.StringAttribute{
				Optional:            true,
				Description:         "The password used to bind to the LDAP server. Maximum length =  128",
				MarkdownDescription: "The password used to bind to the LDAP server. Maximum length =  128",
			},
			"change_password": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable change of the user.",
				MarkdownDescription: "Enable change of the user.",
			},
			"default_authentication_group": schema.StringAttribute{
				Optional:            true,
				Description:         "This is the default group. Maximum length =  64",
				MarkdownDescription: "This is the default group. Maximum length =  64",
			},
			"follow_referrals": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable following LDAP referrals received from LDAP server.",
				MarkdownDescription: "Enable following LDAP referrals received from LDAP server.",
			},
			"group_attr_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The Attribute name for group extraction from the LDAP server. Maximum length =  32",
				MarkdownDescription: "The Attribute name for group extraction from the LDAP server. Maximum length =  32",
			},
			"group_name_identifier": schema.StringAttribute{
				Optional:            true,
				Description:         "Name that uniquely identifies a group in LDAP server. Maximum length =  32",
				MarkdownDescription: "Name that uniquely identifies a group in LDAP server. Maximum length =  32",
			},
			"group_search_attribute": schema.StringAttribute{
				Optional:            true,
				Description:         "LDAP group search attribute. Used to determine to which groups a group belongs. Maximum length =  32",
				MarkdownDescription: "LDAP group search attribute. Used to determine to which groups a group belongs. Maximum length =  32",
			},
			"group_search_filter": schema.StringAttribute{
				Optional:            true,
				Description:         "String to be combined with the default LDAP group search string to form the search value. Maximum length =  128",
				MarkdownDescription: "String to be combined with the default LDAP group search string to form the search value. Maximum length =  128",
			},
			"group_search_subattribute": schema.StringAttribute{
				Optional:            true,
				Description:         "LDAP group search subattribute. Used to determine to which groups a group belongs.. Maximum length =  32",
				MarkdownDescription: "LDAP group search subattribute. Used to determine to which groups a group belongs.. Maximum length =  32",
			},
			"ip_address": schema.StringAttribute{
				Required:            true,
				Description:         "The IP address of the LDAP server.. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "The IP address of the LDAP server.. Minimum length =  1 Maximum length =  64",
			},
			"ldap_host_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Host Name on the certificate from LDAP Server. Maximum length =  128",
				MarkdownDescription: "Host Name on the certificate from LDAP Server. Maximum length =  128",
			},
			"login_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The name attribute used by the system to query the external LDAP server. Maximum length =  32",
				MarkdownDescription: "The name attribute used by the system to query the external LDAP server. Maximum length =  32",
			},
			"max_ldap_referrals": schema.Int64Attribute{
				Optional:            true,
				Description:         "Maximum number of ldap referrals to follow.",
				MarkdownDescription: "Maximum number of ldap referrals to follow.",
			},
			"max_nesting_level": schema.Int64Attribute{
				Optional:            true,
				Description:         "Number of levels at which group extraction is allowed. Maximum value =  ",
				MarkdownDescription: "Number of levels at which group extraction is allowed. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of LDAP server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of LDAP server. Minimum length =  1 Maximum length =  128",
			},
			"nested_group_extraction": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enable Nested Group Extraction.",
				MarkdownDescription: "Enable Nested Group Extraction.",
			},
			"port": schema.Int64Attribute{
				Optional:            true,
				Description:         "The port number on which the LDAP server is running. Maximum value =  ",
				MarkdownDescription: "The port number on which the LDAP server is running. Maximum value =  ",
			},
			"search_filter": schema.StringAttribute{
				Optional:            true,
				Description:         "The String to be combined with the default LDAP user search string to form the value. Maximum length =  256",
				MarkdownDescription: "The String to be combined with the default LDAP user search string to form the value. Maximum length =  256",
			},
			"sec_type": schema.StringAttribute{
				Optional:            true,
				Description:         "The communication type between the system and the LDAP server.",
				MarkdownDescription: "The communication type between the system and the LDAP server.",
			},
			"ssh_public_key": schema.StringAttribute{
				Optional:            true,
				Description:         "SSH public key attribute holds the public keys of the user. Maximum length =  64",
				MarkdownDescription: "SSH public key attribute holds the public keys of the user. Maximum length =  64",
			},
			"subattribute_name": schema.StringAttribute{
				Optional:            true,
				Description:         "The Sub-Attribute name for group extraction from LDAP server. Maximum length =  32",
				MarkdownDescription: "The Sub-Attribute name for group extraction from LDAP server. Maximum length =  32",
			},
			"type": schema.StringAttribute{
				Required:            true,
				Description:         "The type of LDAP server. Minimum length =  2 Maximum length =  64",
				MarkdownDescription: "The type of LDAP server. Minimum length =  2 Maximum length =  64",
			},
			"validate_ldap_server_certs": schema.BoolAttribute{
				Optional:            true,
				Description:         "Validate LDAP Server Certificate.",
				MarkdownDescription: "Validate LDAP Server Certificate.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type ldapServerModel struct {
	AuthTimeout                types.Int64  `tfsdk:"auth_timeout"`
	Authentication             types.Bool   `tfsdk:"authentication"`
	BaseDn                     types.String `tfsdk:"base_dn"`
	BindDn                     types.String `tfsdk:"bind_dn"`
	BindPasswd                 types.String `tfsdk:"bind_passwd"`
	ChangePassword             types.Bool   `tfsdk:"change_password"`
	DefaultAuthenticationGroup types.String `tfsdk:"default_authentication_group"`
	FollowReferrals            types.Bool   `tfsdk:"follow_referrals"`
	GroupAttrName              types.String `tfsdk:"group_attr_name"`
	GroupNameIdentifier        types.String `tfsdk:"group_name_identifier"`
	GroupSearchAttribute       types.String `tfsdk:"group_search_attribute"`
	GroupSearchFilter          types.String `tfsdk:"group_search_filter"`
	GroupSearchSubattribute    types.String `tfsdk:"group_search_subattribute"`
	IpAddress                  types.String `tfsdk:"ip_address"`
	LdapHostName               types.String `tfsdk:"ldap_host_name"`
	LoginName                  types.String `tfsdk:"login_name"`
	MaxLdapReferrals           types.Int64  `tfsdk:"max_ldap_referrals"`
	MaxNestingLevel            types.Int64  `tfsdk:"max_nesting_level"`
	Name                       types.String `tfsdk:"name"`
	NestedGroupExtraction      types.Bool   `tfsdk:"nested_group_extraction"`
	Port                       types.Int64  `tfsdk:"port"`
	SearchFilter               types.String `tfsdk:"search_filter"`
	SecType                    types.String `tfsdk:"sec_type"`
	SshPublicKey               types.String `tfsdk:"ssh_public_key"`
	SubattributeName           types.String `tfsdk:"subattribute_name"`
	Type                       types.String `tfsdk:"type"`
	ValidateLdapServerCerts    types.Bool   `tfsdk:"validate_ldap_server_certs"`
	Id                         types.String `tfsdk:"id"`
}

func ldapServerGetThePayloadFromtheConfig(ctx context.Context, data *ldapServerModel) ldapServerReq {
	tflog.Debug(ctx, "In ldapServerGetThePayloadFromtheConfig Function")
	ldapServerReqPayload := ldapServerReq{
		AuthTimeout:                data.AuthTimeout.ValueInt64(),
		Authentication:             data.Authentication.ValueBool(),
		BaseDn:                     data.BaseDn.ValueString(),
		BindDn:                     data.BindDn.ValueString(),
		BindPasswd:                 data.BindPasswd.ValueString(),
		ChangePassword:             data.ChangePassword.ValueBool(),
		DefaultAuthenticationGroup: data.DefaultAuthenticationGroup.ValueString(),
		FollowReferrals:            data.FollowReferrals.ValueBool(),
		GroupAttrName:              data.GroupAttrName.ValueString(),
		GroupNameIdentifier:        data.GroupNameIdentifier.ValueString(),
		GroupSearchAttribute:       data.GroupSearchAttribute.ValueString(),
		GroupSearchFilter:          data.GroupSearchFilter.ValueString(),
		GroupSearchSubattribute:    data.GroupSearchSubattribute.ValueString(),
		IpAddress:                  data.IpAddress.ValueString(),
		LdapHostName:               data.LdapHostName.ValueString(),
		LoginName:                  data.LoginName.ValueString(),
		MaxLdapReferrals:           data.MaxLdapReferrals.ValueInt64(),
		MaxNestingLevel:            data.MaxNestingLevel.ValueInt64(),
		Name:                       data.Name.ValueString(),
		NestedGroupExtraction:      data.NestedGroupExtraction.ValueBool(),
		Port:                       data.Port.ValueInt64(),
		SearchFilter:               data.SearchFilter.ValueString(),
		SecType:                    data.SecType.ValueString(),
		SshPublicKey:               data.SshPublicKey.ValueString(),
		SubattributeName:           data.SubattributeName.ValueString(),
		Type:                       data.Type.ValueString(),
		ValidateLdapServerCerts:    data.ValidateLdapServerCerts.ValueBool(),
	}
	return ldapServerReqPayload
}

type ldapServerReq struct {
	AuthTimeout                int64  `json:"auth_timeout,omitempty"`
	Authentication             bool   `json:"authentication,omitempty"`
	BaseDn                     string `json:"base_dn,omitempty"`
	BindDn                     string `json:"bind_dn,omitempty"`
	BindPasswd                 string `json:"bind_passwd,omitempty"`
	ChangePassword             bool   `json:"change_password,omitempty"`
	DefaultAuthenticationGroup string `json:"default_authentication_group,omitempty"`
	FollowReferrals            bool   `json:"follow_referrals,omitempty"`
	GroupAttrName              string `json:"group_attr_name,omitempty"`
	GroupNameIdentifier        string `json:"group_name_identifier,omitempty"`
	GroupSearchAttribute       string `json:"group_search_attribute,omitempty"`
	GroupSearchFilter          string `json:"group_search_filter,omitempty"`
	GroupSearchSubattribute    string `json:"group_search_subattribute,omitempty"`
	IpAddress                  string `json:"ip_address,omitempty"`
	LdapHostName               string `json:"ldap_host_name,omitempty"`
	LoginName                  string `json:"login_name,omitempty"`
	MaxLdapReferrals           int64  `json:"max_ldap_referrals,omitempty"`
	MaxNestingLevel            int64  `json:"max_nesting_level,omitempty"`
	Name                       string `json:"name,omitempty"`
	NestedGroupExtraction      bool   `json:"nested_group_extraction,omitempty"`
	Port                       int64  `json:"port,omitempty"`
	SearchFilter               string `json:"search_filter,omitempty"`
	SecType                    string `json:"sec_type,omitempty"`
	SshPublicKey               string `json:"ssh_public_key,omitempty"`
	SubattributeName           string `json:"subattribute_name,omitempty"`
	Type                       string `json:"type,omitempty"`
	ValidateLdapServerCerts    bool   `json:"validate_ldap_server_certs,omitempty"`
}
