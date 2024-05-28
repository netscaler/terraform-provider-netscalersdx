
resource "netscalersdx_provision_vpx" "device1" {
  name                   = "device1"
  ip_address             = "10.10.10.11"
  ipv4_address           = "10.10.10.11"
  netmask                = "255.255.255.0"
  gateway                = "10.10.10.12"
  if_internal_ip_enabled = false
  config_type            = 0

  image_name   = "NSVPX-XEN-13.1-17.42_nc_64.xva"
  profile_name = "nsroot_Verysecret"
  description  = "from tf"

  # License Allocation
  license                    = "Standard"
  throughput_allocation_mode = 0
  throughput                 = 1000
  max_burst_throughput       = 0
  burst_priority             = 0

  # Crypto Allocation
  number_of_acu = 0
  number_of_scu = 0

  # Resource Allocation
  vm_memory_total = 2048
  pps             = 1000000
  number_of_cores = 0

  # Instance Administration
  # username   = "vpxadmin"
  # password   = "secret"
  # cmd_policy = true

  # Network Settings
  l2_enabled = false

  # When no Management Channel created for interfaces
  if_0_1      = true
  vlan_id_0_1 = 0 # VLAN Tag
  if_0_2      = true
  vlan_id_0_2 = 0 # VLAN Tag

  # When Management Channel created for interfaces
  # la_mgmt     = true
  # vlan_id_0_1 = 10  # VLAN tag

  # Data Interfaces (Network Settings)
  network_interfaces = [
    {
      port_name            = "LA/1"
      mac_mode             = "default"
      receiveuntagged      = true
      vlan_whitelist_array = ["2", "3", "5"]
    },
    {
      port_name            = "LA/2"
      mac_mode             = "default"
      receiveuntagged      = true
      vlan_whitelist_array = ["4000", "4001", "4005"] // individual vlan id (Ascending order)
    },
    {
      port_name            = "LA/3"
      mac_mode             = "default"
      receiveuntagged      = true
      vlan_whitelist_array = ["100-105", "4000-4004"] // maintain the order as well (Ascending order)
    }
  ]

  # Management VLAN settings
  nsvlan_id         = 0
  vlan_type         = 1
  nsvlan_tagged     = false
  nsvlan_interfaces = []
}
