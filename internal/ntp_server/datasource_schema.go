package ntp_server

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ntpServerDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a ntp server.",
		Attributes: map[string]schema.Attribute{
			"autokey": schema.BoolAttribute{
				Computed:            true,
				Description:         "Autokey Public Key Authentication.",
				MarkdownDescription: "Autokey Public Key Authentication.",
			},
			"client": schema.StringAttribute{
				Computed:            true,
				Description:         "Sender of request, whether from Setup Wizard or direct NTP configuration.",
				MarkdownDescription: "Sender of request, whether from Setup Wizard or direct NTP configuration.",
			},
			"key_id": schema.Int64Attribute{
				Computed:            true,
				Description:         "Key Identifier for Symmetric Key Authentication. Maximum value =  ",
				MarkdownDescription: "Key Identifier for Symmetric Key Authentication. Maximum value =  ",
			},
			"maxpoll": schema.Int64Attribute{
				Computed:            true,
				Description:         "Maximum Poll Interval. Maximum value =  ",
				MarkdownDescription: "Maximum Poll Interval. Maximum value =  ",
			},
			"minpoll": schema.Int64Attribute{
				Computed:            true,
				Description:         "Minimum Poll Interval. Maximum value =  ",
				MarkdownDescription: "Minimum Poll Interval. Maximum value =  ",
			},
			"preferred_server": schema.BoolAttribute{
				Computed:            true,
				Description:         "NTP Server Preferred.",
				MarkdownDescription: "NTP Server Preferred.",
			},
			"server": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "NTP Time Server Address.",
				MarkdownDescription: "NTP Time Server Address.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource, it is same as the server address.",
			},
		},
	}
}
