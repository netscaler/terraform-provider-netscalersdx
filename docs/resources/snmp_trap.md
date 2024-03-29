---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_snmp_trap Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for SNMP Trap Destinations resource.
---

# netscalersdx_snmp_trap (Resource)

Configuration for SNMP Trap Destinations resource.

## Example Usage

```terraform
resource "netscalersdx_snmp_trap" "tf_snmp_trap" {
  community   = "public"
  version     = "v2"
  dest_server = "10.10.10.10"
  dest_port   = 163
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dest_server` (String) Trap Destination Server Address.

### Optional

- `community` (String) Community Name. Maximum length =  32
- `dest_port` (Number) Destination Port. Minimum value =  1 Maximum value =
- `user_name` (List of String) Name of SNMP Trap User. Minimum length =  1 Maximum length =  32
- `version` (String) SNMP version. Maximum length =  2

### Read-Only

- `id` (String) The ID of this resource.
