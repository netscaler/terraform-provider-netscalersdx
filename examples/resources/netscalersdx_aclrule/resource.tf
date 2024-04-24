resource "netscalersdx_aclrule" "tf_aclrule" {
  name     = "tf_aclrule"
  priority = 100
  protocol = "TCP"
  action   = "Allow"
  dst_port = 80
  src_ip   = "10.10.10.10"
}
