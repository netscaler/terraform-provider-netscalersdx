package citrixsdx

import (
	"fmt"
	"log"
	"os"
	"terraform-provider-citrixsdx/service"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testAccProvisionVpxPlaceholder = `
	resource "citrixsdx_provision_vpx" "device1" {
		name                       = "accTest-device1"
		ip_address                 = "%s"
		if_internal_ip_enabled     = false
		config_type                = 0
		ipv4_address               = "%s"
		netmask                    = "%s"
		gateway                    = "%s"
		nexthop                    = ""
		image_name                 = "%s"
		profile_name               = "%s"
		description                = "from tf"
		throughput_allocation_mode = "0"
		throughput                 = "1000"
		max_burst_throughput       = "0"
		burst_priority             = "0"
		license                    = "Standard"
		number_of_acu              = 0
		number_of_scu              = "0"
		vm_memory_total            = "2048"
		pps                        = "1000000"
		number_of_cores            = "0"
		l2_enabled                 = "false"
		if_0_1                     = true
		vlan_id_0_1                = ""
		if_0_2                     = true
		vlan_id_0_2                = ""
		network_interfaces {
		  port_name       = "LA/1"
		  mac_address     = ""
		  mac_mode        = "default"
		  receiveuntagged = "true"
		}
		nsvlan_id         = ""
		vlan_type         = 1
		nsvlan_tagged     = "false"
		nsvlan_interfaces = []
	  }
	`
)

var testAccProvisionVpxAdd = fmt.Sprintf(testAccProvisionVpxPlaceholder,
	os.Getenv("VPX_IP"),
	os.Getenv("VPX_IP"),
	os.Getenv("VPX_NETMASK"),
	os.Getenv("VPX_GATEWAY"),
	os.Getenv("VPX_IMAGE"),
	os.Getenv("VPX_PROFILE"),
)

// example.Widget represents a concrete Go type that represents an API resource
func TestAccProvisionVpx_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: false,
		PreCheck:   func() { testAccPreCheck(t) },
		// ProviderFactories: providerFactories,
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProvisionVpxDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProvisionVpxAdd,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvisionVpxExists("citrixsdx_provision_vpx.device1", nil),
				),
			},
		},
	})
}

func testAccCheckProvisionVpxExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		log.Println("[DEBUG] testAccCheckProvisionVpxExists")
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vpx Device ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed")
			}

			*id = rs.Primary.ID
		}

		// retrieve the client from the test provider
		c := testAccProvider.Meta().(*service.NitroClient)
		data, err := c.GetResource("ns", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Resource %s not found", n)
		}

		return nil
	}
}

func testAccCheckProvisionVpxDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*service.NitroClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixsdx_provision_vpx" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := c.GetResource("ns", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Provisioned Vpx %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
