package acctest

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccStaticRoutePlaceholder = `
	
	resource "netscalersdx_static_route" "tf_static_route" {
		network = "%s"
		gateway = "%s"
		netmask = "255.255.255.0"
	}				  
	`
)

var testAccStaticRouteConfig = fmt.Sprintf(testAccStaticRoutePlaceholder,
	os.Getenv("NETWORK"),
	os.Getenv("GATEWAY"),
)

func TestAccStaticRoute_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStaticRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStaticRouteConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckStaticRouteExists("netscalersdx_static_route.tf_static_route", nil),
				),
			},
		},
	})
}

func testAccCheckStaticRouteExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Static Rroute not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Static Rroute ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("static_route", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Static Rroute not found: %s", n)
		}
		return nil
	}
}

func testAccCheckStaticRouteDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_static_route" {
			continue
		}
		_, err := client.GetResource("static_route", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Static Rroute still exists")
		}
	}
	return nil
}
