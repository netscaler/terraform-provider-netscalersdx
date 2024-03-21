package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpManagerPlaceholder = `

	resource "netscalersdx_snmp_manager" "tf_snmp_manager" {
		ip_address = "10.10.10.10"
		community  = "public"
		netmask    = "255.255.255.0"
	  }
					
	`
)

func TestAccSnmpManager_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpManagerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpManagerExists("netscalersdx_snmp_manager.tf_snmp_manager"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_manager.tf_snmp_manager", "ip_address", "10.10.10.10"),
				),
			},
		},
	})
}

func testAccCheckSnmpManagerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Device Profile not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Device Profile ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_manager", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Profile not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSnmpManagerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_snmp_manager" {
			continue
		}
		_, err := client.GetResource("snmp_manager", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Profile still exists")
		}
	}
	return nil
}
