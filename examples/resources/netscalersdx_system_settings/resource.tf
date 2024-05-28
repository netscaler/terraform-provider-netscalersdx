resource "netscalersdx_system_settings" "tf_system_settings" {
  secure_access_only               = false
  svm_ns_comm                      = "http"
  enable_cuxip                     = true
  session_timeout_unit             = "Minutes"
  enable_certificate_download      = true
  basicauth                        = true
  enable_session_timeout           = false
  enable_apiproxy_credentials      = false
  enable_shell_access              = true
  is_metering_enabled              = true
  disable_agent_old_password_input = false
  enable_nsrecover_login           = true
  enable_delete_interface_on_adc   = false
}

