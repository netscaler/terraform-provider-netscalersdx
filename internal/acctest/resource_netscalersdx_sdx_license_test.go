package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSdxLicensePlaceholder = `

	resource "netscalersdx_sdx_license" "tf_sdx_license" {
		file_name = "temp.lic"
	}
		  
	`
)

func TestAccSdxLicense_basic(t *testing.T) {
	// TODO: Need to find a way to test this resource
	t.Skip("TODO: Need to find a way to test this resource")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckSdxLicenseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSdxLicensePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSdxLicenseExists("netscalersdx_sdx_license.tf_sdx_license"),
				),
			},
		},
	})
}

func testAccCheckSdxLicenseExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Sdx License not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Sdx License ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetAllResource("sdx_license")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Sdx License not found: %s", n)
		}
		return nil
	}
}
