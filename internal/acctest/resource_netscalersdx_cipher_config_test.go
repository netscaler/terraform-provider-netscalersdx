package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccCipherConfigPlaceholder = `

	resource "netscalersdx_cipher_config" "tf_cipher_config" {
		config_mode = "CipherSuites"
		cipher_name_list_array = [
		  "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384",
		  "TLS1.2-ECDHE-ECDSA-AES256-GCM-SHA384",
		  "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256",
		  "TLS1.2-ECDHE-ECDSA-AES128-GCM-SHA256",
		  "TLS1.2-ECDHE-RSA-AES-256-SHA384",
		  "TLS1.2-ECDHE-ECDSA-AES256-SHA384",
		  "TLS1.2-ECDHE-RSA-AES-128-SHA256",
		  "TLS1.2-ECDHE-ECDSA-AES128-SHA256",
		  "TLS1.2-DHE-DSS-AES256-GCM-SHA384",
		  "TLS1.2-DHE-RSA-AES256-GCM-SHA384",
		  "TLS1.2-DHE-DSS-AES128-GCM-SHA256",
		  "TLS1.2-DHE-RSA-AES128-GCM-SHA256",
		  "TLS1.2-DHE-RSA-AES-256-SHA256",
		  "TLS1.2-DHE-DSS-AES256-SHA256",
		  "TLS1.2-DHE-RSA-AES-128-SHA256",
		  "SSL2-DES-CBC-MD5"
		]
	  }			  
	`
)

func TestAccCipherConfig_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// CheckDestroy:             testAccCheckCipherConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCipherConfigPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCipherConfigExists("netscalersdx_cipher_config.tf_cipher_config"),
				),
			},
		},
	})
}

func testAccCheckCipherConfigExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Cipher Config not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cipher Config ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetAllResource("cipher_config")
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Cipher Config not found: %s", n)
		}
		return nil
	}
}
