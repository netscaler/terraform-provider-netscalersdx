data "netscalersdx_ntp_server" "ntp_server" {
  server = "10.10.10.11"
}
output "ntp_servers" {
  value = data.netscalersdx_ntp_server.ntp_server.id
}