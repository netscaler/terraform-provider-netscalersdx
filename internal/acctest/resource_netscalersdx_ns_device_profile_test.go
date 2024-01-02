package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccNsDeviceProfilePlaceholder = `
	
	resource "netscalersdx_ns_device_profile" "tf_ns_device_profile" {
		name                                         = "tf_ns_device_profile"
		password                                     = "Verysecret@123"
		username                                     = "nsroot"
		use_global_setting_for_communication_with_ns = true
		type                                         = "ns"
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

func TestAccNsDeviceProfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsDeviceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsDeviceProfilePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNsDeviceProfileExists("netscalersdx_ns_device_profile.tf_ns_device_profile", nil),
					resource.TestCheckResourceAttr("netscalersdx_ns_device_profile.tf_ns_device_profile", "name", "tf_ns_device_profile"),
					resource.TestCheckResourceAttr("netscalersdx_ns_device_profile.tf_ns_device_profile", "type", "ns"),
					resource.TestCheckResourceAttr("netscalersdx_ns_device_profile.tf_ns_device_profile", "snmpsecurityname", "device-profile"),
				),
			},
		},
	})
}

func testAccCheckNsDeviceProfileExists(n string, id *string) resource.TestCheckFunc {
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

		data, err := client.GetResource("ns_device_profile", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Profile not found: %s", n)
		}
		return nil
	}
}

func testAccCheckNsDeviceProfileDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_ns_device_profile" {
			continue
		}
		_, err := client.GetResource("ns_device_profile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Profile still exists")
		}
	}
	return nil
}
