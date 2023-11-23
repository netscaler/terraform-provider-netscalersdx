package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	
	"terraform-provider-citrixsdx/internal/service"

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
	resp.TypeName = "citrixsdx"
	resp.Version = p.version
}

func (p *sdxprovider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:    true,
				Description: "Citrix SDX host. Can be specified with `CITRIXSDX_HOST` environment variable. This has to start with https://",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "Citrix SDX username. Can be specified with `CITRIXSDX_USERNAME` environment variable.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Citrix SDX password. Can be specified with `CITRIXSDX_PASSWORD` environment variable.",
			},
			"ssl_verify": schema.BoolAttribute{
				Optional:    true,
				Description: "Ignore validity of SDX TLS certificate if true. Can be specified with `CITRIXSDX_SSL_VERIFY` environment variable.",
			},
			"log_level": schema.StringAttribute{
				Optional:    true,
				Description: "Log level (Default is INFO). Can be specified with `CITRIXSDX_LOG_LEVEL` environment variable.",
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
		resourceProvisionVpx,
		resourceVpxState,
	}
}
func (p *sdxprovider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		dataSourceVpx,
	}
}

// Configure prepares an API client for data sources and resources.
func (p *sdxprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var params service.NitroParams
	log.Printf("req.Config: %s", req.Config)

	diags := req.Config.Get(ctx, &params)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Use values in environment variables if the values aren't set in the configuration.
	hostEnv := os.Getenv("CITRIXSDX_HOST")
	if params.Host.IsNull() && hostEnv != "" {
		params.Host = types.StringValue(hostEnv)
	}
	usernameEnv := os.Getenv("CITRIXSDX_USERNAME")
	if params.Username.IsNull() && usernameEnv != "" {
		params.Username = types.StringValue(usernameEnv)
	}
	passwordEnv := os.Getenv("CITRIXSDX_PASSWORD")
	if params.Password.IsNull() && passwordEnv != "" {
		params.Password = types.StringValue(passwordEnv)
	}
	if params.SslVerify.IsNull() {
		envSslVerify, _ := strconv.ParseBool(os.Getenv("CITRIXSDX_SSL_VERIFY"))
		params.SslVerify = types.BoolValue(envSslVerify)
	}
	if params.LogLevel.IsNull() {
		params.LogLevel = types.StringValue(os.Getenv("CITRIXSDX_LOG_LEVEL"))
	}

	if params.Host.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"host is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Host. Either set the value statically in the configuration, or use the CITRIXSDX_HOST "+
				"environment variable."),
		)
	}
	if params.Username.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"username is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Username. Either set the value statically in the configuration, or use the CITRIXSDX_USERNAME "+
				"environment variable."),
		)
	}
	if params.Password.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"password is required",
			fmt.Sprintf("The provider cannot create the Nitro API client as there is an unknown configuration "+
				"value for the Password. Either set the value statically in the configuration, or use the CITRIXSDX_PASSWORD "+
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
