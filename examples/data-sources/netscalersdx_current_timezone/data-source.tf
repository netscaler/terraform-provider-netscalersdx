data "netscalersdx_current_timezone" "tf_current_timezone" {
  # Some random id to satisfy the provider's requirement for an id.
  id = "tf-test-current-timezone"
}
output "timezone" {
  value = data.netscalersdx_current_timezone.tf_current_timezone.timezone
}