resource "netscalersdx_ssl_cert" "tf_ssl_cert" {
  file_name = "tf_ssl_cert.pem"
  # file_location_path is the local DIRECTORY that contains the file named by
  # file_name, on the machine running `terraform apply`. The uploaded file is
  # "<file_location_path>/<file_name>" (here /path/to/your/certs/tf_ssl_cert.pem)
  # and is stored on the SDX appliance under file_name.
  file_location_path = "/path/to/your/certs"
}
