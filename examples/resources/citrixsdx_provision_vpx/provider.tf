terraform {
  required_providers {
    citrixsdx = {
      source = "citrix/citrixsdx"
    }
  }
}
provider "citrixsdx" {
  host       = "https://10.10.10.10" # Optionally use CITRIXSDX_HOST env var
  username   = "nsroot"                # Optionally use CITRIXSDX_USERNAME env var
  password   = "secretpassword"         # Optionally use CITRIXSDX_PASSWORD env var
  ssl_verify = false                   # Optionally use CITRIXSDX_SSL_VERIFY env var
}
