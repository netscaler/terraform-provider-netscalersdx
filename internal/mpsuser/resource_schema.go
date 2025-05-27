package mpsuser

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-netscalersdx/internal/utils"
)

func mpsuserResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for System User resource.",
		Attributes: map[string]schema.Attribute{
			"enable_session_timeout": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enables session timeout for user.",
				MarkdownDescription: "Enables session timeout for user.",
			},
			"external_authentication": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable external authentication.",
				MarkdownDescription: "Enable external authentication.",
			},
			"groups": schema.ListAttribute{
				ElementType:         types.StringType,
				Required:            true,
				Description:         "Groups to which user belongs.",
				MarkdownDescription: "Groups to which user belongs.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "User Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "User Name. Minimum length =  1 Maximum length =  128",
			},
			"password": schema.StringAttribute{
				Required:            true,
				Description:         "Password. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Password. Minimum length =  1 Maximum length =  128",
			},
			"session_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Session timeout for the user.",
				MarkdownDescription: "Session timeout for the user.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Session timeout unit for the user.",
				MarkdownDescription: "Session timeout unit for the user.",
			},
			"tenant_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Tenant Id of the system users. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Tenant Id of the system users. Minimum length =  1 Maximum length =  128",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type mpsuserModel struct {
	EnableSessionTimeout   types.Bool   `tfsdk:"enable_session_timeout"`
	ExternalAuthentication types.Bool   `tfsdk:"external_authentication"`
	Groups                 types.List   `tfsdk:"groups"`
	Name                   types.String `tfsdk:"name"`
	Password               types.String `tfsdk:"password"`
	SessionTimeout         types.Int64  `tfsdk:"session_timeout"`
	SessionTimeoutUnit     types.String `tfsdk:"session_timeout_unit"`
	TenantId               types.String `tfsdk:"tenant_id"`
	Id                     types.String `tfsdk:"id"`
}

func mpsuserGetThePayloadFromtheConfig(ctx context.Context, data *mpsuserModel) mpsuserReq {
	tflog.Debug(ctx, "In mpsuserGetThePayloadFromtheConfig Function")

	mpsuserReqPayload := mpsuserReq{
		Groups:             utils.TypeListToUnmarshalStringList(data.Groups),
		Name:               data.Name.ValueString(),
		Password:           data.Password.ValueString(),
		SessionTimeoutUnit: data.SessionTimeoutUnit.ValueString(),
		TenantId:           data.TenantId.ValueString(),
	}

	if !data.ExternalAuthentication.IsNull() && !data.ExternalAuthentication.IsUnknown() {
		mpsuserReqPayload.ExternalAuthentication = data.ExternalAuthentication.ValueBoolPointer()
	}
	if !data.EnableSessionTimeout.IsNull() && !data.EnableSessionTimeout.IsUnknown() {
		mpsuserReqPayload.EnableSessionTimeout = data.EnableSessionTimeout.ValueBoolPointer()
	}
	if !data.SessionTimeout.IsNull() && !data.SessionTimeout.IsUnknown() {
		mpsuserReqPayload.SessionTimeout = data.SessionTimeout.ValueInt64Pointer()
	}

	return mpsuserReqPayload
}
func mpsuserSetAttrFromGet(ctx context.Context, data *mpsuserModel, getResponseData map[string]interface{}) *mpsuserModel {
	tflog.Debug(ctx, "In mpsuserSetAttrFromGet Function")

	data.EnableSessionTimeout = types.BoolValue(utils.StringToBool(getResponseData["enable_session_timeout"].(string)))
	data.ExternalAuthentication = types.BoolValue(utils.StringToBool(getResponseData["external_authentication"].(string)))
	data.Groups = utils.StringListToTypeList(utils.ToStringList(getResponseData["groups"].([]interface{})))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.SessionTimeout = types.Int64Value(utils.StringToInt(getResponseData["session_timeout"].(string)))
	data.SessionTimeoutUnit = types.StringValue(getResponseData["session_timeout_unit"].(string))
	data.TenantId = types.StringValue(getResponseData["tenant_id"].(string))

	return data
}

type mpsuserReq struct {
	EnableSessionTimeout   *bool    `json:"enable_session_timeout,omitempty"`
	ExternalAuthentication *bool    `json:"external_authentication,omitempty"`
	Groups                 []string `json:"groups,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Password               string   `json:"password,omitempty"`
	SessionTimeout         *int64   `json:"session_timeout,omitempty"`
	SessionTimeoutUnit     string   `json:"session_timeout_unit,omitempty"`
	TenantId               string   `json:"tenant_id,omitempty"`
}
