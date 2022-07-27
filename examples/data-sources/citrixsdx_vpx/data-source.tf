data "citrixsdx_vpx" "instance1" {
  ip_address = "10.222.74.176"
}

output "instance1_id" {
  value = data.citrixsdx_vpx.instance1.id
}