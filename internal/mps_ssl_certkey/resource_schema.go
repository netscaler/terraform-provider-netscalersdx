package mps_ssl_certkey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func mpsSslCertkeyResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for Install SSL certificate on Management Service resource",
		Attributes: map[string]schema.Attribute{
			"fingerprint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "SHA-1 Fingerprint of NetScaler ADM SSL Certificate. Minimum length =  1 Maximum length =  512",
				MarkdownDescription: "SHA-1 Fingerprint of NetScaler ADM SSL Certificate. Minimum length =  1 Maximum length =  512",
			},
			"password": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "The pass-phrase that was used to encrypt the private-key.. Maximum length =  32",
				MarkdownDescription: "The pass-phrase that was used to encrypt the private-key.. Maximum length =  32",
			},
			"ssl_certificate": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "File name of a valid certificate. The certificate file must be located in the `/var/mps/ssl_certs/` directory on the Management Service virtual appliance(SDX). Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "File name of a valid certificate. The certificate file must be located in the `/var/mps/ssl_certs/` directory on the Management Service virtual appliance(SDX). Minimum length =  1 Maximum length =  128",
			},
			"ssl_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "File name of the private key used to create the certificate. The key file must be located in the `/var/mps/ssl_certs/` directory on the Management Service virtual appliance(SDX). Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "File name of the private key used to create the certificate. The key file must be located in the `/var/mps/ssl_certs/` directory on the Management Service virtual appliance(SDX). Minimum length =  1 Maximum length =  128",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type mpsSslCertkeyModel struct {
	Fingerprint    types.String `tfsdk:"fingerprint"`
	Password       types.String `tfsdk:"password"`
	SslCertificate types.String `tfsdk:"ssl_certificate"`
	SslKey         types.String `tfsdk:"ssl_key"`
	Id             types.String `tfsdk:"id"`
}

func mpsSslCertkeyGetThePayloadFromtheConfig(ctx context.Context, data *mpsSslCertkeyModel) mpsSslCertkeyReq {
	tflog.Debug(ctx, "In mpsSslCertkeyGetThePayloadFromtheConfig Function")
	mpsSslCertkeyReqPayload := mpsSslCertkeyReq{
		Fingerprint:    data.Fingerprint.ValueString(),
		Password:       data.Password.ValueString(),
		SslCertificate: data.SslCertificate.ValueString(),
		SslKey:         data.SslKey.ValueString(),
	}
	return mpsSslCertkeyReqPayload
}
func mpsSslCertkeySetAttrFromGet(ctx context.Context, data *mpsSslCertkeyModel, getResponseData map[string]interface{}) *mpsSslCertkeyModel {
	tflog.Debug(ctx, "In mpsSslCertkeySetAttrFromGet Function")

	data.Fingerprint = types.StringValue(getResponseData["fingerprint"].(string))
	// data.Password = types.StringValue(getResponseData["password"].(string))
	data.SslCertificate = types.StringValue(getResponseData["ssl_certificate"].(string))
	data.SslKey = types.StringValue(getResponseData["ssl_key"].(string))

	return data
}

type mpsSslCertkeyReq struct {
	Fingerprint    string `json:"fingerprint,omitempty"`
	Password       string `json:"password,omitempty"`
	SslCertificate string `json:"ssl_certificate,omitempty"`
	SslKey         string `json:"ssl_key,omitempty"`
}
