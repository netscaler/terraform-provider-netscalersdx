package mpsuser

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mpsuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing mpsuser",
		Attributes: map[string]schema.Attribute{
			"enable_session_timeout": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enables session timeout for user.",
				MarkdownDescription: "Enables session timeout for user.",
			},
			"external_authentication": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable external authentication.",
				MarkdownDescription: "Enable external authentication.",
			},
			"groups": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "Groups to which user belongs.",
				MarkdownDescription: "Groups to which user belongs.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "User Name. Minimum length =  1 Maximum length =  128",
			},
			"password": schema.StringAttribute{
				Computed:            true,
				Description:         "Password. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Password. Minimum length =  1 Maximum length =  128",
			},
			"session_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "Session timeout for the user.",
				MarkdownDescription: "Session timeout for the user.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Computed:            true,
				Description:         "Session timeout unit for the user.",
				MarkdownDescription: "Session timeout unit for the user.",
			},
			"tenant_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Tenant Id of the system users. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Tenant Id of the system users. Minimum length =  1 Maximum length =  128",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource",
			},
		},
	}
}
