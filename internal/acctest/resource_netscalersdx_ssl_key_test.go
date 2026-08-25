package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	// NOTE: requires /tmp/tf_ssl_key_test.key to exist on the test runner before TF_ACC=1 runs.
	// file_location_path is the DIRECTORY; the uploaded file is <dir>/<file_name>.
	testAccSslKeyPlaceholder = `

	resource "netscalersdx_ssl_key" "tf_ssl_key" {
		file_name          = "tf_ssl_key_test.key"
		file_location_path = "/tmp"
	}

	`

	// NOTE: requires /tmp/tf_ssl_key_test.key to exist on the test runner before TF_ACC=1 runs.
	testAccSslKeyDataSource = `

	resource "netscalersdx_ssl_key" "tf_ssl_key" {
		file_name          = "tf_ssl_key_test.key"
		file_location_path = "/tmp"
	}

	data "netscalersdx_ssl_key" "data_tf_ssl_key" {
		id = netscalersdx_ssl_key.tf_ssl_key.id
	}

	`
)

func TestAccSslKey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslKeyPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSslKeyExists("netscalersdx_ssl_key.tf_ssl_key"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_key.tf_ssl_key", "file_name", "tf_ssl_key_test.key"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_key.tf_ssl_key", "file_location_path", "/tmp"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_key.tf_ssl_key", "id", "tf_ssl_key_test.key"),
				),
			},
		},
	})
}

func TestAccSslKey_dataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslKeyDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Pair id and file_name (server-returned). Deliberately skip
					// file_location_path — sslKeySetAttrFromGet does not copy it
					// from the server response (server path != local path), so
					// the data source value would not match the resource value.
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_key.data_tf_ssl_key", "id",
						"netscalersdx_ssl_key.tf_ssl_key", "id",
					),
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_key.data_tf_ssl_key", "file_name",
						"netscalersdx_ssl_key.tf_ssl_key", "file_name",
					),
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_key.data_tf_ssl_key", "file_size",
						"netscalersdx_ssl_key.tf_ssl_key", "file_size",
					),
				),
			},
		},
	})
}

func testAccCheckSslKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("SslKey not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SslKey ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResourceV2("ssl_key", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SslKey not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSslKeyDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_ssl_key" {
			continue
		}
		_, err := client.GetResourceV2("ssl_key", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SslKey still exists")
		}
	}
	return nil
}
