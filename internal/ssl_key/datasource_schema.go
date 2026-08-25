package ssl_key

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func sslKeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description:         "Retrieve an SSL Key by its ID.",
		MarkdownDescription: "Retrieve an SSL Key by its ID.",
		Attributes: map[string]schema.Attribute{
			"file_name": schema.StringAttribute{
				Computed:            true,
				Description:         "File Name. Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "File Name. Minimum length =  1 Maximum length =  256",
			},
			"file_location_path": schema.StringAttribute{
				Computed:            true,
				Description:         "File Location. Not populated by the data source: the SDX appliance returns an empty value for this field on reads.",
				MarkdownDescription: "File Location. Not populated by the data source: the SDX appliance returns an empty value for this field on reads.",
			},
			"file_size": schema.Int64Attribute{
				Computed:            true,
				Description:         "File size in bytes as reported by the SDX appliance.",
				MarkdownDescription: "File size in bytes as reported by the SDX appliance.",
			},
			"file_last_modified": schema.StringAttribute{
				Computed:            true,
				Description:         "Last Modified timestamp of the file on the SDX appliance.",
				MarkdownDescription: "Last Modified timestamp of the file on the SDX appliance.",
			},
			"file_last_modified_epoch": schema.StringAttribute{
				Computed:            true,
				Description:         "Last Modified (Epoch) timestamp of the file on the SDX appliance, in epoch seconds.",
				MarkdownDescription: "Last Modified (Epoch) timestamp of the file on the SDX appliance, in epoch seconds.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The ID of this data source. It is the same as the file_name value.",
				MarkdownDescription: "The ID of this data source. It is the same as the `file_name` value.",
			},
		},
	}
}
