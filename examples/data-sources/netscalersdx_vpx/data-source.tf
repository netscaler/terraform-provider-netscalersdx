data "netscalersdx_vpx" "instance1" {
  ip_address = "10.10.10.11"
}

output "instance1_id" {
  value = data.netscalersdx_vpx.instance1.id
}