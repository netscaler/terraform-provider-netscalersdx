package radius_server

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func radiusServerResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Radius Server configuration resource.",
		Attributes: map[string]schema.Attribute{
			"accounting": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable accounting in the radius server.",
				MarkdownDescription: "Enable accounting in the radius server.",
			},
			"address_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Configuration Type. Values: 0: IPv4, 1: IPv6, -1: Hostname.",
				MarkdownDescription: "Configuration Type. Values: 0: IPv4, 1: IPv6, -1: Hostname.",
			},
			"attribute_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Attribute type for RADIUS group extraction. Maximum value =  ",
				MarkdownDescription: "Attribute type for RADIUS group extraction. Maximum value =  ",
			},
			"auth_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The maximum number of seconds the system will wait for a response from the Radius server. Maximum value =  ",
				MarkdownDescription: "The maximum number of seconds the system will wait for a response from the Radius server. Maximum value =  ",
			},
			"default_authentication_group": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64",
				MarkdownDescription: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64",
			},
			"enable_nas_ip": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable NAS IP extraction.",
				MarkdownDescription: "Enable NAS IP extraction.",
			},
			"group_separator": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         " Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  7",
				MarkdownDescription: " Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  7",
			},
			"groups_prefix": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  31",
				MarkdownDescription: "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction. Maximum length =  31",
			},
			"ip_address": schema.StringAttribute{
				Required:            true,
				Description:         "IP Address of radius server. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IP Address of radius server. Minimum length =  1 Maximum length =  64",
			},
			"ip_attribute_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The attribute type of the remote IP address attribute in a RADIUS response. Maximum value =  ",
				MarkdownDescription: "The attribute type of the remote IP address attribute in a RADIUS response. Maximum value =  ",
			},
			"ip_vendor_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The vendor ID of the attribute in the RADIUS response which denotes the intranet IP. Maximum value =  ",
				MarkdownDescription: "The vendor ID of the attribute in the RADIUS response which denotes the intranet IP. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Name of radius server. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of radius server. Minimum length =  1 Maximum length =  128",
			},
			"nas_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "NAS ID. Maximum length =  128",
				MarkdownDescription: "NAS ID. Maximum length =  128",
			},
			"pass_encoding": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable password encoding in RADIUS packets send to the RADIUS server.",
				MarkdownDescription: "Enable password encoding in RADIUS packets send to the RADIUS server.",
			},
			"port": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Port number of radius server. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Port number of radius server. Minimum value =  1 Maximum value =  ",
			},
			"pwd_attribute_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The attribute type of the password attribute in a RADIUS response.. Maximum value =  ",
				MarkdownDescription: "The attribute type of the password attribute in a RADIUS response.. Maximum value =  ",
			},
			"pwd_vendor_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Vendor ID of the password in the RADIUS response. Used to extract the user password. Maximum value =  ",
				MarkdownDescription: "Vendor ID of the password in the RADIUS response. Used to extract the user password. Maximum value =  ",
			},
			"radius_key": schema.StringAttribute{
				Required:            true,
				Description:         "Key of radius server. Minimum length =  4 Maximum length =  32",
				MarkdownDescription: "Key of radius server. Minimum length =  4 Maximum length =  32",
			},
			"vendor_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Vendor ID for RADIUS group extraction. Maximum value =  ",
				MarkdownDescription: "Vendor ID for RADIUS group extraction. Maximum value =  ",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type radiusServerModel struct {
	Accounting                 types.Bool   `tfsdk:"accounting"`
	AddressType                types.Int64  `tfsdk:"address_type"`
	AttributeType              types.Int64  `tfsdk:"attribute_type"`
	AuthTimeout                types.Int64  `tfsdk:"auth_timeout"`
	DefaultAuthenticationGroup types.String `tfsdk:"default_authentication_group"`
	EnableNasIp                types.Bool   `tfsdk:"enable_nas_ip"`
	GroupSeparator             types.String `tfsdk:"group_separator"`
	GroupsPrefix               types.String `tfsdk:"groups_prefix"`
	IpAddress                  types.String `tfsdk:"ip_address"`
	IpAttributeType            types.Int64  `tfsdk:"ip_attribute_type"`
	IpVendorId                 types.Int64  `tfsdk:"ip_vendor_id"`
	Name                       types.String `tfsdk:"name"`
	NasId                      types.String `tfsdk:"nas_id"`
	PassEncoding               types.String `tfsdk:"pass_encoding"`
	Port                       types.Int64  `tfsdk:"port"`
	PwdAttributeType           types.Int64  `tfsdk:"pwd_attribute_type"`
	PwdVendorId                types.Int64  `tfsdk:"pwd_vendor_id"`
	RadiusKey                  types.String `tfsdk:"radius_key"`
	VendorId                   types.Int64  `tfsdk:"vendor_id"`
	Id                         types.String `tfsdk:"id"`
}

