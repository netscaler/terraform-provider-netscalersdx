package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccMpsSslCertkeyPlaceholder = `
	
	resource "netscalersdx_mps_ssl_certkey" "tf_mps_ssl_certkey" {
		ssl_certificate = "duplicate_cert"
		ssl_key         = "duplicate_key"
	}			
	`
)

func TestAccMpsSslCertkey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsSslCertkeyPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckMpsSslCertkeyExists("netscalersdx_mps_ssl_certkey.tf_mps_ssl_certkey"),
				),
			},
		},
	})
}

func testAccCheckMpsSslCertkeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Mps Ssl Certkey not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Mps Ssl Certkey ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetAllResource("mps_ssl_certkey")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Mps Ssl Certkey not found: %s", n)
		}
		return nil
	}
}
