package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccCurrentHostnamePlaceholder = `
		
	resource "netscalersdx_current_hostname" "tf_current_hostname" {
		hostname            = "example-mgmt"
		hypervisor_hostname = "netscaler-sdx"
	}
						
	`
)

func TestAccCurrentHostname_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCurrentHostnameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCurrentHostnamePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCurrentHostnameExists("netscalersdx_current_hostname.tf_current_hostname"),
				),
			},
		},
	})
}

func testAccCheckCurrentHostnameExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Current hostname not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Current hostname ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("current_hostname", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Current hostname not found: %s", n)
		}
		return nil
	}
}
