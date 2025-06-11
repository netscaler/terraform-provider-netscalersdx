package ssl_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func sslSettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a ssl settings.",
		Attributes: map[string]schema.Attribute{
			"sslreneg": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable SSL Renegotiation.",
				MarkdownDescription: "Enable SSL Renegotiation.",
			},
			"sslv3": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable SSLv3.",
				MarkdownDescription: "Enable SSLv3.",
			},
			"tlsv1": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable TLSv1.",
				MarkdownDescription: "Enable TLSv1.",
			},
			"tlsv1_1": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable TLSv1.1.",
				MarkdownDescription: "Enable TLSv1.1.",
			},
			"tlsv1_2": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable TLSv1.2.",
				MarkdownDescription: "Enable TLSv1.2.",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the unique randomstring.",
			},
		},
	}
}
