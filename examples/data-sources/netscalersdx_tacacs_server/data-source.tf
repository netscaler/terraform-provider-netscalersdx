
data "netscalersdx_tacacs_server" "tacacs_server" {
  name = "tf_tacacs_server"
}
output "tf_tacacs_server_id" {
  value = data.netscalersdx_tacacs_server.tacacs_server.id
}
