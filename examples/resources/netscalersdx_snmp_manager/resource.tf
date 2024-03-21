resource "netscalersdx_snmp_manager" "tf_snmp_manager" {
  ip_address = "10.10.10.10"
  community  = "public"
  netmask    = "255.255.255.0"
}
