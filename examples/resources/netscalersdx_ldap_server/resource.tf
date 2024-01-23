resource "netscalersdx_ldap_server" "tf_ldap_server" {
  name                       = "tf_ldap_server"
  ip_address                 = "10.10.10.10"
  sec_type                   = "PLAINTEXT"
  type                       = "AD"
  port                       = 389
  auth_timeout               = "3"
  validate_ldap_server_certs = false
}
