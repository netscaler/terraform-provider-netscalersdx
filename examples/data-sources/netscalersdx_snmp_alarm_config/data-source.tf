data "netscalersdx_snmp_alarm_config" "example" {
  name = "cpuTempError"
}

output "snmp_alarm_config_id" {
  value = data.netscalersdx_snmp_alarm_config.example.id
}