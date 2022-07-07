package citrixsdx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "fmt"
	// "strings"
	"terraform-provider-citrixsdx/service"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CITRIXSDX_HOST", nil),
				Description: "Citrix SDX host. Can be specified with `CITRIXSDX_HOST` environment variable. This has to start with https://",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					// if the value does not start with http or https, throw an error
					if !(strings.HasPrefix(v.(string), "https://") || strings.HasPrefix(v.(string), "http://")) {
						errors = append(errors, fmt.Errorf("host must start with https://"))
					}
					return
				},
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CITRIXSDX_USERNAME", nil),
				Description: "Citrix SDX username. Can be specified with `CITRIXSDX_USERNAME` environment variable.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CITRIXSDX_PASSWORD", nil),
				Description: "Citrix SDX password. Can be specified with `CITRIXSDX_PASSWORD` environment variable.",
			},
			"ssl_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: envDefaultFunc("CITRIXSDX_SSL_VERIFY", true),
				Description: "Ignore validity of SDX TLS certificate if true. Can be specified with `CITRIXSDX_SSL_VERIFY` environment variable.",
			},
			"log_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "INFO",
				DefaultFunc: schema.EnvDefaultFunc("CITRIXSDX_LOG_LEVEL", nil),
				Description: "Log level (Default is INFO). Can be specified with `CITRIXSDX_LOG_LEVEL` environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"citrixsdx_provision_vpx": resourceProvisionVpx(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func envDefaultFunc(k string, dv interface{}) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		v := os.Getenv(k)
		if v == "" {
			return dv, nil
		}
		if v == "true" || v == "false" {
			return strconv.ParseBool(v)
		}
		return v, nil
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	tflog.Trace(ctx, "In providerConfigure")
	var diags diag.Diagnostics
	params := service.NitroParams{
		Host:      d.Get("host").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		SslVerify: d.Get("ssl_verify").(bool),
		LogLevel:  d.Get("log_level").(string),
	}
	c, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
