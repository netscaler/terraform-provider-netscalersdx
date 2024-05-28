package ssl_settings

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func sslSettingsResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SSL Settings resource.",
		Attributes: map[string]schema.Attribute{
			"sslreneg": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Enable SSL Renegotiation.",
				MarkdownDescription: "Enable SSL Renegotiation.",
			},
			"sslv3": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Enable SSLv3.",
				MarkdownDescription: "Enable SSLv3.",
			},
			"tlsv1": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Enable TLSv1.",
				MarkdownDescription: "Enable TLSv1.",
			},
			"tlsv1_1": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Enable TLSv1.1.",
				MarkdownDescription: "Enable TLSv1.1.",
			},
			"tlsv1_2": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description:         "Enable TLSv1.2.",
				MarkdownDescription: "Enable TLSv1.2.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type sslSettingsModel struct {
	Sslreneg types.Bool   `tfsdk:"sslreneg"`
	Sslv3    types.Bool   `tfsdk:"sslv3"`
	Tlsv1    types.Bool   `tfsdk:"tlsv1"`
	Tlsv11   types.Bool   `tfsdk:"tlsv1_1"`
	Tlsv12   types.Bool   `tfsdk:"tlsv1_2"`
	Id       types.String `tfsdk:"id"`
}

func sslSettingsGetThePayloadFromtheConfig(ctx context.Context, data *sslSettingsModel) sslSettingsReq {
	tflog.Debug(ctx, "In sslSettingsGetThePayloadFromtheConfig Function")
	sslSettingsReqPayload := sslSettingsReq{
		Sslreneg: data.Sslreneg.ValueBoolPointer(),
		Sslv3:    data.Sslv3.ValueBoolPointer(),
		Tlsv1:    data.Tlsv1.ValueBoolPointer(),
		Tlsv11:   data.Tlsv11.ValueBoolPointer(),
		Tlsv12:   data.Tlsv12.ValueBoolPointer(),
	}
	return sslSettingsReqPayload
}
func sslSettingsSetAttrFromGet(ctx context.Context, data *sslSettingsModel, getResponseData map[string]interface{}) *sslSettingsModel {
	tflog.Debug(ctx, "In sslSettingsSetAttrFromGet Function")
	if !data.Sslreneg.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["sslreneg"].(string))
		data.Sslreneg = types.BoolValue(val)
	}
	if !data.Sslv3.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["sslv3"].(string))
		data.Sslv3 = types.BoolValue(val)
	}
	if !data.Tlsv1.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["tlsv1"].(string))
		data.Tlsv1 = types.BoolValue(val)
	}
	if !data.Tlsv11.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["tlsv1_1"].(string))
		data.Tlsv11 = types.BoolValue(val)
	}
	if !data.Tlsv12.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["tlsv1_2"].(string))
		data.Tlsv12 = types.BoolValue(val)
	}
	return data
}

type sslSettingsReq struct {
	Sslreneg *bool `json:"sslreneg,omitempty"`
	Sslv3    *bool `json:"sslv3,omitempty"`
	Tlsv1    *bool `json:"tlsv1,omitempty"`
	Tlsv11   *bool `json:"tlsv1_1,omitempty"`
	Tlsv12   *bool `json:"tlsv1_2,omitempty"`
}
