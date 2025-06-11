data "netscalersdx_device_profile" "test" {
  name = "xen_nsroot_profile"
}
output "name" {
  value = data.netscalersdx_device_profile.test.id
}