package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpTrapPlaceholder = `
	
	resource "netscalersdx_snmp_trap" "tf_snmp_trap" {
		community   = "public"
		version     = "v2"
		dest_server = "10.10.10.10"
		dest_port   = 163
	}				  
	`
)

func TestAccSnmpTrap_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpTrapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpTrapPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpTrapExists("netscalersdx_snmp_trap.tf_snmp_trap", nil),
					resource.TestCheckResourceAttr("netscalersdx_snmp_trap.tf_snmp_trap", "version", "v2"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_trap.tf_snmp_trap", "dest_port", "163"),
				),
			},
		},
	})
}

func testAccCheckSnmpTrapExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Snmp Trap not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Snmp Trap ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_trap", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Snmp Trap not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSnmpTrapDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_snmp_trap" {
			continue
		}
		_, err := client.GetResource("snmp_trap", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Snmp Trap still exists")
		}
	}
	return nil
}
