package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"terraform-provider-netscalersdx/internal/aaa_server"
	"terraform-provider-netscalersdx/internal/aclrule"
	"terraform-provider-netscalersdx/internal/blx_device_profile"
	"terraform-provider-netscalersdx/internal/cipher_config"
	"terraform-provider-netscalersdx/internal/cipher_group"
	"terraform-provider-netscalersdx/internal/current_hostname"
	"terraform-provider-netscalersdx/internal/current_timezone"
	"terraform-provider-netscalersdx/internal/device_group"
	"terraform-provider-netscalersdx/internal/device_profile"
	"terraform-provider-netscalersdx/internal/ldap_server"
	"terraform-provider-netscalersdx/internal/mps"
	"terraform-provider-netscalersdx/internal/mps_feature"
	"terraform-provider-netscalersdx/internal/mps_ssl_certkey"
	"terraform-provider-netscalersdx/internal/mpsgroup"
	"terraform-provider-netscalersdx/internal/mpsuser"
	"terraform-provider-netscalersdx/internal/ns"
	"terraform-provider-netscalersdx/internal/ns_device_profile"
	"terraform-provider-netscalersdx/internal/ns_save_config"
	"terraform-provider-netscalersdx/internal/ntp_param"
	"terraform-provider-netscalersdx/internal/ntp_server"
	"terraform-provider-netscalersdx/internal/ntp_sync"
	"terraform-provider-netscalersdx/internal/radius_server"
	"terraform-provider-netscalersdx/internal/sdx_license"
	"terraform-provider-netscalersdx/internal/smtp_server"
	"terraform-provider-netscalersdx/internal/snmp_alarm_config"
	"terraform-provider-netscalersdx/internal/snmp_manager"
	"terraform-provider-netscalersdx/internal/snmp_mib"
	"terraform-provider-netscalersdx/internal/snmp_trap"
	"terraform-provider-netscalersdx/internal/snmp_user"
	"terraform-provider-netscalersdx/internal/snmp_view"
	"terraform-provider-netscalersdx/internal/ssl_settings"
	"terraform-provider-netscalersdx/internal/static_route"
	"terraform-provider-netscalersdx/internal/syslog_params"
	"terraform-provider-netscalersdx/internal/syslog_server"
	"terraform-provider-netscalersdx/internal/system_settings"
	"terraform-provider-netscalersdx/internal/tacacs_server"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ provider.Provider = &sdxprovider{}
)

// Provider -
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &sdxprovider{
			version: version,
		}
	}
}

type sdxprovider struct {
	version string
}

func (p *sdxprovider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netscalersdx"
	resp.Version = p.version
}

func (p *sdxprovider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:    true,
				Description: "NetScaler SDX host. Can be specified with `NETSCALERSDX_HOST` environment variable. This has to start with https://",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "NetScaler SDX username. Can be specified with `NETSCALERSDX_USERNAME` environment variable.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "NetScaler SDX password. Can be specified with `NETSCALERSDX_PASSWORD` environment variable.",
			},
			"ssl_verify": schema.BoolAttribute{
				Optional:    true,
				Description: "Ignore validity of SDX TLS certificate if true. Can be specified with `NETSCALERSDX_SSL_VERIFY` environment variable.",
			},
			"log_level": schema.StringAttribute{
				Optional:    true,
				Description: "Log level (Default is INFO). Can be specified with `NETSCALERSDX_LOG_LEVEL` environment variable.",
			},
			"server_name": schema.StringAttribute{
				Optional:    true,
				Description: "TODO",
			},
			"headers": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "TODO",
			},
			"json_log_format": schema.BoolAttribute{
				Optional:    true,
				Description: "TODO",
			},
			"root_ca_path": schema.StringAttribute{
				Optional:    true,
				Description: "TODO",
			},
		},
	}
}

