package snmp_alarm_config

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func snmpAlarmConfigDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a Snmp Alarm Config.",
		Attributes: map[string]schema.Attribute{
			"enable": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable Alarm.",
				MarkdownDescription: "Enable Alarm.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:            true,
				Description:         "Alarm Name. Maximum length =  128",
				MarkdownDescription: "Alarm Name. Maximum length =  128",
			},
			"severity": schema.StringAttribute{
				Computed:            true,
				Description:         "Alarm severity. Supported values: Critical, Major, Minor, Warning, Informational . Maximum length =  128",
				MarkdownDescription: "Alarm severity. Supported values: Critical, Major, Minor, Warning, Informational . Maximum length =  128",
			},
			"threshold": schema.Int64Attribute{
				Computed:            true,
				Description:         "Threshold Value for the alarm. Maximum value =  ",
				MarkdownDescription: "Threshold Value for the alarm. Maximum value =  ",
			},
			"time": schema.Int64Attribute{
				Computed:            true,
				Description:         "Frequency of the alarm in minutes.",
				MarkdownDescription: "Frequency of the alarm in minutes.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource, it is same as name.",
			},
		},
	}
}
