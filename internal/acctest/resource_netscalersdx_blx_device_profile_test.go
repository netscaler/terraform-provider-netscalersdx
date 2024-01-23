package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccBlxDeviceProfilePlaceholder = `
	
	resource "netscalersdx_blx_device_profile" "tf_blx_device_profile" {
		name                                         = "tf_blx_device_profile"
		username                                     = "nsroot"
		password                                     = "Verysecret@123"
		host_username                                = "nsroot"
		host_password                                = "Verysecret@123"
		use_global_setting_for_communication_with_ns = true
		type                                         = "blx"
		http_port                                    = 80
		https_port                                   = 443
		snmpversion                                  = "v3"
		snmpsecuritylevel                            = "NoAuthNoPriv"
		snmpsecurityname                             = "device-profile"
		ssh_port                                     = 22
		svm_ns_comm                                  = "http"
	  }						
	`
)

func TestAccBlxDeviceProfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBlxDeviceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBlxDeviceProfilePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBlxDeviceProfileExists("netscalersdx_blx_device_profile.tf_blx_device_profile", nil),
					resource.TestCheckResourceAttr("netscalersdx_blx_device_profile.tf_blx_device_profile", "name", "tf_blx_device_profile"),
					resource.TestCheckResourceAttr("netscalersdx_blx_device_profile.tf_blx_device_profile", "type", "blx"),
					resource.TestCheckResourceAttr("netscalersdx_blx_device_profile.tf_blx_device_profile", "snmpsecurityname", "device-profile"),
				),
			},
		},
	})
}

func testAccCheckBlxDeviceProfileExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Blx Device Profile not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Blx Device Profile ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("blx_device_profile", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Blx Device Profile not found: %s", n)
		}
		return nil
	}
}

func testAccCheckBlxDeviceProfileDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_blx_device_profile" {
			continue
		}
		_, err := client.GetResource("blx_device_profile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Blx Device Profile still exists")
		}
	}
	return nil
}
