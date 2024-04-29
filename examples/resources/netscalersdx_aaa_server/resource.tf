resource "netscalersdx_radius_server" "tf_radius_server" {
  name         = "tf_radius_server"
  ip_address   = "10.10.10.10"
  radius_key   = "Verysecretkey"
  port         = 389
  auth_timeout = "3"
}
resource "netscalersdx_tacacs_server" "tf_tacacs_server" {
  name       = "tf_tacacs_server"
  port       = 545
  tacacs_key = "key"
  ip_address = "10.10.10.11"
}

resource "netscalersdx_aaa_server" "tf_aaa_server" {
  fallback_local_authentication = "true"
  primary_server_type           = "RADIUS"
  primary_server_name           = netscalersdx_radius_server.tf_radius_server.name
  external_servers = [
    {
      external_server_type = "TACACS"
      priority             = 3
      external_server_name = netscalersdx_tacacs_server.tf_tacacs_server.name
    }
  ]
}
