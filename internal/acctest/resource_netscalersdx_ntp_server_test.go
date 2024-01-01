package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccNtpServerPlaceholder = `
	
	resource "netscalersdx_ntp_server" "tf_ntp_server" {
		server  = "10.10.10.11"
		key_id  = 123
		minpoll = 5
		maxpoll = 12
		autokey = false
	}				  
	`
)

func TestAccNtpServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNtpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckNtpServerExists("netscalersdx_ntp_server.tf_ntp_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_ntp_server.tf_ntp_server", "server", "10.10.10.11"),
				),
			},
		},
	})
}

func testAccCheckNtpServerExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Ntp Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ntp Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("ntp_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Ntp Server not found: %s", n)
		}
		return nil
	}
}

func testAccCheckNtpServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_ntp_server" {
			continue
		}
		_, err := client.GetResource("ntp_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Ntp Server still exists")
		}
	}
	return nil
}
