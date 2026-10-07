# Standard offline activation
resource "netscalersdx_nslaslicense_offline" "license" {
  entitlement_name = "SDX 9195 Premium"
  las_secrets_json = "${path.module}/las_secrets.json"

  # Pinned SDX SSH host key used to verify the appliance identity for the SCP
  # license transfer. Capture it once with: ssh-keyscan -t rsa <sdx-mgmt-ip>
  ssh_host_pubkey = "ssh-rsa AAAAB3NzaC1yc2E..."
}

# Restricted offline activation (no file upload to the LAS service)
resource "netscalersdx_nslaslicense_offline" "license_restricted" {
  entitlement_name = "SDX 9195 Premium"
  las_secrets_json = "${path.module}/las_secrets.json"
  ssh_host_pubkey  = "ssh-rsa AAAAB3NzaC1yc2E..."
  restricted_mode  = true
}
