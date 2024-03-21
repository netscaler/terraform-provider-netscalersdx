resource "netscalersdx_snmp_alarm_config" "tf_snmp_alarm_config" {
  name      = "cpuTempError"
  severity  = "Critical"
  threshold = "60"
  enable    = "true"
}
