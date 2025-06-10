data "netscalersdx_blx_device_profile" "test" {
  name = "tf_test_blx_device_profile"
}
output "tf_test_blx_device_profile_id" {
  value = data.netscalersdx_blx_device_profile.test.id
}