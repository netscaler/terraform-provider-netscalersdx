package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccAclrulePlaceholder = `
	
	resource "netscalersdx_aclrule" "tf_aclrule" {
		name     = "tf_aclrule"
		priority = 100
		protocol = "TCP"
		action   = "Allow"
		dst_port = 80
		src_ip   = "10.10.10.10"
	  }
					  
	`
)

func TestAccAclrule_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAclruleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAclrulePlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAclruleExists("netscalersdx_aclrule.tf_aclrule"),
					resource.TestCheckResourceAttr("netscalersdx_aclrule.tf_aclrule", "name", "tf_aclrule"),
					resource.TestCheckResourceAttr("netscalersdx_aclrule.tf_aclrule", "protocol", "TCP"),
					resource.TestCheckResourceAttr("netscalersdx_aclrule.tf_aclrule", "action", "Allow"),
				),
			},
		},
	})
}

func testAccCheckAclruleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Aclrule not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Aclrule ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("aclrule", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Aclrule not found: %s", n)
		}
		return nil
	}
}

func testAccCheckAclruleDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_aclrule" {
			continue
		}
		_, err := client.GetResource("aclrule", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Aclrule still exists")
		}
	}
	return nil
}
