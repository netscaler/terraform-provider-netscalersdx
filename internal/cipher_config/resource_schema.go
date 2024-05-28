package cipher_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-netscalersdx/internal/utils"
)

func cipherConfigResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SSL Cipher Config resource.",
		Attributes: map[string]schema.Attribute{
			"cipher_group_name": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
			},
			"cipher_name_list_array": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description:         "list of cipher suites in form of array of strings.",
				MarkdownDescription: "list of cipher suites in form of array of strings.",
			},
			"config_mode": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

type cipherConfigModel struct {
	CipherGroupName     types.String `tfsdk:"cipher_group_name"`
	CipherNameListArray types.List   `tfsdk:"cipher_name_list_array"`
	ConfigMode          types.String `tfsdk:"config_mode"`
	Id                  types.String `tfsdk:"id"`
}

func cipherConfigGetThePayloadFromtheConfig(ctx context.Context, data *cipherConfigModel) cipherConfigReq {
	tflog.Debug(ctx, "In cipherConfigGetThePayloadFromtheConfig Function")
	cipherConfigReqPayload := cipherConfigReq{
		CipherGroupName:     data.CipherGroupName.ValueString(),
		CipherNameListArray: utils.TypeListToUnmarshalStringList(data.CipherNameListArray),
		ConfigMode:          data.ConfigMode.ValueString(),
	}
	return cipherConfigReqPayload
}
func cipherConfigSetAttrFromGet(ctx context.Context, data *cipherConfigModel, getResponseData map[string]interface{}) *cipherConfigModel {
	tflog.Debug(ctx, "In cipherConfigSetAttrFromGet Function")
	if !data.CipherGroupName.IsNull() {
		data.CipherGroupName = types.StringValue(getResponseData["cipher_group_name"].(string))
	}
	if !data.CipherNameListArray.IsNull() {
		data.CipherNameListArray = utils.StringListToTypeList(utils.ToStringList(getResponseData["cipher_name_list_array"].([]interface{})))
	}
	if !data.ConfigMode.IsNull() {
		data.ConfigMode = types.StringValue(getResponseData["config_mode"].(string))
	}
	return data
}

type cipherConfigReq struct {
	CipherGroupName     string   `json:"cipher_group_name,omitempty"`
	CipherNameListArray []string `json:"cipher_name_list_array,omitempty"`
	ConfigMode          string   `json:"config_mode,omitempty"`
}
