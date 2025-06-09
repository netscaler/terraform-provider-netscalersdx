package radius_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func radiusServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Retrieve an Radius Server.",
		Attributes: map[string]schema.Attribute{
			"accounting": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable accounting in the radius server.",
				MarkdownDescription: "Enable accounting in the radius server.",
			},
			"address_type": schema.Int64Attribute{
				Computed:            true,
				Description:         "Configuration Type. Values: 0: IPv4, 1: IPv6, -1: Hostname.",
				MarkdownDescription: "Configuration Type. Values: 0: IPv4, 1: IPv6, -1: Hostname.",
			},
			"attribute_type": schema.Int64Attribute{
				Computed:            true,
				Description:         "Attribute type for RADIUS group extraction. Maximum value =  ",
				MarkdownDescription: "Attribute type for RADIUS group extraction. Maximum value =  ",
			},
			"auth_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the Radius server. Maximum value =  ",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the Radius server. Maximum value =  ",
			},
			"default_authentication_group": schema.StringAttribute{
				Computed:            true,
				Description:         "This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64",
				MarkdownDescription: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64",
			},
			"enable_nas_ip": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable NAS IP extraction.",
				MarkdownDescription: "Enable NAS IP extraction.",
			},
			"group_separator": schema.StringAttribute{
				Computed:            true,
				Description:         " Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  7",
				MarkdownDescription: " Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  7",
			},
			"groups_prefix": schema.StringAttribute{
				Computed:            true,
				Description:         "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  31",
				MarkdownDescription: "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  31",
			},
			"ip_address": schema.StringAttribute{
				Computed:            true,
				Description:         "IP Address of radius server. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IP Address of radius server. Minimum length =  1 Maximum length =  64",
			},
			"ip_attribute_type": schema.Int64Attribute{
				Computed:            true,
				Description:         "The attribute type of the remote IP address attribute in a RADIUS response. Maximum value =  ",
				MarkdownDescription: "The attribute type of the remote IP address attribute in a RADIUS response. Maximum value =  ",
			},
			"ip_vendor_id": schema.Int64Attribute{
				Computed:            true,
				Description:         "The vendor ID of the attribute in the RADIUS response which denotes the intranet IP. Maximum value =  ",
				MarkdownDescription: "The vendor ID of the attribute in the RADIUS response which denotes the intranet IP. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of radius server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of radius server. Minimum length =  1 Maximum length =  128",
			},
			"nas_id": schema.StringAttribute{
				Computed:            true,
				Description:         "NAS ID. Maximum length =  128",
				MarkdownDescription: "NAS ID. Maximum length =  128",
			},
			"pass_encoding": schema.StringAttribute{
				Computed:            true,
				Description:         "Enable password encoding in RADIUS packets send to the RADIUS server.",
				MarkdownDescription: "Enable password encoding in RADIUS packets send to the RADIUS server.",
			},
			"port": schema.Int64Attribute{
				Computed:            true,
				Description:         "Port number of radius server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Port number of radius server. Minimum value =  1 Maximum value =  ",
			},
			"pwd_attribute_type": schema.Int64Attribute{
				Computed:            true,
				Description:         "The attribute type of the password attribute in a RADIUS response.. Maximum value =  ",
				MarkdownDescription: "The attribute type of the password attribute in a RADIUS response.. Maximum value =  ",
			},
			"pwd_vendor_id": schema.Int64Attribute{
				Computed:            true,
				Description:         "Vendor ID of the password in the RADIUS response. Used to extract the user password. Maximum value =  ",
				MarkdownDescription: "Vendor ID of the password in the RADIUS response. Used to extract the user password. Maximum value =  ",
			},
			"radius_key": schema.StringAttribute{
				Computed:            true,
				Description:         "Key of radius server. Minimum length =  4 Maximum length =  32",
				MarkdownDescription: "Key of radius server. Minimum length =  4 Maximum length =  32",
			},
			"vendor_id": schema.Int64Attribute{
				Computed:            true,
				Description:         "Vendor ID for RADIUS group extraction.",
				MarkdownDescription: "Vendor ID for RADIUS group extraction.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource",
			},
		},
	}
}
