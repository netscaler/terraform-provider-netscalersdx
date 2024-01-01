package acctest

import (
	"os"
	"terraform-provider-netscalersdx/internal/provider"
	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"netscalersdx": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func testAccApiClient() (*service.NitroClient, error) {
	paramsapi := testAccApiClientParams()

	c, err := service.NewNitroClientFromParams(paramsapi)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func testAccApiClientParams() service.NitroParamsapi {
	paramsapi := service.NitroParamsapi{}

	if k := os.Getenv("NETSCALERSDX_HOST"); k != "" {
		paramsapi.Host = k
	}
	if k := os.Getenv("NETSCALERSDX_USERNAME"); k != "" {
		paramsapi.Username = k
	}
	if k := os.Getenv("NETSCALERSDX_PASSWORD"); k != "" {
		paramsapi.Password = k
	}

	return paramsapi
}
