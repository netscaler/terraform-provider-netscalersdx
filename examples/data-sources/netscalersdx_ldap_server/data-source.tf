
data "netscalersdx_ldap_server" "tf_ldap_server" {
  name = "tf_ldap_server"
}

output "name" {
  value = data.netscalersdx_ldap_server.tf_ldap_server.id
}