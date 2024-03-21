resource "netscalersdx_mpsuser" "tf_mpsuser" {
  name                    = "tf_mpsuser"
  password                = "VerySecret@1234"
  external_authentication = "false"
  groups                  = ["read_only"]
  session_timeout        = "20"
  session_timeout_unit   = "Minutes"
  enable_session_timeout = "true"
}
