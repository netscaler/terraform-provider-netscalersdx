package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccMpsFeaturePlaceholder = `
	
		resource "netscalersdx_mps_feature" "tf_mps_feature" {
			admin_toggle = 3
			built_in     = "false"
			feature_name = "Device_Syslog"
			ops_toggle   = 0
		}	  
	`
)

func TestAccMpsFeature_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckMpsFeatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsFeaturePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckMpsFeatureExists("netscalersdx_mps_feature.tf_mps_feature"),
				),
			},
		},
	})
}

func testAccCheckMpsFeatureExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Mps Feature not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Mps Feature ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("mps_feature", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Mps Feature not found: %s", n)
		}
		return nil
	}
}
