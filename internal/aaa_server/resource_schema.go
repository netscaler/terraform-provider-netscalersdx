package aaa_server

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func aaaServerResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for AAA Server configuration resource.",
		Attributes: map[string]schema.Attribute{
			"external_servers": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"external_server_name": schema.StringAttribute{
							Required:            true,
							Description:         "Name of external server. Minimum length =  1 Maximum length =  128",
							MarkdownDescription: "Name of external server. Minimum length =  1 Maximum length =  128",
						},
						"external_server_type": schema.StringAttribute{
							Required:            true,
							Description:         "Type of external server. Supported types 1.RADIUS 2.LDAP 3.TACACS 4.KEYSTONE. Minimum length =  1 Maximum length =  32",
							MarkdownDescription: "Type of external server. Supported types 1.RADIUS 2.LDAP 3.TACACS 4.KEYSTONE. Minimum length =  1 Maximum length =  32",
						},
						"priority": schema.Int64Attribute{
							Required:            true,
							Description:         "Priority of external server. Minimum value =  2 Maximum value =  ",
							MarkdownDescription: "Priority of external server. Minimum value =  2 Maximum value =  ",
						},
					},
				},
				Optional:            true,
				Computed:            false,
				Description:         "List of external servers.",
				MarkdownDescription: "List of external servers.",
			},
			"fallback_local_authentication": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable local fallback authentication.",
				MarkdownDescription: "Enable local fallback authentication.",
			},
			"log_ext_group_info": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Log external group info.",
				MarkdownDescription: "Log external group info.",
			},
			"primary_server_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of primary server name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of primary server name. Minimum length =  1 Maximum length =  128",
			},
			"primary_server_type": schema.StringAttribute{
				Required:            true,
				Description:         "Type of primary server. Supported types 1. LOCAL 2.RADIUS 3.LDAP 4.TACACS 5.KEYSTONE. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Type of primary server. Supported types 1. LOCAL 2.RADIUS 3.LDAP 4.TACACS 5.KEYSTONE. Minimum length =  1 Maximum length =  32",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func aaaServerGetThePayloadFromtheConfig(ctx context.Context, data *aaaServerModel) aaaServerReq {
	tflog.Debug(ctx, "In aaaServerGetThePayloadFromtheConfig Function")
	aaaServerReqPayload := aaaServerReq{
		ExternalServers:             externalServersFromConfigToRequest(ctx, data.ExternalServers),
		FallbackLocalAuthentication: data.FallbackLocalAuthentication.ValueBoolPointer(),
		LogExtGroupInfo:             data.LogExtGroupInfo.ValueBoolPointer(),
		PrimaryServerName:           data.PrimaryServerName.ValueString(),
		PrimaryServerType:           data.PrimaryServerType.ValueString(),
	}
	return aaaServerReqPayload
}

func externalServersFromConfigToRequest(ctx context.Context, externalServersData basetypes.ListValue) *[]externalServersReq {
	tflog.Debug(ctx, "In externalServersFromConfigToRequest Function")

	var externalServersReqList []externalServersReq

	for _, externalServersMap := range externalServersData.Elements() {

		externalServerObject := externalServersMap.(basetypes.ObjectValue)
		newExternalServers := externalServersReq{}
		for key, val := range externalServerObject.Attributes() {
			if !val.IsNull() {
				if key == "external_server_name" {
					newExternalServers.ExternalServerName = val.(basetypes.StringValue).ValueString()
				} else if key == "priority" {
					newExternalServers.Priority = val.(basetypes.Int64Value).ValueInt64()

				} else if key == "external_server_type" {
					newExternalServers.ExternalServerType = val.(basetypes.StringValue).ValueString()
				}
			}
		}
		externalServersReqList = append(externalServersReqList, newExternalServers)
	}

	return &externalServersReqList
}
func aaaServerSetAttrFromGet(ctx context.Context, data *aaaServerModel, getResponseData map[string]interface{}) *aaaServerModel {
	tflog.Debug(ctx, "In aaaServerSetAttrFromGet Function")
	data.ExternalServers = externalServerToTFValue(ctx, getResponseData["external_servers"].([]interface{}))
	val1, _ := strconv.ParseBool(getResponseData["fallback_local_authentication"].(string))
	data.FallbackLocalAuthentication = types.BoolValue(val1)
	val2, _ := strconv.ParseBool(getResponseData["log_ext_group_info"].(string))
	data.LogExtGroupInfo = types.BoolValue(val2)
	data.PrimaryServerName = types.StringValue(getResponseData["primary_server_name"].(string))
	data.PrimaryServerType = types.StringValue(getResponseData["primary_server_type"].(string))
	return data
}

func externalServerToTFValue(ctx context.Context, esListValue []interface{}) basetypes.ListValue {
	tflog.Debug(ctx, "In externalServerToTFValue Function")

	var esResList []attr.Value

	elementTypes := map[string]attr.Type{
		"external_server_name": types.StringType,
		"external_server_type": types.StringType,
		"priority":             types.Int64Type,
	}

	if len(esListValue) == 0 {
		objechType := basetypes.ObjectType{AttrTypes: elementTypes}
		listValuefinal := basetypes.NewListNull(objechType)
		return listValuefinal
	}

	for _, es := range esListValue {
		es := es.(map[string]interface{})
		elements := map[string]attr.Value{
			"external_server_name": basetypes.NewStringValue(es["external_server_name"].(string)),
			"external_server_type": basetypes.NewStringValue(es["external_server_type"].(string)),
		}

		val, _ := strconv.Atoi(es["priority"].(string))
		elements["priority"] = basetypes.NewInt64Value(int64(val))

		objectValue, _ := basetypes.NewObjectValue(elementTypes, elements)

		esResList = append(esResList, objectValue)
	}

	objechType := basetypes.ObjectType{AttrTypes: elementTypes}
	listValuefinal, _ := basetypes.NewListValueFrom(ctx, objechType, esResList)

	return listValuefinal
}

type aaaServerReq struct {
	ExternalServers             *[]externalServersReq `json:"external_servers,omitempty"`
	FallbackLocalAuthentication *bool                 `json:"fallback_local_authentication,omitempty"`
	LogExtGroupInfo             *bool                 `json:"log_ext_group_info,omitempty"`
	PrimaryServerName           string                `json:"primary_server_name,omitempty"`
	PrimaryServerType           string                `json:"primary_server_type,omitempty"`
}
type externalServersReq struct {
	ExternalServerType string `json:"external_server_type,omitempty"`
	ExternalServerName string `json:"external_server_name,omitempty"`
	Priority           int64  `json:"priority,omitempty"`
}

type aaaServerModel struct {
	ExternalServers             types.List   `tfsdk:"external_servers"`
	FallbackLocalAuthentication types.Bool   `tfsdk:"fallback_local_authentication"`
	LogExtGroupInfo             types.Bool   `tfsdk:"log_ext_group_info"`
	PrimaryServerName           types.String `tfsdk:"primary_server_name"`
	PrimaryServerType           types.String `tfsdk:"primary_server_type"`
	Id                          types.String `tfsdk:"id"`
}
