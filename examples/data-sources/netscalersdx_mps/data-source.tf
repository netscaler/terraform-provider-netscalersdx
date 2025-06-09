data "netscalersdx_mps" "tf_mps" {
  id = "tf_mps"
}

output "result" {
  value = data.netscalersdx_mps.tf_mps.motd
}