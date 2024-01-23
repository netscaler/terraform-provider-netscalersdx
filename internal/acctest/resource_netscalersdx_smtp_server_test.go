package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSmtpServerPlaceholder = `
	
	resource "netscalersdx_smtp_server" "tf_smtp_server" {
		server_name    = "tf_smtp_server"
		sender_mail_id = "abc.com"
		password       = "secret"
		port           = 587
		username       = "user"
		is_ssl         = "true"
		is_auth        = "true"
	}				  
	`
)

func TestAccSmtpServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmtpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSmtpServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSmtpServerExists("netscalersdx_smtp_server.tf_smtp_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_smtp_server.tf_smtp_server", "server_name", "tf_smtp_server"),
					resource.TestCheckResourceAttr("netscalersdx_smtp_server.tf_smtp_server", "sender_mail_id", "abc.com"),
				),
			},
		},
	})
}

func testAccCheckSmtpServerExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Smtp Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Smtp Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("smtp_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Smtp Server not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSmtpServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_smtp_server" {
			continue
		}
		_, err := client.GetResource("smtp_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Smtp Server still exists")
		}
	}
	return nil
}
