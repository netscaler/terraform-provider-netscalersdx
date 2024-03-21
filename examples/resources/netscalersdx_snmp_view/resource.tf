resource "netscalersdx_snmp_view" "tf_snmp_view" {
  name    = "tf_snmp_view"
  subtree = "1.3.6.1.2.1.1"
  type    = "false"
}
