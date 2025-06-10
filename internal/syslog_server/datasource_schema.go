package syslog_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func syslogServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a syslog server.",
		Attributes: map[string]schema.Attribute{
			"ip_address": schema.StringAttribute{
				Computed:            true,
				Description:         "Syslog server IP address. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Syslog server IP address. Minimum length =  1 Maximum length =  64",
			},
			"log_level_all": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send logs of all levels to this syslog server.",
				MarkdownDescription: "Send logs of all levels to this syslog server.",
			},
			"log_level_critical": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send logs of level critical to this syslog server.",
				MarkdownDescription: "Send logs of level critical to this syslog server.",
			},
			"log_level_error": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send logs of level error to this syslog server.",
				MarkdownDescription: "Send logs of level error to this syslog server.",
			},
			"log_level_info": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send logs of level info to this syslog server.",
				MarkdownDescription: "Send logs of level info to this syslog server.",
			},
			"log_level_none": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send no logs to this syslog server.",
				MarkdownDescription: "Send no logs to this syslog server.",
			},
			"log_level_warning": schema.BoolAttribute{
				Computed:            true,
				Description:         "Send logs of level warning to this syslog server.",
				MarkdownDescription: "Send logs of level warning to this syslog server.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Syslog server name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Syslog server name. Minimum length =  1 Maximum length =  128",
			},
			"port": schema.Int64Attribute{
				Computed:            true,
				Description:         "Syslog server port. Maximum value =  ",
				MarkdownDescription: "Syslog server port. Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the same as the name value.",
			},
		},
	}
}
