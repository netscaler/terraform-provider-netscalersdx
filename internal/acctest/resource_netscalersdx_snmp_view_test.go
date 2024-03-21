package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpViewPlaceholder = `
	
	resource "netscalersdx_snmp_view" "tf_snmp_view" {
		name    = "tf_snmp_view"
		subtree = "1.3.6.1.2.1.1"
		type    = "false"
	}	  
	`
)

func TestAccSnmpView_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpViewDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpViewPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpViewExists("netscalersdx_snmp_view.tf_snmp_view"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_view.tf_snmp_view", "name", "tf_snmp_view"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_view.tf_snmp_view", "subtree", "1.3.6.1.2.1.1"),
				),
			},
		},
	})
}

func testAccCheckSnmpViewExists(n string,) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Snmp View not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Snmp View ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_view", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Snmp View not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSnmpViewDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_snmp_view" {
			continue
		}
		_, err := client.GetResource("snmp_view", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Snmp View still exists")
		}
	}
	return nil
}
