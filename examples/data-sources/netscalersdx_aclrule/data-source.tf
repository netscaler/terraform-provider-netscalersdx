data "netscalersdx_aclrule" "test" {
  name = "tf_aclrule"
}
output "acl_id" {
  value = data.netscalersdx_aclrule.test.id
}
