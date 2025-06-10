data "netscalersdx_syslog_server" "test" {
  name = "tf_syslog_server"
}

output "syslog_server_id" {
  value = data.netscalersdx_syslog_server.test.id
}