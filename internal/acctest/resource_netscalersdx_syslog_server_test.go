package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	testAccSyslogServerPlaceholder = `
	
	resource "netscalersdx_syslog_server" "tf_syslog_server" {
		name           = "tf_syslog_server"
		ip_address     = "10.10.10.10"
		port           = 514
		log_level_all  = true
		log_level_none = false
	}				  
	`
)

func TestAccSyslogServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSyslogServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogServerPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSyslogServerExists("netscalersdx_syslog_server.tf_syslog_server", nil),
					resource.TestCheckResourceAttr("netscalersdx_syslog_server.tf_syslog_server", "name", "tf_syslog_server"),
					resource.TestCheckResourceAttr("netscalersdx_syslog_server.tf_syslog_server", "ip_address", "10.10.10.10"),
				),
			},
		},
	})
}

func testAccCheckSyslogServerExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Syslog Server not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Syslog Server ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResource("syslog_server", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Syslog Server not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSyslogServerDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_syslog_server" {
			continue
		}
		_, err := client.GetResource("syslog_server", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Syslog Server still exists")
		}
	}
	return nil
}