func (p *sdxprovider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		ns.ProvisionVpxResource,
		ns.VpxStateResource,
		ns_device_profile.NsDeviceProfileResource,
		ntp_server.NtpServerResource,
		ns_save_config.NsSaveConfigResource,
		cipher_group.CipherGroupResource,
		syslog_server.SyslogServerResource,
		static_route.StaticRouteResource,
		smtp_server.SmtpServerResource,
		tacacs_server.TacacsServerResource,
		ldap_server.LdapServerResource,
		device_profile.DeviceProfileResource,
		radius_server.RadiusServerResource,
		device_group.DeviceGroupResource,
		blx_device_profile.BlxDeviceProfileResource,
		snmp_user.SnmpUserResource,
		syslog_params.SyslogParamsResource,
		ntp_param.NtpParamResource,
		snmp_view.SnmpViewResource,
		mps_feature.MpsFeatureResource,
		current_timezone.CurrentTimezoneResource,
		snmp_alarm_config.SnmpAlarmConfigResource,
		snmp_trap.SnmpTrapResource,
		snmp_manager.SnmpManagerResource,
		mpsuser.MpsuserResource,
		mpsgroup.MpsgroupResource,
		aclrule.AclruleResource,
		aaa_server.AaaServerResource,
		snmp_mib.SnmpMibResource,
		ntp_sync.NtpSyncResource,
		mps.MpsResource,
		cipher_config.CipherConfigResource,
		ssl_settings.SslSettingsResource,
		system_settings.SystemSettingsResource,
		current_hostname.CurrentHostnameResource,
		mps_ssl_certkey.MpsSslCertkeyResource,
		sdx_license.LicenseFileResource,
	}
}
func (p *sdxprovider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		ns.VpxDataSource,
		aaa_server.AaaServerDataSource,
		cipher_config.CipherConfigDataSource,
		cipher_group.CipherGroupDataSource,
		radius_server.RadiusServerDataSource,
		device_group.DeviceGroupDataSource,
		ldap_server.LdapServerDataSource,
		mps.MpsDataSource,
		mps_feature.MpsFeatureDataSource,
		mpsgroup.MpsgroupDataSource,
		mpsuser.MpsuserDataSource,
	}
}

// Configure prepares an API client for data sources and resources.
func (p *sdxprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var params service.NitroParams

	diags := req.Config.Get(ctx, &params)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Use values in environment variables if the values aren't set in the configuration.
	hostEnv := os.Getenv("NETSCALERSDX_HOST")
	if params.Host.IsNull() && hostEnv != "" {
		params.Host = types.StringValue(hostEnv)
	}
	usernameEnv := os.Getenv("NETSCALERSDX_USERNAME")
	if params.Username.IsNull() && usernameEnv != "" {
		params.Username = types.StringValue(usernameEnv)
	}
	passwordEnv := os.Getenv("NETSCALERSDX_PASSWORD")
	if params.Password.IsNull() && passwordEnv != "" {
		params.Password = types.StringValue(passwordEnv)
	}
	if params.SslVerify.IsNull() {
		envSslVerify, _ := strconv.ParseBool(os.Getenv("NETSCALERSDX_SSL_VERIFY"))
		params.SslVerify = types.BoolValue(envSslVerify)
	}
	if params.LogLevel.IsNull() {
		params.LogLevel = types.StringValue(os.Getenv("NETSCALERSDX_LOG_LEVEL"))
	}

	if params.Host.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"host is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Host. Either set the value statically in the configuration, or use the NETSCALERSDX_HOST "+
				"environment variable."),
		)
	}
	if params.Username.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"username is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Username. Either set the value statically in the configuration, or use the NETSCALERSDX_USERNAME "+
				"environment variable."),
		)
	}
	if params.Password.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"password is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Password. Either set the value statically in the configuration, or use the NETSCALERSDX_PASSWORD "+
				"environment variable."),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	paramsapi := service.NitroParamsapi{
		Host:      params.Host.ValueString(),
		Username:  params.Username.ValueString(),
		Password:  params.Password.ValueString(),
		SslVerify: params.SslVerify.ValueBool(),
		LogLevel:  params.LogLevel.ValueString(),
	}

	c, err := service.NewNitroClientFromParams(paramsapi)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NitroClient",
			fmt.Sprintf("The provider cannot create the Nitro API client: %s", err),
		)
		return
	}

	resp.DataSourceData = &c
	resp.ResourceData = &c

}
