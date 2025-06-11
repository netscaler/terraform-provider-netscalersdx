data "netscalersdx_ssl_settings" "test" {
  id = "test-tf_ssl_settings"
}
output "tf_ssl_settings" {
  value = data.netscalersdx_ssl_settings.test
}