resource "netscalersdx_blx_device_profile" "tf_blx_device_profile" {
  name                                         = "tf_blx_device_profile"
  username                                     = "nsroot"
  password                                     = "Verysecret@123"
  host_username                                = "root"
  host_password                                = "Verysecret@123"
  use_global_setting_for_communication_with_ns = true
  type                                         = "blx"
  http_port                                    = 80
  https_port                                   = 443
  snmpversion                                  = "v3"
  snmpsecuritylevel                            = "NoAuthNoPriv"
  snmpsecurityname                             = "device-profile"
  ssh_port                                     = 22
  svm_ns_comm                                  = "http"
}
