package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccMpsuserPlaceholder = `
	
	resource "netscalersdx_mpsuser" "tf_mpsuser" {
		name                    = "tf_mpsuser"
		password                = "VerySecret@1234"
		external_authentication = "false"
		groups                  = ["read_only"]
		session_timeout        = "20"
		session_timeout_unit   = "Minutes"
		enable_session_timeout = "true"
	}				  
	`
)

func TestAccMpsuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMpsuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsuserPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckMpsuserExists("netscalersdx_mpsuser.tf_mpsuser"),
					resource.TestCheckResourceAttr("netscalersdx_mpsuser.tf_mpsuser", "name", "tf_mpsuser"),
					resource.TestCheckResourceAttr("netscalersdx_mpsuser.tf_mpsuser", "session_timeout_unit", "Minutes"),
					resource.TestCheckResourceAttr("netscalersdx_mpsuser.tf_mpsuser", "session_timeout", "20"),
				),
			},
		},
	})
}

func testAccCheckMpsuserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("mpsuser not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mpsuser ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("mpsuser", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mpsuser not found: %s", n)
		}
		return nil
	}
}

func testAccCheckMpsuserDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_mpsuser" {
			continue
		}
		_, err := client.GetResource("mpsuser", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mpsuser still exists")
		}
	}
	return nil
}
