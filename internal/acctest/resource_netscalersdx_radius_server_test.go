package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccRadiusServerPlaceholder = `
	
	resource "netscalersdx_radius_server" "tf_radius_server" {
		name         = "tf_radius_server"
		ip_address   = "10.10.10.10"
		radius_key   = "Verysecretkey"
		port         = 389
		auth_timeout = "3"
	}				  
	`
)

func TestAccRadiusServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckRadiusServerExists("netscalersdx_radius_server.tf_radius_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_radius_server.tf_radius_server", "name", "tf_radius_server"),
					resource.TestCheckResourceAttr("netscalersdx_radius_server.tf_radius_server", "ip_address", "10.10.10.10"),
					resource.TestCheckResourceAttr("netscalersdx_radius_server.tf_radius_server", "port", "389"),
				),
			},
		},
	})
}

func testAccCheckRadiusServerExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Radius Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Radius Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("radius_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Radius Server not found: %s", n)
		}
		return nil
	}
}

func testAccCheckRadiusServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_radius_server" {
			continue
		}
		_, err := client.GetResource("radius_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Radius Server still exists")
		}
	}
	return nil
}
