package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccLdapServerPlaceholder = `
	
	resource "netscalersdx_ldap_server" "tf_ldap_server" {
		name                       = "tf_ldap_server"
		ip_address                 = "10.10.10.10"
		sec_type                   = "PLAINTEXT"
		type                       = "AD"
		port                       = 389
		auth_timeout               = "3"
		validate_ldap_server_certs = false
	}				  
	`
)

func TestAccLdapServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLdapServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLdapServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckLdapServerExists("netscalersdx_ldap_server.tf_ldap_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_ldap_server.tf_ldap_server", "name", "tf_ldap_server"),
					resource.TestCheckResourceAttr("netscalersdx_ldap_server.tf_ldap_server", "ip_address", "10.10.10.10"),
					resource.TestCheckResourceAttr("netscalersdx_ldap_server.tf_ldap_server", "sec_type", "PLAINTEXT"),
				),
			},
		},
	})
}

func testAccCheckLdapServerExists(n string, id *string) resource.TestCheckFunc {
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

		data, err := client.GetResource("ldap_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Device Profile not found: %s", n)
		}
		return nil
	}
}

func testAccCheckLdapServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_ldap_server" {
			continue
		}
		_, err := client.GetResource("ldap_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Device Profile still exists")
		}
	}
	return nil
}
