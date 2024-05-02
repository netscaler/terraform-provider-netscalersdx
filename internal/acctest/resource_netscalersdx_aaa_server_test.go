package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccAaaServerPlaceholder = `

	resource "netscalersdx_aaa_server" "tf_aaa_server" {
		primary_server_type           = "LOCAL"
	}			
	`
)

func TestAccAaaServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckAaaServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAaaServerExists("netscalersdx_aaa_server.tf_aaa_server"),
				),
			},
		},
	})
}

func testAccCheckAaaServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Aaa Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Aaa Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("aaa_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Aaa Server not found: %s", n)
		}
		return nil
	}
}
