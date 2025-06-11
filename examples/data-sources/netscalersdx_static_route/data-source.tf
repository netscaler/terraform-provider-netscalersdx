data "netscalersdx_static_route" "tf_test" {
  # id = "1234-asdf-1234-asdd"
  network = "10.10.10.10"
}
output "static_route_id" {
  value = data.netscalersdx_static_route.tf_test.id
}
