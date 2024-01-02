---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_cipher_group Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for Cipher Group resource
---

# netscalersdx_cipher_group (Resource)

Configuration for Cipher Group resource

## Example Usage

```terraform
resource "netscalersdx_cipher_group" "tf_cipher_group" {
  cipher_group_description = "from terraform"
  cipher_group_name        = "tf_cipher_group"
  cipher_name_list_array   = ["TLS1-AES-256-CBC-SHA", "TLS1-ECDHE-RSA-DES-CBC3-SHA"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cipher_group_description` (String) Describing the Cipher Group algorithms created. Minimum length =  1 Maximum length =  256
- `cipher_group_name` (String) Name of Cipher Group. Minimum length =  1 Maximum length =  128
- `cipher_name_list_array` (List of String) list of cipher suites in form of array of strings.

### Read-Only

- `id` (String) The ID of this resource.