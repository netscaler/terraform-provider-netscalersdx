data "netscalersdx_system_settings" "test" {
  id = "tf-system_settings-12345"
}

output "system_settings_result" {
  value = data.netscalersdx_system_settings.test
}