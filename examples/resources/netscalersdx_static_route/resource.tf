resource "netscalersdx_static_route" "tf_static_route" {
  network = "10.10.10.166"
  gateway = "10.10.10.166"
  netmask = "255.255.255.0"
}
