resource "netscalersdx_tacacs_server" "tf_tacacs_server" {
  name       = "tf_tacacs_server"
  port       = 545
  tacacs_key = "key"
  ip_address = "10.10.10.11"
}
