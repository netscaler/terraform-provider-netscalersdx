data "netscalersdx_snmp_manager" "tf_demo" {
  ip_address = "10.10.10.10"
}

output "name" {
  value = data.netscalersdx_snmp_manager.tf_demo.id
}