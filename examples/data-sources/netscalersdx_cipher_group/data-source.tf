data "netscalersdx_cipher_group" "tf_cipher_group" {
  id = "tf_cipher_group"
}

output "cipher_group_name" {
  value = data.netscalersdx_cipher_group.tf_cipher_group.cipher_group_name
}