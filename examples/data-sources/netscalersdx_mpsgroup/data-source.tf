data "netscalersdx_mpsgroup" "tf_mpsgroup" {
  name = "tf_mpsgroup"
}

output "mpsgroup_name" {
  value = data.netscalersdx_mpsgroup.tf_mpsgroup.id
}
