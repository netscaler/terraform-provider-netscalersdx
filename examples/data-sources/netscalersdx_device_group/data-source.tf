data "netscalersdx_device_group" "tf_device_group" {
  name = "tf_device_group"
}

output "resource_id" {
  value = data.netscalersdx_device_group.tf_device_group.id
}