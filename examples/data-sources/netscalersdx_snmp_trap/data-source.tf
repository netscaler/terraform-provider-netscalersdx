data "netscalersdx_snmp_trap" "demo" {
  dest_server = "10.10.10.10"
}

output "name" {
  value = data.netscalersdx_snmp_trap.demo.id
}