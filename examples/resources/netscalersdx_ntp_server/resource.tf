resource "netscalersdx_ntp_server" "ntp_server" {
  server  = "10.10.10.11"
  key_id  = 123
  minpoll = 5
  maxpoll = 12
  autokey = false
}
