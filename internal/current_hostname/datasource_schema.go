package current_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func currentHostnameDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a current hostname.",
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Computed:            true,
				Description:         "Hostname. Minimum length =  1 Maximum length =  63",
				MarkdownDescription: "Hostname. Minimum length =  1 Maximum length =  63",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed:            true,
				Description:         "Hypervisor Hostname. Minimum length =  1 Maximum length =  63",
				MarkdownDescription: "Hypervisor Hostname. Minimum length =  1 Maximum length =  63",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the random string",
			},
		},
	}
}
