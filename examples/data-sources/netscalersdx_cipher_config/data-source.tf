
data "netscalersdx_cipher_config" "name" {
  id = "cipher_config-a54b680f01ac46962dc13870a4ba8cc"
}

output "name" {
  value = data.netscalersdx_cipher_config.name.cipher_name_list_array
}