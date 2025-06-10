package syslog_params

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func syslogParamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a syslog params.",
		Attributes: map[string]schema.Attribute{
			"date_format": schema.StringAttribute{
				Computed:            true,
				Description:         "Format of date to be added in the syslog message.",
				MarkdownDescription: "Format of date to be added in the syslog message.",
			},
			"timezone": schema.StringAttribute{
				Computed:            true,
				Description:         "Timezone to be used in the syslog message.",
				MarkdownDescription: "Timezone to be used in the syslog message.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the unique randomstring.",
			},
		},
	}
}
