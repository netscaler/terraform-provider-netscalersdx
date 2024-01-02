resource "netscalersdx_syslog_server" "tf_syslog_server" {
  name           = "tf_syslog_server"
  ip_address     = "10.10.10.10"
  port           = 514
  log_level_all  = true
  log_level_none = false
}
