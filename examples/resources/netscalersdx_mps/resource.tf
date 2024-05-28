resource "netscalersdx_mps" "tf_mps" {
  is_cloud                   = "false"
  is_passive                 = "false"
  is_member_of_default_group = "true"
  is_thirdparty_vm_supported = "false"
  is_container               = "false"
  hist_mig_inprog            = "false"
  config_motd                = "true"
  motd                       = "Welcome to NetScaler SDX"
}
