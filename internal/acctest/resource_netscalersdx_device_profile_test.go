package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccDeviceProfilePlaceholder = `
	
	resource "netscalersdx_device_profile" "tf_device_profile" {
		name                                         = "tf_device_profile"
		username                                     = "user"
		password                                     = "Verysecret@123"
		host_username                                = "root"
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

func TestAccDeviceProfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDeviceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceProfilePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceProfileExists("netscalersdx_device_profile.tf_device_profile", nil),
					resource.TestCheckResourceAttr("netscalersdx_device_profile.tf_device_profile", "name", "tf_device_profile"),
					resource.TestCheckResourceAttr("netscalersdx_device_profile.tf_device_profile", "type", "blx"),
					resource.TestCheckResourceAttr("netscalersdx_device_profile.tf_device_profile", "snmpsecurityname", "device-profile"),
				),
			},
		},
	})
}

func testAccCheckDeviceProfileExists(n string, id *string) resource.TestCheckFunc {
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

		data, err := client.GetResource("device_profile", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Profile not found: %s", n)
		}
		return nil
	}
}

func testAccCheckDeviceProfileDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_device_profile" {
			continue
		}
		_, err := client.GetResource("device_profile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Profile still exists")
		}
	}
	return nil
}
