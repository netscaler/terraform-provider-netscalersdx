package ns_device_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nsDeviceProfileResourceModel describes the resource data model.
type nsDeviceProfileResourceModel struct {
	Name                                   types.String `tfsdk:"name"`
	SvmNsComm                              types.String `tfsdk:"svm_ns_comm"`
	UseGlobalSettingForCommunicationWithNs types.Bool   `tfsdk:"use_global_setting_for_communication_with_ns"`
	Id                                     types.String `tfsdk:"id"`
	Type                                   types.String `tfsdk:"type"`
	NsProfileName                          types.String `tfsdk:"ns_profile_name"`
	Password                               types.String `tfsdk:"password"`
	Snmpsecuritylevel                      types.String `tfsdk:"snmpsecuritylevel"`
	Username                               types.String `tfsdk:"username"`
	Snmpauthprotocol                       types.String `tfsdk:"snmpauthprotocol"`
	SshPort                                types.String `tfsdk:"ssh_port"`
	Snmpprivprotocol                       types.String `tfsdk:"snmpprivprotocol"`
	HostPassword                           types.String `tfsdk:"host_password"`
	Snmpversion                            types.String `tfsdk:"snmpversion"`
	Passphrase                             types.String `tfsdk:"passphrase"`
	HostUsername                           types.String `tfsdk:"host_username"`
	Snmpsecurityname                       types.String `tfsdk:"snmpsecurityname"`
	SslPrivateKey                          types.String `tfsdk:"ssl_private_key"`
	SslCert                                types.String `tfsdk:"ssl_cert"`
	HttpPort                               types.Int64  `tfsdk:"http_port"`
	Snmpcommunity                          types.String `tfsdk:"snmpcommunity"`
	HttpsPort                              types.Int64  `tfsdk:"https_port"`
	MaxWaitTimeReboot                      types.String `tfsdk:"max_wait_time_reboot"`
	Snmpprivpassword                       types.String `tfsdk:"snmpprivpassword"`
	CbProfileName                          types.String `tfsdk:"cb_profile_name"`
	Snmpauthpassword                       types.String `tfsdk:"snmpauthpassword"`
}

type nsDeviceProfileResourceReq struct {
	Name                                   string `json:"name,omitempty"`
	SvmNsComm                              string `json:"svm_ns_comm,omitempty"`
	UseGlobalSettingForCommunicationWithNs bool   `json:"use_global_setting_for_communication_with_ns"`
	Id                                     string `json:"id,omitempty"`
	Type                                   string `json:"type,omitempty"`
	NsProfileName                          string `json:"ns_profile_name,omitempty"`
	Password                               string `json:"password,omitempty"`
	Snmpsecuritylevel                      string `json:"snmpsecuritylevel,omitempty"`
	Username                               string `json:"username,omitempty"`
	Snmpauthprotocol                       string `json:"snmpauthprotocol,omitempty"`
	SshPort                                string `json:"ssh_port,omitempty"`
	Snmpprivprotocol                       string `json:"snmpprivprotocol,omitempty"`
	HostPassword                           string `json:"host_password,omitempty"`
	Snmpversion                            string `json:"snmpversion,omitempty"`
	Passphrase                             string `json:"passphrase,omitempty"`
	HostUsername                           string `json:"host_username,omitempty"`
	Snmpsecurityname                       string `json:"snmpsecurityname,omitempty"`
	SslPrivateKey                          string `json:"ssl_private_key,omitempty"`
	SslCert                                string `json:"ssl_cert,omitempty"`
	HttpPort                               int64  `json:"http_port,omitempty"`
	Snmpcommunity                          string `json:"snmpcommunity,omitempty"`
	HttpsPort                              int64  `json:"https_port,omitempty"`
	MaxWaitTimeReboot                      string `json:"max_wait_time_reboot,omitempty"`
	Snmpprivpassword                       string `json:"snmpprivpassword,omitempty"`
	CbProfileName                          string `json:"cb_profile_name,omitempty"`
	Snmpauthpassword                       string `json:"snmpauthpassword,omitempty"`
}

// function that returns the request body for the API call.
func nsDeviceProfileGetThePayloadFromtheConfig(ctx context.Context, data *nsDeviceProfileResourceModel) nsDeviceProfileResourceReq {
	tflog.Debug(ctx, "In nsDeviceProfileGetThePayloadFromtheConfig Function of nsDeviceProfileResource")

	nsDeviceProfileReq := nsDeviceProfileResourceReq{
		Name:                                   data.Name.ValueString(),
		SvmNsComm:                              data.SvmNsComm.ValueString(),
		UseGlobalSettingForCommunicationWithNs: data.UseGlobalSettingForCommunicationWithNs.ValueBool(),
		Type:                                   data.Type.ValueString(),
		NsProfileName:                          data.NsProfileName.ValueString(),
		Password:                               data.Password.ValueString(),
		Snmpsecuritylevel:                      data.Snmpsecuritylevel.ValueString(),
		Username:                               data.Username.ValueString(),
		Snmpauthprotocol:                       data.Snmpauthprotocol.ValueString(),
		SshPort:                                data.SshPort.ValueString(),
		Snmpprivprotocol:                       data.Snmpprivprotocol.ValueString(),
		HostPassword:                           data.HostPassword.ValueString(),
		Snmpversion:                            data.Snmpversion.ValueString(),
		Passphrase:                             data.Passphrase.ValueString(),
		HostUsername:                           data.HostUsername.ValueString(),
		Snmpsecurityname:                       data.Snmpsecurityname.ValueString(),
		SslPrivateKey:                          data.SslPrivateKey.ValueString(),
		SslCert:                                data.SslCert.ValueString(),
		HttpPort:                               data.HttpPort.ValueInt64(),
		Snmpcommunity:                          data.Snmpcommunity.ValueString(),
		HttpsPort:                              data.HttpsPort.ValueInt64(),
		MaxWaitTimeReboot:                      data.MaxWaitTimeReboot.ValueString(),
		Snmpprivpassword:                       data.Snmpprivpassword.ValueString(),
		CbProfileName:                          data.CbProfileName.ValueString(),
		Snmpauthpassword:                       data.Snmpauthpassword.ValueString(),
	}

	return nsDeviceProfileReq
}
