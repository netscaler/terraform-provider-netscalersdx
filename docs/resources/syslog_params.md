---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_syslog_params Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for Syslog Parameters resource.
---

# netscalersdx_syslog_params (Resource)

Configuration for Syslog Parameters resource.

## Example Usage

```terraform
resource "netscalersdx_syslog_params" "tf_syslog_params" {
  date_format = "DDMMYYYY"
  timezone    = "GMT"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `date_format` (String) Format of date to be added in the syslog message.
- `timezone` (String) Timezone to be used in the syslog message.

### Read-Only

- `id` (String) The ID of this resource.
