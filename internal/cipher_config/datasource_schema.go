package cipher_config

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func cipherConfigDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Retrieve an Cipher Config by its ID.",
		Attributes: map[string]schema.Attribute{
			"cipher_group_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
			},
			"cipher_name_list_array": schema.SetAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "list of cipher suites in form of array of strings.",
				MarkdownDescription: "list of cipher suites in form of array of strings.",
			},
			"config_mode": schema.StringAttribute{
				Computed:            true,
				Description:         "SSL Ciphers Config Mode [CipherGroup, CipherSuites].",
				MarkdownDescription: "SSL Ciphers Config Mode [CipherGroup, CipherSuites].",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}
