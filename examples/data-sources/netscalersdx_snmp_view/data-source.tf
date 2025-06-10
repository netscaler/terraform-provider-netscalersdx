data "netscalersdx_snmp_view" "test" {
  name = "tf_snmp_view"
}

output "netscalersdx_snmp_view_id" {
  value = data.netscalersdx_snmp_view.test.id
}