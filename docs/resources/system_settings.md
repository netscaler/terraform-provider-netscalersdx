---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_system_settings Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for System Settings resource.
---

# netscalersdx_system_settings (Resource)

Configuration for System Settings resource.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `authorize_deviceapiproxy` (Boolean) Authorize the DeviceAPIProxy request.
- `basicauth` (Boolean) Allow Basic Authentication Protocol.
- `disable_agent_old_password_input` (Boolean) Disable old password input requirement while changing ADM agent password.
- `disk_utilization_threshold` (Number) Disk utilization threshold after which data processing it stopped.
- `enable_apiproxy_credentials` (Boolean) Enable API Proxy Credentials.
- `enable_certificate_download` (Boolean) Enable Certificate Download.
- `enable_cuxip` (Boolean) Used to enable/disable CUXIP(Customer User Experience Improvement Program).
- `enable_delete_interface_on_adc` (Boolean) Flag to enable/disable deleting interface from ADCs on SDX.
- `enable_nsrecover_login` (Boolean) This setting enalbes nsrecover login for SVM.
- `enable_session_timeout` (Boolean) Enables session timeout feature.
- `enable_shell_access` (Boolean) Enable Shell access for non-nsroot User(s).
- `is_metering_enabled` (Boolean) Enable Metering for NetScaler VPX on SDX.
- `keep_adc_image_count` (Number) Count for number of NetScaler images to be saved in Agent.
- `keep_alive_ping_interval` (Number) Agent web socket keep alive ping interval for the system.
- `prompt_creds_for_stylebooks` (Boolean) Prompt Credentials for Stylebooks.
- `secure_access_only` (Boolean) Secure Access only.
- `session_timeout` (Number) Session timeout for the system.
- `session_timeout_unit` (String) Session timeout unit for the system. Possible Values: [ Minutes, Hours ]
- `svm_ns_comm` (String) Communication with Instances. Minimum length =  1 Maximum length =  10. Possible Values: [ http, https ]

### Read-Only

- `id` (String) The ID of this resource.
