resource "netscalersdx_ssl_key" "tf_ssl_key" {
  file_name = "tf_ssl_key.key"
  # file_location_path is the local DIRECTORY that contains the file named by
  # file_name, on the machine running `terraform apply`. The uploaded file is
  # "<file_location_path>/<file_name>" (here /path/to/your/keys/tf_ssl_key.key)
  # and is stored on the SDX appliance under file_name.
  file_location_path = "/path/to/your/keys"
}

data "netscalersdx_ssl_key" "data_tf_ssl_key" {
  id = netscalersdx_ssl_key.tf_ssl_key.id
}

output "ssl_key_file_size" {
  value = data.netscalersdx_ssl_key.data_tf_ssl_key.file_size
}
