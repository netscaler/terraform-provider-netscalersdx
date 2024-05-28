package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSystemSettingsPlaceholder = `

	resource "netscalersdx_system_settings" "tf_system_settings" {
		secure_access_only               = false
		svm_ns_comm                      = "http"
		enable_cuxip                     = true
		session_timeout_unit             = "Minutes"
		enable_certificate_download      = true
		basicauth                        = true
		session_timeout                  = 15
		enable_session_timeout           = false
		enable_apiproxy_credentials      = false
		enable_shell_access              = true
		keep_alive_ping_interval         = 50
		is_metering_enabled              = true
		disable_agent_old_password_input = false
		enable_nsrecover_login           = true
		enable_delete_interface_on_adc   = false
	  }	  
	`
)

func TestAccSystemSettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSystemSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemSettingsPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSystemSettingsExists("netscalersdx_system_settings.tf_system_settings"),
				),
			},
		},
	})
}

func testAccCheckSystemSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("System Settings not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Settings ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("system_settings", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("System Settings not found: %s", n)
		}
		return nil
	}
}
