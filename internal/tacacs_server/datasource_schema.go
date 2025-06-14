package tacacs_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func tacacsServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a tacacs server.",
		Attributes: map[string]schema.Attribute{
			"accounting": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable accounting in the tacacs server.",
				MarkdownDescription: "Enable accounting in the tacacs server.",
			},
			"auth_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the TACACS server. Minimum value =  1 Maximum value =  ",
			},
			"group_attr_name": schema.StringAttribute{
				Computed:            true,
				Description:         "The Attribute name for group extraction from the ACS server. If not passed, then groups will not be extracted. No other harm. Maximum length =  64",
				MarkdownDescription: "The Attribute name for group extraction from the ACS server. If not passed, then groups will not be extracted. No other harm. Maximum length =  64",
			},
			"ip_address": schema.StringAttribute{
				Computed:            true,
				Description:         "IP Address of TACACS server. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IP Address of TACACS server. Minimum length =  1 Maximum length =  64",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Name of TACACS server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of TACACS server. Minimum length =  1 Maximum length =  128",
			},
			"port": schema.Int64Attribute{
				Computed:            true,
				Description:         "port number of TACACS server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "port number of TACACS server. Minimum value =  1 Maximum value =  ",
			},
			"tacacs_key": schema.StringAttribute{
				Computed:            true,
				Description:         "Key shared between the TACACS+ server and clients. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Key shared between the TACACS+ server and clients. Minimum length =  1 Maximum length =  64",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource, it is the random string generated by remote SDX.",
			},
		},
	}
}
