package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSnmpUserPlaceholder = `
	
	resource "netscalersdx_snmp_user" "tf_snmp_user" {
		name             = "tf_snmp_user"
		security_level   = 2
		auth_protocol    = 1
		auth_password    = "Verysecret@123"
		privacy_protocol = 1
		privacy_password = "Verysecret@123"
	}				  
	`
)

func TestAccSnmpUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpUserPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnmpUserExists("netscalersdx_snmp_user.tf_snmp_user", nil),
					resource.TestCheckResourceAttr("netscalersdx_snmp_user.tf_snmp_user", "name", "tf_snmp_user"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_user.tf_snmp_user", "security_level", "2"),
					resource.TestCheckResourceAttr("netscalersdx_snmp_user.tf_snmp_user", "auth_protocol", "1"),
				),
			},
		},
	})
}

func testAccCheckSnmpUserExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Snmp User not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Snmp User ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("snmp_user", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Snmp User not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSnmpUserDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_snmp_user" {
			continue
		}
		_, err := client.GetResource("snmp_user", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Snmp User still exists")
		}
	}
	return nil
}
