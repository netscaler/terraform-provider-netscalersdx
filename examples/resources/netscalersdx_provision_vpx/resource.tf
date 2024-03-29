resource "netscalersdx_provision_vpx" "device1" {
  name                   = "device1"
  ip_address             = "10.10.10.11"
  ipv4_address           = "10.10.10.11"
  netmask                = "255.255.255.0"
  gateway                = "10.10.10.12"
  if_internal_ip_enabled = false
  config_type            = 0

  nexthop      = ""
  image_name   = "NSVPX-XEN-13.1-17.42_nc_64.xva"
  profile_name = "nsroot_Verysecret"
  description  = "from tf"

  # License Allocation
  license                    = "Standard"
  throughput_allocation_mode = "0"
  throughput                 = "1000"
  max_burst_throughput       = "0"
  burst_priority             = "0"

  # Crypto Allocation
  number_of_acu = 0
  number_of_scu = "0"

  # Resource Allocation
  vm_memory_total = 2048
  pps             = 1000000
  number_of_cores = "0"

  # Instance Administration
  # username = "vpxadmin"
  # password = "secret"
  # cmd_policy = "true"

  # Network Settings
  l2_enabled  = "false"
  if_0_1      = true
  vlan_id_0_1 = "0"
  if_0_2      = true
  vlan_id_0_2 = "0"

  # Data Interfaces (Network Settings)
  network_interfaces = [{
    port_name            = "LA/1"
    mac_mode             = "default"
    receiveuntagged      = "true"
    vlan_whitelist       = "2,3,5"
    vlan_whitelist_array = [2, 3, 5]
    },
    {
      port_name            = "LA/2"
      mac_mode             = "default"
      receiveuntagged      = "true"
      vlan_whitelist       = ""
      vlan_whitelist_array = []
      is_vlan_applied      = false
    },
  ]

  # Management VLAN settings
  nsvlan_id         = "0"
  vlan_type         = 1
  nsvlan_tagged     = "false"
  nsvlan_interfaces = []
}
