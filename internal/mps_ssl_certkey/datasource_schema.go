package mps_ssl_certkey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func mpsSslCertkeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a mps ssl Certkey.",
		Attributes: map[string]schema.Attribute{
			"fingerprint": schema.StringAttribute{
				Computed:            true,
				Description:         "SHA-1 Fingerprint of NetScaler ADM SSL Certificate. Minimum length =  1 Maximum length =  512",
				MarkdownDescription: "SHA-1 Fingerprint of NetScaler ADM SSL Certificate. Minimum length =  1 Maximum length =  512",
			},
			"password": schema.StringAttribute{
				Computed:            true,
				Description:         "The pass-phrase that was used to encrypt the private-key.. Maximum length =  32",
				MarkdownDescription: "The pass-phrase that was used to encrypt the private-key.. Maximum length =  32",
			},
			"ssl_certificate": schema.StringAttribute{
				Computed:            true,
				Description:         "Certificate. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Certificate. Minimum length =  1 Maximum length =  128",
			},
			"ssl_key": schema.StringAttribute{
				Computed:            true,
				Description:         "Key. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Key. Minimum length =  1 Maximum length =  128",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the unique randomstring.",
			},
		},
	}
}
