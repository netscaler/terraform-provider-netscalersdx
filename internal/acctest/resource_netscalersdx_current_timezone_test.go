package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccCurrentTimezonePlaceholder = `
	
	resource "netscalersdx_current_timezone" "tf_current_timezone" {
		timezone = "UTC+0000 GMT Europe/London"
	}
						
	`
)

func TestAccCurrentTimezone_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCurrentTimezoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCurrentTimezonePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCurrentTimezoneExists("netscalersdx_current_timezone.tf_current_timezone"),
				),
			},
		},
	})
}

func testAccCheckCurrentTimezoneExists(n string) resource.TestCheckFunc {
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

		data, err := client.GetResource("current_timezone", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Profile not found: %s", n)
		}
		return nil
	}
}
