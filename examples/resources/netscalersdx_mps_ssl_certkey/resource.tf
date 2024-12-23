resource "netscalersdx_mps_ssl_certkey" "tf_mps_ssl_certkey" {
  ssl_certificate = "duplicate_cert"
  ssl_key         = "duplicate_key"
}
