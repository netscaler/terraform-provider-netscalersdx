package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccDeviceGroupPlaceholder = `
	
	resource "netscalersdx_device_group" "tf_device_group" {
		name                   = "tf_device_group"
		duration               = 10
		category               = "default"
		criteria_value         = "sample"
		static_device_list_arr = ["10.10.10.10"]
	}
	`
)

func TestAccDeviceGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDeviceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceGroupPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceGroupExists("netscalersdx_device_group.tf_device_group", nil),
					resource.TestCheckResourceAttr("netscalersdx_device_group.tf_device_group", "name", "tf_device_group"),
					resource.TestCheckResourceAttr("netscalersdx_device_group.tf_device_group", "duration", "10"),
					resource.TestCheckResourceAttr("netscalersdx_device_group.tf_device_group", "category", "default"),
				),
			},
		},
	})
}

func testAccCheckDeviceGroupExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Device Group not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Device Group ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("device_group", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Group not found: %s", n)
		}
		return nil
	}
}

func testAccCheckDeviceGroupDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_device_group" {
			continue
		}
		_, err := client.GetResource("device_group", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Group still exists")
		}
	}
	return nil
}
