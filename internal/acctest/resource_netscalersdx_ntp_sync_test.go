package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccNtpSyncPlaceholder = `
	
	resource "netscalersdx_ntp_sync" "tf_ntp_sync" {
		ntpd_status = false
	}					
	`
)

func TestAccNtpSync_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpSyncPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNtpSyncExists("netscalersdx_ntp_sync.tf_ntp_sync"),
				),
			},
		},
	})
}

func testAccCheckNtpSyncExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Ntp Sync not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ntp Sync ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetAllResource("ntp_sync")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Ntp Sync not found: %s", n)
		}
		return nil
	}
}
