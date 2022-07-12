terraform {
  required_providers {
    citrixsdx = {
      source = "citrix/citrixsdx"
    }
  }
}
provider "citrixsdx" {
  host       = "https://10.222.74.135" # Optionally use CITRIXSDX_HOST env var
  username   = "nsroot"                # Optionally use CITRIXSDX_USERNAME env var
  password   = "Notnsroot250$"         # Optionally use CITRIXSDX_PASSWORD env var
  ssl_verify = false                   # Optionally use CITRIXSDX_SSL_VERIFY env var
}
