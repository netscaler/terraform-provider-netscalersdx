---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_snmp_alarm_config Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for SNMP Alarm Configurations resource.
---

# netscalersdx_snmp_alarm_config (Resource)

Configuration for SNMP Alarm Configurations resource.

## Example Usage

```terraform
resource "netscalersdx_snmp_alarm_config" "tf_snmp_alarm_config" {
  name      = "cpuTempError"
  severity  = "Critical"
  threshold = "60"
  enable    = "true"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Alarm Name. Maximum length =  128

### Optional

- `enable` (Boolean) Enable Alarm.
- `severity` (String) Alarm severity. Supported values: Critical, Major, Minor, Warning, Informational . Maximum length =  128
- `threshold` (Number) Threshold Value for the alarm.
- `time` (Number) Frequency of the alarm in minutes.

### Read-Only

- `id` (String) The ID of this resource.