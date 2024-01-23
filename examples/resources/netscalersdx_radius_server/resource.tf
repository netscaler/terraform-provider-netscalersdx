resource "netscalersdx_radius_server" "tf_radius_server" {
  name         = "tf_radius_server"
  ip_address   = "10.10.10.10"
  radius_key   = "Verysecretkey"
  port         = 389
  auth_timeout = "3"
}
