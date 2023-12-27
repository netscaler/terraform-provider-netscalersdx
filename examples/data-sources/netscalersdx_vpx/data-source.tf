data "netscalersdx_vpx" "instance1" {
  ip_address = "10.222.74.177"
}

output "instance1_id" {
  value = data.netscalersdx_vpx.instance1.id
}