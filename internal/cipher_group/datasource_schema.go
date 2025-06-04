package cipher_group

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func cipherGroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source for managing Cipher Groups on Citrix ADC instance.",
		Attributes: map[string]schema.Attribute{
			"cipher_group_description": schema.StringAttribute{
				Computed:            true,
				Description:         "Describing the Cipher Group algorithms created. Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "Describing the Cipher Group algorithms created. Minimum length =  1 Maximum length =  256",
			},
			"cipher_group_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Name of Cipher Group. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of Cipher Group. Minimum length =  1 Maximum length =  128",
			},
			"cipher_name_list_array": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "list of cipher suites in form of array of strings.",
				MarkdownDescription: "list of cipher suites in form of array of strings.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}
