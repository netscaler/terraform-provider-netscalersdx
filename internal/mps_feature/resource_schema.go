package mps_feature

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func mpsFeatureResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for Feature Toggle Status resource.",
		Attributes: map[string]schema.Attribute{
			"admin_toggle": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "This is Admin controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend.. Maximum value =  ",
				MarkdownDescription: "This is Admin controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend.. Maximum value =  ",
			},
			"built_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "This is Ops controllable and will not be visible to the Admin to control. If true: Ops controllable feature, false: Admin controllable feature..",
				MarkdownDescription: "This is Ops controllable and will not be visible to the Admin to control. If true: Ops controllable feature, false: Admin controllable feature..",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Feature Description..",
				MarkdownDescription: "Feature Description..",
			},
			"feature_name": schema.StringAttribute{
				Required:            true,
				Description:         "Feature Name.",
				MarkdownDescription: "Feature Name.",
			},
			"ops_toggle": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "This is Ops controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend. Ops controlled takes higher precedence than Admin Controlled.. Maximum value =  ",
				MarkdownDescription: "This is Ops controllable. 0: Disable UI and Backend, 1: Disable UI and enable Backend, 2: Enable UI and disable Backend, 3: Enable UI and Backend. Ops controlled takes higher precedence than Admin Controlled.. Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type mpsFeatureModel struct {
	AdminToggle types.Int64  `tfsdk:"admin_toggle"`
	BuiltIn     types.Bool   `tfsdk:"built_in"`
	Description types.String `tfsdk:"description"`
	FeatureName types.String `tfsdk:"feature_name"`
	OpsToggle   types.Int64  `tfsdk:"ops_toggle"`
	Id          types.String `tfsdk:"id"`
}

func mpsFeatureGetThePayloadFromtheConfig(ctx context.Context, data *mpsFeatureModel) mpsFeatureReq {
	tflog.Debug(ctx, "In mpsFeatureGetThePayloadFromtheConfig Function")
	mpsFeatureReqPayload := mpsFeatureReq{
		Description: data.Description.ValueString(),
		FeatureName: data.FeatureName.ValueString(),
	}

	if !data.AdminToggle.IsNull() && !data.AdminToggle.IsUnknown() {
		mpsFeatureReqPayload.AdminToggle = data.AdminToggle.ValueInt64Pointer()
	}
	if !data.BuiltIn.IsNull() && !data.BuiltIn.IsUnknown() {
		mpsFeatureReqPayload.BuiltIn = data.BuiltIn.ValueBoolPointer()
	}
	if !data.OpsToggle.IsNull() && !data.OpsToggle.IsUnknown() {
		mpsFeatureReqPayload.OpsToggle = data.OpsToggle.ValueInt64Pointer()
	}

	return mpsFeatureReqPayload
}
func mpsFeatureSetAttrFromGet(ctx context.Context, data *mpsFeatureModel, getResponseData map[string]interface{}) *mpsFeatureModel {
	tflog.Debug(ctx, "In mpsFeatureSetAttrFromGet Function")

	data.AdminToggle = types.Int64Value(utils.StringToInt(getResponseData["admin_toggle"].(string)))
	data.BuiltIn = types.BoolValue(utils.StringToBool(getResponseData["built_in"].(string)))
	data.Description = types.StringValue(getResponseData["description"].(string))
	data.FeatureName = types.StringValue(getResponseData["feature_name"].(string))
	data.OpsToggle = types.Int64Value(utils.StringToInt(getResponseData["ops_toggle"].(string)))

	return data
}

type mpsFeatureReq struct {
	AdminToggle *int64 `json:"admin_toggle,omitempty"`
	BuiltIn     *bool  `json:"built_in,omitempty"`
	Description string `json:"description,omitempty"`
	FeatureName string `json:"feature_name,omitempty"`
	OpsToggle   *int64 `json:"ops_toggle,omitempty"`
}
