data "netscalersdx_ns_device_profile" "tf_ns_device_profile" {
  name = "ns_nsroot_profile"
}

output "name" {
  value = data.netscalersdx_ns_device_profile.tf_ns_device_profile.id
}