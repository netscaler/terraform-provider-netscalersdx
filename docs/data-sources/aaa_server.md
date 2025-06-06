---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_aaa_server Data Source - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Retrieve an AAA server by its ID.
---

# netscalersdx_aaa_server (Data Source)

Retrieve an AAA server by its ID.

## Example Usage

```terraform
resource "netscalersdx_radius_server" "tf_radius_server" {
  name         = "tf_radius_server"
  ip_address   = "10.10.10.10"
  radius_key   = "Verysecretkey"
  port         = 389
  auth_timeout = "3"
}
resource "netscalersdx_tacacs_server" "tf_tacacs_server" {
  name       = "tf_tacacs_server"
  port       = 545
  tacacs_key = "key"
  ip_address = "10.10.10.11"
}

resource "netscalersdx_aaa_server" "tf_aaa_server" {
  fallback_local_authentication = "true"
  primary_server_type           = "RADIUS"
  primary_server_name           = netscalersdx_radius_server.tf_radius_server.name
  external_servers = [
    {
      external_server_type = "TACACS"
      priority             = 3
      external_server_name = netscalersdx_tacacs_server.tf_tacacs_server.name
    }
  ]
}

data "netscalersdx_aaa_server" "data_tf_aaa_server" {
  id = netscalersdx_aaa_server.tf_aaa_server.id
}

output "primary_server_type" {
  value = data.netscalersdx_aaa_server.data_tf_aaa_server.primary_server_type
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `external_servers` (Attributes List) List of external servers. (see [below for nested schema](#nestedatt--external_servers))
- `fallback_local_authentication` (Boolean) Enable local fallback authentication.
- `log_ext_group_info` (Boolean) Log external group info.
- `primary_server_name` (String) Name of primary server name. Minimum length =  1 Maximum length =  128
- `primary_server_type` (String) Type of primary server. Supported types 1. LOCAL 2.RADIUS 3.LDAP 4.TACACS 5.KEYSTONE. Minimum length =  1 Maximum length =  32

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--external_servers"></a>
### Nested Schema for `external_servers`

Required:

- `external_server_name` (String) Name of external server. Minimum length =  1 Maximum length =  128
- `external_server_type` (String) Type of external server. Supported types 1.RADIUS 2.LDAP 3.TACACS 4.KEYSTONE. Minimum length =  1 Maximum length =  32
- `priority` (Number) Priority of external server. Minimum value =  2 Maximum value =
