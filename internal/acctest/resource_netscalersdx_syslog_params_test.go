package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSyslogParamsPlaceholder = `
	
	resource "netscalersdx_syslog_params" "tf_syslog_params" {
		date_format = "MMDDYYYY"
		timezone    = "GMT"
	}				  
	`
)

func TestAccSyslogParams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSyslogParamsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogParamsPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSyslogParamsExists("netscalersdx_syslog_params.tf_syslog_params", nil),
					resource.TestCheckResourceAttr("netscalersdx_syslog_params.tf_syslog_params", "date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("netscalersdx_syslog_params.tf_syslog_params", "timezone", "GMT"),
				),
			},
		},
	})
}

func testAccCheckSyslogParamsExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Syslog Params not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Syslog Params ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("syslog_params", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Syslog Params not found: %s", n)
		}
		return nil
	}
}
