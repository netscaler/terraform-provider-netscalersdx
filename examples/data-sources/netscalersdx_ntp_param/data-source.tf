data "netscalersdx_ntp_param" "tf_ntp_param" {
  id = "tf_ntp_param"
}

output "automax_logsec" {
  value = data.netscalersdx_ntp_param.tf_ntp_param.automax_logsec
}