terraform {
  required_providers {
    netscalersdx = {
      source = "netscaler/netscalersdx"
    }
  }
}
provider "netscalersdx" {
  host       = "https://10.10.10.10" # Optionally use NETSCALERSDX_HOST env var
  username   = "nsroot"              # Optionally use NETSCALERSDX_USERNAME env var
  password   = "secretpassword"      # Optionally use NETSCALERSDX_PASSWORD env var
  ssl_verify = false                 # Optionally use NETSCALERSDX_SSL_VERIFY env var
}