func radiusServerGetThePayloadFromtheConfig(ctx context.Context, data *radiusServerModel) radiusServerReq {
	tflog.Debug(ctx, "In radiusServerGetThePayloadFromtheConfig Function")
	radiusServerReqPayload := radiusServerReq{
		DefaultAuthenticationGroup: data.DefaultAuthenticationGroup.ValueString(),
		GroupSeparator:             data.GroupSeparator.ValueString(),
		GroupsPrefix:               data.GroupsPrefix.ValueString(),
		IpAddress:                  data.IpAddress.ValueString(),
		Name:                       data.Name.ValueString(),
		NasId:                      data.NasId.ValueString(),
		PassEncoding:               data.PassEncoding.ValueString(),
		RadiusKey:                  data.RadiusKey.ValueString(),
	}

	if !data.Accounting.IsNull() && !data.Accounting.IsUnknown() {
		radiusServerReqPayload.Accounting = data.Accounting.ValueBoolPointer()
	}
	if !data.AddressType.IsNull() && !data.AddressType.IsUnknown() {
		radiusServerReqPayload.AddressType = data.AddressType.ValueInt64Pointer()
	}
	if !data.AttributeType.IsNull() && !data.AttributeType.IsUnknown() {
		radiusServerReqPayload.AttributeType = data.AttributeType.ValueInt64Pointer()
	}
	if !data.AuthTimeout.IsNull() && !data.AuthTimeout.IsUnknown() {
		radiusServerReqPayload.AuthTimeout = data.AuthTimeout.ValueInt64Pointer()
	}
	if !data.EnableNasIp.IsNull() && !data.EnableNasIp.IsUnknown() {
		radiusServerReqPayload.EnableNasIp = data.EnableNasIp.ValueBoolPointer()
	}
	if !data.IpAttributeType.IsNull() && !data.IpAttributeType.IsUnknown() {
		radiusServerReqPayload.IpAttributeType = data.IpAttributeType.ValueInt64Pointer()
	}
	if !data.IpVendorId.IsNull() && !data.IpVendorId.IsUnknown() {
		radiusServerReqPayload.IpVendorId = data.IpVendorId.ValueInt64Pointer()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		radiusServerReqPayload.Port = data.Port.ValueInt64Pointer()
	}
	if !data.PwdAttributeType.IsNull() && !data.PwdAttributeType.IsUnknown() {
		radiusServerReqPayload.PwdAttributeType = data.PwdAttributeType.ValueInt64Pointer()
	}
	if !data.PwdVendorId.IsNull() && !data.PwdVendorId.IsUnknown() {
		radiusServerReqPayload.PwdVendorId = data.PwdVendorId.ValueInt64Pointer()
	}
	if !data.VendorId.IsNull() && !data.VendorId.IsUnknown() {
		radiusServerReqPayload.VendorId = data.VendorId.ValueInt64Pointer()
	}

	return radiusServerReqPayload
}

type radiusServerReq struct {
	Accounting                 *bool  `json:"accounting,omitempty"`
	AddressType                *int64 `json:"address_type,omitempty"`
	AttributeType              *int64 `json:"attribute_type,omitempty"`
	AuthTimeout                *int64 `json:"auth_timeout,omitempty"`
	DefaultAuthenticationGroup string `json:"default_authentication_group,omitempty"`
	EnableNasIp                *bool  `json:"enable_nas_ip,omitempty"`
	GroupSeparator             string `json:"group_separator,omitempty"`
	GroupsPrefix               string `json:"groups_prefix,omitempty"`
	IpAddress                  string `json:"ip_address,omitempty"`
	IpAttributeType            *int64 `json:"ip_attribute_type,omitempty"`
	IpVendorId                 *int64 `json:"ip_vendor_id,omitempty"`
	Name                       string `json:"name,omitempty"`
	NasId                      string `json:"nas_id,omitempty"`
	PassEncoding               string `json:"pass_encoding,omitempty"`
	Port                       *int64 `json:"port,omitempty"`
	PwdAttributeType           *int64 `json:"pwd_attribute_type,omitempty"`
	PwdVendorId                *int64 `json:"pwd_vendor_id,omitempty"`
	RadiusKey                  string `json:"radius_key,omitempty"`
	VendorId                   *int64 `json:"vendor_id,omitempty"`
}

func radiusServerSetAttrFromGet(ctx context.Context, data *radiusServerModel, getResponseData map[string]interface{}) *radiusServerModel {
	tflog.Debug(ctx, "In radiusServerSetAttrFromGet Function")

	data.Accounting = types.BoolValue(utils.StringToBool(getResponseData["accounting"].(string)))
	data.AddressType = types.Int64Value(utils.StringToInt(getResponseData["address_type"].(string)))
	data.AttributeType = types.Int64Value(utils.StringToInt(getResponseData["attribute_type"].(string)))
	data.AuthTimeout = types.Int64Value(utils.StringToInt(getResponseData["auth_timeout"].(string)))
	data.DefaultAuthenticationGroup = types.StringValue(getResponseData["default_authentication_group"].(string))
	data.EnableNasIp = types.BoolValue(utils.StringToBool(getResponseData["enable_nas_ip"].(string)))
	data.GroupSeparator = types.StringValue(getResponseData["group_separator"].(string))
	data.GroupsPrefix = types.StringValue(getResponseData["groups_prefix"].(string))
	data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	data.IpAttributeType = types.Int64Value(utils.StringToInt(getResponseData["ip_attribute_type"].(string)))
	data.IpVendorId = types.Int64Value(utils.StringToInt(getResponseData["ip_vendor_id"].(string)))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.NasId = types.StringValue(getResponseData["nas_id"].(string))
	data.PassEncoding = types.StringValue(getResponseData["pass_encoding"].(string))
	data.Port = types.Int64Value(utils.StringToInt(getResponseData["port"].(string)))
	data.PwdAttributeType = types.Int64Value(utils.StringToInt(getResponseData["pwd_attribute_type"].(string)))
	data.PwdVendorId = types.Int64Value(utils.StringToInt(getResponseData["pwd_vendor_id"].(string)))
	data.VendorId = types.Int64Value(utils.StringToInt(getResponseData["vendor_id"].(string)))

	return data
}
