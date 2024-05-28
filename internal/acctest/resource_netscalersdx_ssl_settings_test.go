package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSslSettingsPlaceholder = `

	resource "netscalersdx_ssl_settings" "tf_ssl_settings" {
		sslreneg = false
		tlsv1_1  = true
		sslv3    = false
		tlsv1_2  = true
		tlsv1    = false
	  }
				  
	`
)

func TestAccSslSettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSslSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslSettingsPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSslSettingsExists("netscalersdx_ssl_settings.tf_ssl_settings"),
				),
			},
		},
	})
}

func testAccCheckSslSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Ssl Settings not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ssl Settings ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetAllResource("ssl_settings")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Ssl Settings not found: %s", n)
		}
		return nil
	}
}
