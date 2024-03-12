package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpAlarmConfigPlaceholder = `
	
	resource "netscalersdx_snmp_alarm_config" "tf_snmp_alarm_config" {
		name      = "cpuTempError"
		severity  = "Critical"
		threshold = "60"
		enable    = "true"
	}		  
	`
)

func TestAccSnmpAlarmConfig_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSnmpAlarmConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpAlarmConfigPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpAlarmConfigExists("netscalersdx_snmp_alarm_config.tf_snmp_alarm_config"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_alarm_config.tf_snmp_alarm_config", "name", "cpuTempError"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_alarm_config.tf_snmp_alarm_config", "severity", "Critical"),
				),
			},
		},
	})
}

func testAccCheckSnmpAlarmConfigExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Snmp Alarm Config not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Snmp Alarm Config ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_alarm_config", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Snmp Alarm Config not found: %s", n)
		}
		return nil
	}
}
