resource "netscalersdx_ntp_param" "tf_ntp_param" {
  automax_logsec = 12
  revoke_logsec  = 16
  authentication = true
}
