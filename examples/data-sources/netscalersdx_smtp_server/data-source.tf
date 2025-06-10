
data "netscalersdx_smtp_server" "tf_smtp_server" {
  server_name = "tf_smtp_server"
}

output "server_id" {
  value = data.netscalersdx_smtp_server.name.id
}