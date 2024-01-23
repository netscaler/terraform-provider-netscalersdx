resource "netscalersdx_cipher_group" "tf_cipher_group" {
  cipher_group_description = "from terraform"
  cipher_group_name        = "tf_cipher_group"
  cipher_name_list_array   = ["TLS1-AES-256-CBC-SHA", "TLS1-ECDHE-RSA-DES-CBC3-SHA"]
}
