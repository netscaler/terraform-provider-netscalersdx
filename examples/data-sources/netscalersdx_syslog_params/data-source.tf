data "netscalersdx_syslog_params" "test" {
  id = "tf-id-123"
}

output "date_format" {
  value = data.netscalersdx_syslog_params.test.date_format
}