resource "netscalersdx_snmp_trap" "tf_snmp_trap" {
  community   = "public"
  version     = "v2"
  dest_server = "10.10.10.10"
  dest_port   = 163
}
