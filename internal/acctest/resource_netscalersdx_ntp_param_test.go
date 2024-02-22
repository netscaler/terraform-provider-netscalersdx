package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccNtpParamPlaceholder = `
	
	resource "netscalersdx_ntp_param" "tf_ntp_param" {
		automax_logsec = 12
		revoke_logsec  = 16
		authentication = true
	}				  
	`
)

func TestAccNtpParam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckNtpParamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpParamPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNtpParamExists("netscalersdx_ntp_param.tf_ntp_param", nil),
					resource.TestCheckResourceAttr("netscalersdx_ntp_param.tf_ntp_param", "automax_logsec", "12"),
					resource.TestCheckResourceAttr("netscalersdx_ntp_param.tf_ntp_param", "revoke_logsec", "16"),
				),
			},
		},
	})
}

func testAccCheckNtpParamExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Ntp Param not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ntp Param ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("ntp_param", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Ntp Param not found: %s", n)
		}
		return nil
	}
}
