
data "netscalersdx_mpsuser" "tf_mpsuser" {
  name = "tf-test"
}
output "mpsuser_id" {
  value = data.netscalersdx_mpsuser.tf_mpsuser.id
}