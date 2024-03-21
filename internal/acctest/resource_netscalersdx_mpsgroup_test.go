package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccMpsgroupPlaceholder = `
	
	resource "netscalersdx_mpsgroup" "tf_mpsgroup" {
		name                     = "tf_mpsgroup"
		permission               = "readonly"
		select_individual_entity = "false"
		assign_all_apps          = "true"
		description              = "test-terraform"
		role                     = "nonadmin"
	}				  
	`
)

func TestAccMpsgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMpsgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsgroupPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckMpsgroupExists("netscalersdx_mpsgroup.tf_mpsgroup"),
					resource.TestCheckResourceAttr("netscalersdx_mpsgroup.tf_mpsgroup", "name", "tf_mpsgroup"),
					resource.TestCheckResourceAttr("netscalersdx_mpsgroup.tf_mpsgroup", "permission", "readonly"),
					resource.TestCheckResourceAttr("netscalersdx_mpsgroup.tf_mpsgroup", "description", "test-terraform"),
				),
			},
		},
	})
}

func testAccCheckMpsgroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("mpsgroup not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mpsgroup ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("mpsgroup", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mpsgroup not found: %s", n)
		}
		return nil
	}
}

func testAccCheckMpsgroupDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_mpsgroup" {
			continue
		}
		_, err := client.GetResource("mpsgroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mpsgroup still exists")
		}
	}
	return nil
}
