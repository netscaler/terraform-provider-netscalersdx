
data "netscalersdx_radius_server" "tf_radius_server" {
  name = "tf_radius_server"
}

output "name" {
  value = data.netscalersdx_radius_server.tf_radius_server.name
}

output "ip_address" {
  value = data.netscalersdx_radius_server.tf_radius_server.ip_address
}