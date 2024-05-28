resource "netscalersdx_ssl_settings" "tf_ssl_settings" {
  sslreneg = false
  tlsv1_1  = true
  sslv3    = false
  tlsv1_2  = true
  tlsv1    = false
}
