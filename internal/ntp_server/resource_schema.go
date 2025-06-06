// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ntp_server

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func ntpServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for NTP Server Information resource.",
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "NTP Time Server Address.",
				MarkdownDescription: "NTP Time Server Address.",
			},
			"autokey": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Autokey Public Key Authentication.",
				MarkdownDescription: "Autokey Public Key Authentication.",
			},
			"client": schema.StringAttribute{
				Optional:            true,
				Description:         "Sender of request, whether from Setup Wizard or direct NTP configuration.",
				MarkdownDescription: "Sender of request, whether from Setup Wizard or direct NTP configuration.",
			},
			"key_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Key Identifier for Symmetric Key Authentication.",
				MarkdownDescription: "Key Identifier for Symmetric Key Authentication.",
			},
			"maxpoll": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Maximum Poll Interval.",
				MarkdownDescription: "Maximum Poll Interval.",
			},
			"minpoll": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Minimum Poll Interval.",
				MarkdownDescription: "Minimum Poll Interval.",
			},
			"preferred_server": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "NTP Server Preferred.",
				MarkdownDescription: "NTP Server Preferred.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type ntpServerModel struct {
	Autokey         types.Bool   `tfsdk:"autokey"`
	Client          types.String `tfsdk:"client"`
	KeyId           types.Int64  `tfsdk:"key_id"`
	Maxpoll         types.Int64  `tfsdk:"maxpoll"`
	Minpoll         types.Int64  `tfsdk:"minpoll"`
	PreferredServer types.Bool   `tfsdk:"preferred_server"`
	Server          types.String `tfsdk:"server"`
	Id              types.String `tfsdk:"id"`
}

func ntpServerSetAttrFromGet(ctx context.Context, data *ntpServerModel, getResponseData map[string]interface{}) *ntpServerModel {

	data.Autokey = types.BoolValue(utils.StringToBool(getResponseData["autokey"].(string)))
	data.KeyId = types.Int64Value(utils.StringToInt(getResponseData["key_id"].(string)))
	data.Maxpoll = types.Int64Value(utils.StringToInt(getResponseData["maxpoll"].(string)))
	data.Minpoll = types.Int64Value(utils.StringToInt(getResponseData["minpoll"].(string)))
	data.PreferredServer = types.BoolValue(utils.StringToBool(getResponseData["preferred_server"].(string)))
	data.Server = types.StringValue(getResponseData["server"].(string))

	return data

}

func ntpServerGetThePayloadFromtheConfig(ctx context.Context, data *ntpServerModel) ntpServerReq {
	tflog.Debug(ctx, "In ntpServerGetThePayloadFromtheConfig Function")
	ntpServerReqPayload := ntpServerReq{
		Client: data.Client.ValueString(),
		Server: data.Server.ValueString(),
	}

	if !data.Autokey.IsNull() && !data.Autokey.IsUnknown() {
		ntpServerReqPayload.Autokey = data.Autokey.ValueBoolPointer()
	}
	if !data.KeyId.IsNull() && !data.KeyId.IsUnknown() {
		ntpServerReqPayload.KeyId = data.KeyId.ValueInt64Pointer()
	}
	if !data.Maxpoll.IsNull() && !data.Maxpoll.IsUnknown() {
		ntpServerReqPayload.Maxpoll = data.Maxpoll.ValueInt64Pointer()
	}
	if !data.Minpoll.IsNull() && !data.Minpoll.IsUnknown() {
		ntpServerReqPayload.Minpoll = data.Minpoll.ValueInt64Pointer()
	}
	if !data.PreferredServer.IsNull() && !data.PreferredServer.IsUnknown() {
		ntpServerReqPayload.PreferredServer = data.PreferredServer.ValueBoolPointer()
	}

	return ntpServerReqPayload
}

type ntpServerReq struct {
	Autokey         *bool  `json:"autokey"`
	Client          string `json:"client,omitempty"`
	KeyId           *int64 `json:"key_id,omitempty"`
	Maxpoll         *int64 `json:"maxpoll,omitempty"`
	Minpoll         *int64 `json:"minpoll,omitempty"`
	PreferredServer *bool  `json:"preferred_server,omitempty"`
	Server          string `json:"server,omitempty"`
}
