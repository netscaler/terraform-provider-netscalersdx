package snmp_user

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpUserResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP User resource.",
		Attributes: map[string]schema.Attribute{
			"auth_password": schema.StringAttribute{
				Optional:            true,
				Description:         "Authentication Password of SNMP User. Minimum length =  8 Maximum length =  32",
				MarkdownDescription: "Authentication Password of SNMP User. Minimum length =  8 Maximum length =  32",
			},
			"auth_protocol": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication Protocol of SNMP User. Values: 0:noValue, 1: MD5, 2: SHA1.",
				MarkdownDescription: `
Authentication Protocol of SNMP User.
	* Values:
		* 0: noValue
		* 1: MD5
		* 2: SHA1`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Name of SNMP User. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Name of SNMP User. Minimum length =  1 Maximum length =  32",
			},
			"privacy_password": schema.StringAttribute{
				Optional:            true,
				Description:         "Privacy Password of SNMP User. Minimum length =  8 Maximum length =  32",
				MarkdownDescription: "Privacy Password of SNMP User. Minimum length =  8 Maximum length =  32",
			},
			"privacy_protocol": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Privacy Protocol of SNMP User. Values: 0:noValue, 1: DES, 2: AES.",
				MarkdownDescription: `
Privacy Protocol of SNMP User.
	* Values: 
		* 0:noValue, 
		* 1: DES, 
		* 2: AES.`,
			},
			"security_level": schema.Int64Attribute{
				Required:    true,
				Description: "Security Level of SNMP User. Values: 0: noAuthNoPriv, 1: authNoPriv, 2: authPriv.",
				MarkdownDescription: `
Security Level of SNMP User. 
	* Values: 
		* 0: noAuthNoPriv, 
		* 1: authNoPriv, 
		* 2: authPriv.`,
			},
			"view_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "SNMP View Name attached to the SNMP User. Maximum length =  32",
				MarkdownDescription: "SNMP View Name attached to the SNMP User. Maximum length =  32",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpUserModel struct {
	AuthPassword    types.String `tfsdk:"auth_password"`
	AuthProtocol    types.Int64  `tfsdk:"auth_protocol"`
	Name            types.String `tfsdk:"name"`
	PrivacyPassword types.String `tfsdk:"privacy_password"`
	PrivacyProtocol types.Int64  `tfsdk:"privacy_protocol"`
	SecurityLevel   types.Int64  `tfsdk:"security_level"`
	ViewName        types.String `tfsdk:"view_name"`
	Id              types.String `tfsdk:"id"`
}

func snmpUserGetThePayloadFromtheConfig(ctx context.Context, data *snmpUserModel) snmpUserReq {
	tflog.Debug(ctx, "In snmpUserGetThePayloadFromtheConfig Function")
	snmpUserReqPayload := snmpUserReq{
		AuthPassword:    data.AuthPassword.ValueString(),
		Name:            data.Name.ValueString(),
		PrivacyPassword: data.PrivacyPassword.ValueString(),
		ViewName:        data.ViewName.ValueString(),
	}

	if !data.AuthProtocol.IsNull() && !data.AuthProtocol.IsUnknown() {
		snmpUserReqPayload.AuthProtocol = data.AuthProtocol.ValueInt64Pointer()
	}
	if !data.PrivacyProtocol.IsNull() && !data.PrivacyProtocol.IsUnknown() {
		snmpUserReqPayload.PrivacyProtocol = data.PrivacyProtocol.ValueInt64Pointer()
	}
	if !data.SecurityLevel.IsNull() && !data.SecurityLevel.IsUnknown() {
		snmpUserReqPayload.SecurityLevel = data.SecurityLevel.ValueInt64Pointer()
	}

	return snmpUserReqPayload
}
func snmpUserSetAttrFromGet(ctx context.Context, data *snmpUserModel, getResponseData map[string]interface{}) *snmpUserModel {
	tflog.Debug(ctx, "In snmpUserSetAttrFromGet Function")

	data.AuthProtocol = types.Int64Value(utils.StringToInt(getResponseData["auth_protocol"].(string)))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.PrivacyProtocol = types.Int64Value(utils.StringToInt(getResponseData["privacy_protocol"].(string)))
	data.SecurityLevel = types.Int64Value(utils.StringToInt(getResponseData["security_level"].(string)))
	data.ViewName = types.StringValue(getResponseData["view_name"].(string))

	return data
}

type snmpUserReq struct {
	AuthPassword    string `json:"auth_password,omitempty"`
	AuthProtocol    *int64 `json:"auth_protocol,omitempty"`
	Name            string `json:"name,omitempty"`
	PrivacyPassword string `json:"privacy_password,omitempty"`
	PrivacyProtocol *int64 `json:"privacy_protocol,omitempty"`
	SecurityLevel   *int64 `json:"security_level,omitempty"`
	ViewName        string `json:"view_name,omitempty"`
}
