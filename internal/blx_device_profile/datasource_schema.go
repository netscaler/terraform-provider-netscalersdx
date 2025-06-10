package blx_device_profile

import (
	"context"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func blxDeviceProfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a blx device profile.",
		Attributes: map[string]schema.Attribute{
			"cb_profile_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Profile Name, This is one of the already created NetScaler SD-WAN profiles.",
				MarkdownDescription: "Profile Name, This is one of the already created NetScaler SD-WAN profiles.",
			},
			"host_password": schema.StringAttribute{
				Computed:            true,
				Description:         "Host Password for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Host Password for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"host_username": schema.StringAttribute{
				Computed:            true,
				Description:         "Host User Name for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Host User Name for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"http_port": schema.Int64Attribute{
				Computed:            true,
				Description:         "HTTP port to connect to the device.",
				MarkdownDescription: "HTTP port to connect to the device.",
			},
			"https_port": schema.Int64Attribute{
				Computed:            true,
				Description:         "HTTPS port to connect to the device.",
				MarkdownDescription: "HTTPS port to connect to the device.",
			},
			"max_wait_time_reboot": schema.StringAttribute{
				Computed:            true,
				Description:         "Max waiting time to reboot NetScaler.",
				MarkdownDescription: "Max waiting time to reboot NetScaler.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Profile Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Profile Name. Minimum length =  1 Maximum length =  128",
			},
			"ns_profile_name": schema.StringAttribute{
				Computed:            true,
				Description:         "Profile Name, This is one of the already created NetScaler profiles.",
				MarkdownDescription: "Profile Name, This is one of the already created NetScaler profiles.",
			},
			"passphrase": schema.StringAttribute{
				Computed:            true,
				Description:         "Passphrase with which private key is encrypted.",
				MarkdownDescription: "Passphrase with which private key is encrypted.",
			},
			"password": schema.StringAttribute{
				Computed:            true,
				Description:         "Instance credentials.Password for this profile. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Instance credentials.Password for this profile. Minimum length =  1 Maximum length =  127",
			},
			"snmpauthpassword": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 auth password for this profile. Minimum length =  8 Maximum length =  31",
				MarkdownDescription: "SNMP v3 auth password for this profile. Minimum length =  8 Maximum length =  31",
			},
			"snmpauthprotocol": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 auth protocol for this profile.",
				MarkdownDescription: "SNMP v3 auth protocol for this profile.",
			},
			"snmpcommunity": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP community for this profile. Maximum length =  31",
				MarkdownDescription: "SNMP community for this profile. Maximum length =  31",
			},
			"snmpprivpassword": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 priv password for this profile. Minimum length =  8 Maximum length =  31",
				MarkdownDescription: "SNMP v3 priv password for this profile. Minimum length =  8 Maximum length =  31",
			},
			"snmpprivprotocol": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 priv protocol for this profile.",
				MarkdownDescription: "SNMP v3 priv protocol for this profile.",
			},
			"snmpsecuritylevel": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 security level for this profile.",
				MarkdownDescription: "SNMP v3 security level for this profile.",
			},
			"snmpsecurityname": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP v3 security name for this profile. Maximum length =  31",
				MarkdownDescription: "SNMP v3 security name for this profile. Maximum length =  31",
			},
			"snmpversion": schema.StringAttribute{
				Computed:            true,
				Description:         "SNMP version for this profile.",
				MarkdownDescription: "SNMP version for this profile.",
			},
			"ssh_port": schema.StringAttribute{
				Computed:            true,
				Description:         "SSH port to connect to the device.",
				MarkdownDescription: "SSH port to connect to the device.",
			},
			"ssl_cert": schema.StringAttribute{
				Computed:            true,
				Description:         "SSL Certificate for certificate based authentication.",
				MarkdownDescription: "SSL Certificate for certificate based authentication.",
			},
			"ssl_private_key": schema.StringAttribute{
				Computed:            true,
				Description:         "SSL Private Key for key based authentication.",
				MarkdownDescription: "SSL Private Key for key based authentication.",
			},
			"svm_ns_comm": schema.StringAttribute{
				Computed:            true,
				Description:         "Communication protocol (http or https) with Instances. Minimum length =  1 Maximum length =  10",
				MarkdownDescription: "Communication protocol (http or https) with Instances. Minimum length =  1 Maximum length =  10",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				Description:         "Profile Type, This must be with in specified supported instance types: blx,ns,nssdx,cpx. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Profile Type, This must be with in specified supported instance types: blx,ns,nssdx,cpx. Minimum length =  1 Maximum length =  128",
			},
			"use_global_setting_for_communication_with_ns": schema.BoolAttribute{
				Computed:            true,
				Description:         "True, if the communication with Instance needs to be global and not device specific.",
				MarkdownDescription: "True, if the communication with Instance needs to be global and not device specific.",
			},
			"username": schema.StringAttribute{
				Computed:            true,
				Description:         "Instance credentials.Username provided in the profile will be used to contact the instance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Instance credentials.Username provided in the profile will be used to contact the instance. Minimum length =  1 Maximum length =  127",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this resource",
			},
		},
	}
}

func blxDeviceProfileSetAttrFromGet(ctx context.Context, data *blxDeviceProfileModel, getResponseData map[string]interface{}) *blxDeviceProfileModel {

	data.CbProfileName = types.StringValue(getResponseData["cb_profile_name"].(string))
	data.HostUsername = types.StringValue(getResponseData["host_username"].(string))
	data.HttpPort = types.Int64Value(utils.StringToInt(getResponseData["http_port"].(string)))
	data.HttpsPort = types.Int64Value(utils.StringToInt(getResponseData["https_port"].(string)))
	data.MaxWaitTimeReboot = types.StringValue(getResponseData["max_wait_time_reboot"].(string))
	data.Name = types.StringValue(getResponseData["name"].(string))
	data.NsProfileName = types.StringValue(getResponseData["ns_profile_name"].(string))
	data.Snmpauthprotocol = types.StringValue(getResponseData["snmpauthprotocol"].(string))
	data.Snmpcommunity = types.StringValue(getResponseData["snmpcommunity"].(string))
	data.Snmpprivprotocol = types.StringValue(getResponseData["snmpprivprotocol"].(string))
	data.Snmpsecuritylevel = types.StringValue(getResponseData["snmpsecuritylevel"].(string))
	data.Snmpsecurityname = types.StringValue(getResponseData["snmpsecurityname"].(string))
	data.Snmpversion = types.StringValue(getResponseData["snmpversion"].(string))
	data.SshPort = types.StringValue(getResponseData["ssh_port"].(string))
	data.SslCert = types.StringValue(getResponseData["ssl_cert"].(string))
	data.SvmNsComm = types.StringValue(getResponseData["svm_ns_comm"].(string))
	data.Type = types.StringValue(getResponseData["type"].(string))
	data.UseGlobalSettingForCommunicationWithNs = types.BoolValue(utils.StringToBool(getResponseData["use_global_setting_for_communication_with_ns"].(string)))
	data.Username = types.StringValue(getResponseData["username"].(string))

	return data
}
