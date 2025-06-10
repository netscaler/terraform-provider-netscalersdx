package ntp_param

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ntpParamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a ntp param.",
		Attributes: map[string]schema.Attribute{
			"authentication": schema.BoolAttribute{
				Computed:            true,
				Description:         "Authentication Enabled.",
				MarkdownDescription: "Authentication Enabled.",
			},
			"automax_logsec": schema.Int64Attribute{
				Computed:            true,
				Description:         "Automax Interval (as power of 2 in seconds).",
				MarkdownDescription: "Automax Interval (as power of 2 in seconds).",
			},
			"revoke_logsec": schema.Int64Attribute{
				Computed:            true,
				Description:         "Revoke Interval (as power of 2 in seconds).",
				MarkdownDescription: "Revoke Interval (as power of 2 in seconds).",
			},
			"trusted_key_list": schema.ListAttribute{
				ElementType:         types.Int64Type,
				Computed:            true,
				Description:         "List of Trusted Key Identifiers for Symmetric Key Cryptography. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "List of Trusted Key Identifiers for Symmetric Key Cryptography. Minimum value =  1 Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource is the random string",
			},
		},
	}
}
