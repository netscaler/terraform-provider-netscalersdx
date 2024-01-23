package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccCipherGroupPlaceholder = `
	
	resource "netscalersdx_cipher_group" "tf_cipher_group" {
		cipher_group_description = "from terraform"
		cipher_group_name        = "tf_cipher_group"
		cipher_name_list_array   = ["TLS1-AES-256-CBC-SHA", "TLS1-ECDHE-RSA-DES-CBC3-SHA"]
	}				  
	`
)

func TestAccCipherGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCipherGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCipherGroupPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCipherGroupExists("netscalersdx_cipher_group.tf_cipher_group", nil),
					resource.TestCheckResourceAttr("netscalersdx_cipher_group.tf_cipher_group", "cipher_group_name", "tf_cipher_group"),
					resource.TestCheckResourceAttr("netscalersdx_cipher_group.tf_cipher_group", "cipher_group_description", "from terraform"),
				),
			},
		},
	})
}

func testAccCheckCipherGroupExists(n string, id *string) resource.TestCheckFunc {
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

		endpoint := "cipher_group"
		returnArr, err := client.GetAllResource(endpoint)

		if err != nil {
			return err
		}

		found := false

		returnData := returnArr[endpoint].([]interface{})
		for _, v := range returnData {
			m := v.(map[string]interface{})
			if m["cipher_group_name"] == rs.Primary.ID {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("cipher_group %s not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckCipherGroupDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_cipher_group" {
			continue
		}

		endpoint := "cipher_group"
		returnArr, err := client.GetAllResource(endpoint)

		if err != nil {
			return fmt.Errorf("Cipher Group still exists")
		}

		found := false

		returnData := returnArr[endpoint].([]interface{})
		for _, v := range returnData {
			m := v.(map[string]interface{})
			if m["cipher_group_name"] == rs.Primary.ID {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("Cipher Group %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
