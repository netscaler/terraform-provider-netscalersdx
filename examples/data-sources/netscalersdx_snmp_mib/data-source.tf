data "netscalersdx_snmp_mib" "demo" {
  name = "tf-testing"
}

output "snmp_mib_id" {
  value = data.netscalersdx_snmp_mib.demo.id
}