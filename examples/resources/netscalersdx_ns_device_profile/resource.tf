resource "netscalersdx_ns_device_profile" "tf_ns_device_profile" {
  name                                         = "tf_ns_device_profile"
  password                                     = "Verysecret@123"
  username                                     = "nsroot"
  use_global_setting_for_communication_with_ns = true
  type                                         = "ns"
  http_port                                    = 80
  https_port                                   = 443
  snmpversion                                  = "v3"
  snmpsecuritylevel                            = "NoAuthNoPriv"
  snmpsecurityname                             = "device-profile"
  ssh_port                                     = 22
  svm_ns_comm                                  = "http"
}
