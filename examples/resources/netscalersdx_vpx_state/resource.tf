data "netscalersdx_vpx" "instance1" {
  ip_address = "10.10.10.176"
}

resource "netscalersdx_vpx_state" "stop_device1" {
  vpx_id = data.netscalersdx_vpx.instance1.id
  state  = "stop" # "start" | "stop" | "force_stop" | "reboot" | "force_reboot"
}