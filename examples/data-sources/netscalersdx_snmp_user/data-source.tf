
data "netscalersdx_snmp_user" "name" {
  name = "hello"
}
output "id" {
  value = data.netscalersdx_snmp_user.is
}