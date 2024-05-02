package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpMibPlaceholder = `
	
	resource "netscalersdx_snmp_mib" "tf_snmp_mib" {
		name     = "tf-testing"
		contact  = "NetScaler"
		location = "NetScaler123"
	  }
						  
	`
)

func TestAccSnmpMib_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSnmpMibDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpMibPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpMibExists("netscalersdx_snmp_mib.tf_snmp_mib"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_mib.tf_snmp_mib", "name", "tf-testing"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_mib.tf_snmp_mib", "contact", "NetScaler"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_mib.tf_snmp_mib", "location", "NetScaler123"),
				),
			},
		},
	})
}

func testAccCheckSnmpMibExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Snmp Mib not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Snmp Mib ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_mib", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Snmp Mib not found: %s", n)
		}
		return nil
	}
}
