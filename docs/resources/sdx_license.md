---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_sdx_license Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Upload and apply license file to the Citrix SDX appliance.
---

# netscalersdx_sdx_license (Resource)

Upload and apply license file to the Citrix SDX appliance.

## Example Usage

```terraform
resource "netscalersdx_sdx_license" "tf_sdx_license" {
  file_name = "temp.lic"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `file_name` (String) License file name to be uploaded and applied. Note: License file should be present in the current directory where the Terraform configuration file is located

### Read-Only

- `id` (String) The ID of this resource.
