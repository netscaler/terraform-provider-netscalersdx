package cipher_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of the Cipher Group. Minimum length =  1 Maximum length =  128",
			},
			"cipher_name_list_array": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Description:         "list of cipher suites in form of array of strings.",
				MarkdownDescription: "list of cipher suites in form of array of strings.",
			},
			"config_mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
	CipherNameListArray types.Set    `tfsdk:"cipher_name_list_array"`
	ConfigMode          types.String `tfsdk:"config_mode"`
	Id                  types.String `tfsdk:"id"`
}

func cipherConfigGetThePayloadFromtheConfig(ctx context.Context, data *cipherConfigModel) cipherConfigReq {
	tflog.Debug(ctx, "In cipherConfigGetThePayloadFromtheConfig Function")
	cipherConfigReqPayload := cipherConfigReq{
		CipherGroupName:     data.CipherGroupName.ValueString(),
		CipherNameListArray: utils.TypeListToUnmarshalStringSet(data.CipherNameListArray),
		ConfigMode:          data.ConfigMode.ValueString(),
	}
	return cipherConfigReqPayload
}
func cipherConfigSetAttrFromGet(ctx context.Context, data *cipherConfigModel, getResponseData map[string]interface{}) *cipherConfigModel {
	tflog.Debug(ctx, "In cipherConfigSetAttrFromGet Function")

	data.CipherGroupName = types.StringValue(getResponseData["cipher_group_name"].(string))
	data.CipherNameListArray = utils.StringListToTypeSet(utils.ToStringList(getResponseData["cipher_name_list_array"].([]interface{})))
	data.ConfigMode = types.StringValue(getResponseData["config_mode"].(string))

	return data
}

type cipherConfigReq struct {
	CipherGroupName     string   `json:"cipher_group_name,omitempty"`
	CipherNameListArray []string `json:"cipher_name_list_array,omitempty"`
	ConfigMode          string   `json:"config_mode,omitempty"`
}
