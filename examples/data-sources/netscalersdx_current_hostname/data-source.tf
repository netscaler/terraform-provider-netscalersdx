data "netscalersdx_current_hostname" "test" {
  id  = "tf-test-current-hostname"
}
output "current_hostname" {
  value = data.netscalersdx_current_hostname.test.hostname
}