data "netscalersdx_mps_ssl_certkey" "test" {
  id = "example-id"
}
output "certificate_name" {
  value = data.netscalersdx_mps_ssl_certkey.test.ssl_certificate
}