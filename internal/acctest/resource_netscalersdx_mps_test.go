package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccMpsPlaceholder = `
	
	resource "netscalersdx_mps" "tf_mps" {
		is_cloud                   = "false"
		is_passive                 = "false"
		is_member_of_default_group = "true"
		is_thirdparty_vm_supported = "false"
		is_container               = "false"
		hist_mig_inprog            = "false"
		config_motd                = "true"
		motd                       = "Welcome to NetScaler SDX"
	  }
	`
)

func TestAccMps_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckMpsExists("netscalersdx_mps.tf_mps"),
				),
			},
		},
	})
}

func testAccCheckMpsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Mps not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Mps ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("mps", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Mps not found: %s", n)
		}
		return nil
	}
}
