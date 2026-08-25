package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	// NOTE: requires /tmp/tf_ssl_cert_test.pem to exist on the test runner before TF_ACC=1 runs.
	// file_location_path is the DIRECTORY; the uploaded file is <dir>/<file_name>.
	testAccSslCertPlaceholder = `

	resource "netscalersdx_ssl_cert" "tf_ssl_cert" {
		file_name          = "tf_ssl_cert_test.pem"
		file_location_path = "/tmp"
	}

	`

	// NOTE: requires /tmp/tf_ssl_cert_test.pem to exist on the test runner before TF_ACC=1 runs.
	testAccSslCertDataSource = `

	resource "netscalersdx_ssl_cert" "tf_ssl_cert" {
		file_name          = "tf_ssl_cert_test.pem"
		file_location_path = "/tmp"
	}

	data "netscalersdx_ssl_cert" "data_tf_ssl_cert" {
		id = netscalersdx_ssl_cert.tf_ssl_cert.id
	}

	`
)

func TestAccSslCert_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslCertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslCertPlaceholder,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSslCertExists("netscalersdx_ssl_cert.tf_ssl_cert"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_cert.tf_ssl_cert", "file_name", "tf_ssl_cert_test.pem"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_cert.tf_ssl_cert", "file_location_path", "/tmp"),
					resource.TestCheckResourceAttr("netscalersdx_ssl_cert.tf_ssl_cert", "id", "tf_ssl_cert_test.pem"),
				),
			},
		},
	})
}

func TestAccSslCert_dataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslCertDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Pair id and file_name (server-returned). Deliberately skip
					// file_location_path — sslCertSetAttrFromGet does not copy it
					// from the server response (server path != local path), so
					// the data source value would not match the resource value.
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_cert.data_tf_ssl_cert", "id",
						"netscalersdx_ssl_cert.tf_ssl_cert", "id",
					),
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_cert.data_tf_ssl_cert", "file_name",
						"netscalersdx_ssl_cert.tf_ssl_cert", "file_name",
					),
					resource.TestCheckResourceAttrPair(
						"data.netscalersdx_ssl_cert.data_tf_ssl_cert", "file_size",
						"netscalersdx_ssl_cert.tf_ssl_cert", "file_size",
					),
				),
			},
		},
	})
}

func testAccCheckSslCertExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("SslCert not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SslCert ID is set")
		}

		client, err := testAccApiClient()
		if err != nil {
			return err
		}

		data, err := client.GetResourceV2("ssl_cert", rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SslCert not found: %s", n)
		}
		return nil
	}
}

func testAccCheckSslCertDestroy(s *terraform.State) error {
	client, err := testAccApiClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscalersdx_ssl_cert" {
			continue
		}
		_, err := client.GetResourceV2("ssl_cert", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SslCert still exists")
		}
	}
	return nil
}
