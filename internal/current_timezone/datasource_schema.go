package current_timezone

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func currentTimezoneDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a current timezone.",
		Attributes: map[string]schema.Attribute{
			"timezone": schema.StringAttribute{
				Computed:            true,
				Description:         "Timezone.",
				MarkdownDescription: "Timezone.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the random string",
			},
		},
	}
}
