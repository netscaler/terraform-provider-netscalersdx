package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccTacacsServerPlaceholder = `
	
	resource "netscalersdx_tacacs_server" "tf_tacacs_server" {
		name       = "tf_tacacs_server"
		port       = 545
		tacacs_key = "key"
		ip_address = "10.10.10.11"
	}				  
	`
)

func TestAccTacacsServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTacacsServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTacacsServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTacacsServerExists("netscalersdx_tacacs_server.tf_tacacs_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_tacacs_server.tf_tacacs_server", "name", "tf_tacacs_server"),
					resource.TestCheckResourceAttr("netscalersdx_tacacs_server.tf_tacacs_server", "port", "545"),
					resource.TestCheckResourceAttr("netscalersdx_tacacs_server.tf_tacacs_server", "ip_address", "10.10.10.11"),
				),
			},
		},
	})
}

func testAccCheckTacacsServerExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Tacacs Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Tacacs Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("tacacs_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Tacacs Server not found: %s", n)
		}
		return nil
	}
}

func testAccCheckTacacsServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_tacacs_server" {
			continue
		}
		_, err := client.GetResource("tacacs_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Tacacs Server still exists")
		}
	}
	return nil
}
