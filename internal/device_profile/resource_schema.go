package device_profile

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func deviceProfileResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Device Profile resource.",
		Attributes: map[string]schema.Attribute{
			"cb_profile_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Profile Name, This is one of the already created Citrix SD-WAN profiles.",
				MarkdownDescription: "Profile Name, This is one of the already created Citrix SD-WAN profiles.",
			},
			"host_password": schema.StringAttribute{
				Optional:            true,
				Description:         "Host Password for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Host Password for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"host_username": schema.StringAttribute{
				Optional:            true,
				Description:         "Host User Name for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Host User Name for this profile.Used for BLX form factor of ADC. Minimum length =  1 Maximum length =  127",
			},
			"http_port": schema.Int64Attribute{
				Optional:            true,
				Description:         "HTTP port to connect to the device.",
				MarkdownDescription: "HTTP port to connect to the device.",
			},
			"https_port": schema.Int64Attribute{
				Optional:            true,
				Description:         "HTTPS port to connect to the device.",
				MarkdownDescription: "HTTPS port to connect to the device.",
			},
			"max_wait_time_reboot": schema.StringAttribute{
				Optional:            true,
				Description:         "Max waiting time to reboot Citrix ADC.",
				MarkdownDescription: "Max waiting time to reboot Citrix ADC.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Profile Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Profile Name. Minimum length =  1 Maximum length =  128",
			},
			"ns_profile_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Profile Name, This is one of the already created Citrix ADC profiles.",
				MarkdownDescription: "Profile Name, This is one of the already created Citrix ADC profiles.",
			},
			"passphrase": schema.StringAttribute{
				Optional:            true,
				Description:         "Passphrase with which private key is encrypted.",
				MarkdownDescription: "Passphrase with which private key is encrypted.",
			},
			"password": schema.StringAttribute{
				Required:            true,
				Description:         "Instance credentials.Password for this profile. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Instance credentials.Password for this profile. Minimum length =  1 Maximum length =  127",
			},
			"snmpauthpassword": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 auth password for this profile. Minimum length =  8 Maximum length =  31",
				MarkdownDescription: "SNMP v3 auth password for this profile. Minimum length =  8 Maximum length =  31",
			},
			"snmpauthprotocol": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 auth protocol for this profile.",
				MarkdownDescription: "SNMP v3 auth protocol for this profile.",
			},
			"snmpcommunity": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP community for this profile. Maximum length =  31",
				MarkdownDescription: "SNMP community for this profile. Maximum length =  31",
			},
			"snmpprivpassword": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 priv password for this profile. Minimum length =  8 Maximum length =  31",
				MarkdownDescription: "SNMP v3 priv password for this profile. Minimum length =  8 Maximum length =  31",
			},
			"snmpprivprotocol": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 priv protocol for this profile.",
				MarkdownDescription: "SNMP v3 priv protocol for this profile.",
			},
			"snmpsecuritylevel": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 security level for this profile.",
				MarkdownDescription: "SNMP v3 security level for this profile.",
			},
			"snmpsecurityname": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP v3 security name for this profile. Maximum length =  31",
				MarkdownDescription: "SNMP v3 security name for this profile. Maximum length =  31",
			},
			"snmpversion": schema.StringAttribute{
				Optional:            true,
				Description:         "SNMP version for this profile.",
				MarkdownDescription: "SNMP version for this profile.",
			},
			"ssh_port": schema.StringAttribute{
				Optional:            true,
				Description:         "SSH port to connect to the device.",
				MarkdownDescription: "SSH port to connect to the device.",
			},
			"ssl_cert": schema.StringAttribute{
				Optional:            true,
				Description:         "SSL Certificate for certificate based authentication.",
				MarkdownDescription: "SSL Certificate for certificate based authentication.",
			},
			"ssl_private_key": schema.StringAttribute{
				Optional:            true,
				Description:         "SSL Private Key for key based authentication.",
				MarkdownDescription: "SSL Private Key for key based authentication.",
			},
			"svm_ns_comm": schema.StringAttribute{
				Optional:            true,
				Description:         "Communication protocol (http or https) with Instances. Minimum length =  1 Maximum length =  10",
				MarkdownDescription: "Communication protocol (http or https) with Instances. Minimum length =  1 Maximum length =  10",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Description:         "Profile Type, This must be with in specified supported instance types: blx,sdvanvw,ns,nssdx,cbwanopt,cpx. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Profile Type, This must be with in specified supported instance types: blx,sdvanvw,ns,nssdx,cbwanopt,cpx. Minimum length =  1 Maximum length =  128",
			},
			"use_global_setting_for_communication_with_ns": schema.BoolAttribute{
				Optional:            true,
				Description:         "True, if the communication with Instance needs to be global and not device specific.",
				MarkdownDescription: "True, if the communication with Instance needs to be global and not device specific.",
			},
			"username": schema.StringAttribute{
				Required:            true,
				Description:         "Instance credentials.Username provided in the profile will be used to contact the instance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Instance credentials.Username provided in the profile will be used to contact the instance. Minimum length =  1 Maximum length =  127",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type deviceProfileModel struct {
	CbProfileName                          types.String `tfsdk:"cb_profile_name"`
	HostPassword                           types.String `tfsdk:"host_password"`
	HostUsername                           types.String `tfsdk:"host_username"`
	HttpPort                               types.Int64  `tfsdk:"http_port"`
	HttpsPort                              types.Int64  `tfsdk:"https_port"`
	MaxWaitTimeReboot                      types.String `tfsdk:"max_wait_time_reboot"`
	Name                                   types.String `tfsdk:"name"`
	NsProfileName                          types.String `tfsdk:"ns_profile_name"`
	Passphrase                             types.String `tfsdk:"passphrase"`
	Password                               types.String `tfsdk:"password"`
	Snmpauthpassword                       types.String `tfsdk:"snmpauthpassword"`
	Snmpauthprotocol                       types.String `tfsdk:"snmpauthprotocol"`
	Snmpcommunity                          types.String `tfsdk:"snmpcommunity"`
	Snmpprivpassword                       types.String `tfsdk:"snmpprivpassword"`
	Snmpprivprotocol                       types.String `tfsdk:"snmpprivprotocol"`
	Snmpsecuritylevel                      types.String `tfsdk:"snmpsecuritylevel"`
	Snmpsecurityname                       types.String `tfsdk:"snmpsecurityname"`
	Snmpversion                            types.String `tfsdk:"snmpversion"`
	SshPort                                types.String `tfsdk:"ssh_port"`
	SslCert                                types.String `tfsdk:"ssl_cert"`
	SslPrivateKey                          types.String `tfsdk:"ssl_private_key"`
	SvmNsComm                              types.String `tfsdk:"svm_ns_comm"`
	Type                                   types.String `tfsdk:"type"`
	UseGlobalSettingForCommunicationWithNs types.Bool   `tfsdk:"use_global_setting_for_communication_with_ns"`
	Username                               types.String `tfsdk:"username"`
	Id                                     types.String `tfsdk:"id"`
}

func deviceProfileGetThePayloadFromtheConfig(ctx context.Context, data *deviceProfileModel) deviceProfileReq {
	tflog.Debug(ctx, "In deviceProfileGetThePayloadFromtheConfig Function")
	deviceProfileReqPayload := deviceProfileReq{
		CbProfileName:                          data.CbProfileName.ValueString(),
		HostPassword:                           data.HostPassword.ValueString(),
		HostUsername:                           data.HostUsername.ValueString(),
		HttpPort:                               data.HttpPort.ValueInt64(),
		HttpsPort:                              data.HttpsPort.ValueInt64(),
		MaxWaitTimeReboot:                      data.MaxWaitTimeReboot.ValueString(),
		Name:                                   data.Name.ValueString(),
		NsProfileName:                          data.NsProfileName.ValueString(),
		Passphrase:                             data.Passphrase.ValueString(),
		Password:                               data.Password.ValueString(),
		Snmpauthpassword:                       data.Snmpauthpassword.ValueString(),
		Snmpauthprotocol:                       data.Snmpauthprotocol.ValueString(),
		Snmpcommunity:                          data.Snmpcommunity.ValueString(),
		Snmpprivpassword:                       data.Snmpprivpassword.ValueString(),
		Snmpprivprotocol:                       data.Snmpprivprotocol.ValueString(),
		Snmpsecuritylevel:                      data.Snmpsecuritylevel.ValueString(),
		Snmpsecurityname:                       data.Snmpsecurityname.ValueString(),
		Snmpversion:                            data.Snmpversion.ValueString(),
		SshPort:                                data.SshPort.ValueString(),
		SslCert:                                data.SslCert.ValueString(),
		SslPrivateKey:                          data.SslPrivateKey.ValueString(),
		SvmNsComm:                              data.SvmNsComm.ValueString(),
		Type:                                   data.Type.ValueString(),
		UseGlobalSettingForCommunicationWithNs: data.UseGlobalSettingForCommunicationWithNs.ValueBool(),
		Username:                               data.Username.ValueString(),
	}
	return deviceProfileReqPayload
}

// func deviceProfileSetAttrFromGet(ctx context.Context, data *deviceProfileModel, getResponseData map[string]interface{}) *deviceProfileModel {
// 	tflog.Debug(ctx, "In deviceProfileSetAttrFromGet Function")
// 	return data
// }

type deviceProfileReq struct {
	CbProfileName                          string `json:"cb_profile_name,omitempty"`
	HostPassword                           string `json:"host_password,omitempty"`
	HostUsername                           string `json:"host_username,omitempty"`
	HttpPort                               int64  `json:"http_port,omitempty"`
	HttpsPort                              int64  `json:"https_port,omitempty"`
	MaxWaitTimeReboot                      string `json:"max_wait_time_reboot,omitempty"`
	Name                                   string `json:"name,omitempty"`
	NsProfileName                          string `json:"ns_profile_name,omitempty"`
	Passphrase                             string `json:"passphrase,omitempty"`
	Password                               string `json:"password,omitempty"`
	Snmpauthpassword                       string `json:"snmpauthpassword,omitempty"`
	Snmpauthprotocol                       string `json:"snmpauthprotocol,omitempty"`
	Snmpcommunity                          string `json:"snmpcommunity,omitempty"`
	Snmpprivpassword                       string `json:"snmpprivpassword,omitempty"`
	Snmpprivprotocol                       string `json:"snmpprivprotocol,omitempty"`
	Snmpsecuritylevel                      string `json:"snmpsecuritylevel,omitempty"`
	Snmpsecurityname                       string `json:"snmpsecurityname,omitempty"`
	Snmpversion                            string `json:"snmpversion,omitempty"`
	SshPort                                string `json:"ssh_port,omitempty"`
	SslCert                                string `json:"ssl_cert,omitempty"`
	SslPrivateKey                          string `json:"ssl_private_key,omitempty"`
	SvmNsComm                              string `json:"svm_ns_comm,omitempty"`
	Type                                   string `json:"type,omitempty"`
	UseGlobalSettingForCommunicationWithNs bool   `json:"use_global_setting_for_communication_with_ns,omitempty"`
	Username                               string `json:"username,omitempty"`
}
