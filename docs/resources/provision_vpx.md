---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netscalersdx_provision_vpx Resource - terraform-provider-netscalersdx"
subcategory: ""
description: |-
  Configuration for Provision VPX resource
---

# netscalersdx_provision_vpx (Resource)

Configuration for Provision VPX resource

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ip_address` (String) IP Address for this managed device. Minimum length =  1 Maximum length =  64

### Optional

- `backplane` (String) Backplane Interface. Minimum length =  1
- `burst_priority` (Number) Burst Priority of the VM Instance between 1 and 4.
- `cmd_policy` (String) true if you want to allow shell/sftp/scp access to NetScaler Instance administrator. Minimum length =  1 Maximum length =  1024
- `config_type` (Number) Configuration Type. Values: 0: IPv4, 1: IPv6, 2: Both.
- `crypto_change_requires_reboot` (Boolean) `true` if the current changes made by user requires a reboot of the VM else `false`.
- `customid` (String) Custom ID.
- `datacenter_id` (String) Datacenter Id is system generated key for data center.
- `description` (String) Description of managed device. Minimum length =  1 Maximum length =  512
- `device_family` (String) Device Family. Minimum length =  1 Maximum length =  64
- `display_name` (String) Display Name for this managed device. For HA pair it will be A-B, and for Cluster it will be CLIP. Minimum length =  1 Maximum length =  128
- `domain_name` (String) Domain name of VM Device. Minimum length =  1 Maximum length =  128
- `ent_bw_available` (Number) Enterprise Bandwidth configured.
- `ent_bw_config` (Number) Enterprise Bandwidth configured.
- `ent_bw_total` (Number) Enterprise Bandwidth Total.
- `fips_partition_name` (String) FIPS Partition Name. Minimum length =  1 Maximum length =  128
- `gateway` (String) Default Gateway of managed device. Minimum length =  1 Maximum length =  64
- `gateway_ipv6` (String) Gateway IPv6 Address.
- `host_ip_address` (String) Host IPAddress where VM is provisioned. Minimum length =  1 Maximum length =  64
- `hostname` (String) Assign hostname to managed device, if this is not provided, name will be set as host name . Minimum length =  1 Maximum length =  256
- `id` (String) ID of the NetScaler ADC Instance.
- `if_0_1` (Boolean) Network 0/1 on VM Instance, Select this option to assign 0/1 Interface
- `if_0_2` (Boolean) Network 0/2 on VM Instance, Select this option to assign 0/2 Interface
- `if_internal_ip_enabled` (Boolean) Set as true if VPX is managed by internal network (not required to be set for SDWAN).
- `image_name` (String) Image Name, This parameter is used while provisioning VM Instance with XVA image, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128
- `instance_mode` (String) Denotes state- primary,secondary,clip,clusternode.
- `internal_ip_address` (String) Internal IP Address for this managed device. Minimum length =  1 Maximum length =  64
- `ipv4_address` (String) IPv4 Address. Minimum length =  1 Maximum length =  64
- `ipv6_address` (String) IPv6 Address.
- `is_clip` (Boolean) Is Clip.
- `is_managed` (Boolean) Is Managed.
- `is_new_crypto` (Boolean) `true` if number_of_acu/number_of_scu are used, `false` if number_of_ssl_cores is used.
- `iscco` (Boolean) Is CCO.
- `l2_enabled` (Boolean) L2mode status of VM Instance. Select this option to allow L2 mode on all the Data Interfaces on this NetScaler ADC Instance
- `la_mgmt` (Boolean) Bond consisting of management ports on VM Instance. When Management Channel created for interfaces, this will be set to `true`
- `last_updated_time` (Number) Last Updated Time.
- `license` (String) Feature License for NetScaler ADC Instance, needs to be set while provisioning [Possible values: Standard, Enterprise, Platinum].
- `license_edition` (String) Edition of instance.
- `license_grace_time` (Number) Grace for this NetScaler Instance..
- `mastools_version` (String) Mastools version if the device is embedded agent.
- `max_burst_throughput` (Number) Maximum burst throughput in Mbps of VM Instance.
- `metrics_collection` (Boolean) Flag to check if metrics collection is enabled or disabled..
- `mgmt_ip_address` (String) Management IP Address for this Managed Device. Minimum length =  1 Maximum length =  64
- `name` (String) Name of managed device. Minimum length =  1 Maximum length =  128
- `netmask` (String) Netmask of managed device. Minimum length =  1 Maximum length =  64
- `network_interfaces` (Attributes List) Network Interfaces. (see [below for nested schema](#nestedatt--network_interfaces))
- `nexthop` (String) Next Hop IP address. Minimum length =  1 Maximum length =  64
- `nexthop_v6` (String) Next Hop IPv6 Address.
- `node_id` (String) Node identification of a device.
- `ns_ip_address` (String) NetScaler IP Address for this managed device. Minimum length =  1 Maximum length =  128
- `nsvlan_id` (Number) VLAN for Management Traffic.
- `nsvlan_interfaces` (List of String) VLAN Interfaces. Minimum length =  1 Maximum length =  50
- `nsvlan_tagged` (Boolean) When this option is selected, selected interfaces are added as tagged members of Management VLAN
- `num_pes` (Number) Total number of PEs.
- `number_of_acu` (Number) Assign number of asymmetric crypto units to VM Instance.
- `number_of_cores` (Number) Number of cores that are assigned to VM Instance.
- `number_of_scu` (Number) Assign number of asymmetric crypto units to VM Instance.
- `number_of_ssl_cards` (Number) Number of SSL Cards.
- `number_of_ssl_cores` (Number) Assign number of ssl virtual functions to VM Instance.
- `number_of_ssl_cores_up` (Number) Number of SSL Cores Up.
- `password` (String) Password for specified user on NetScaler Instance. Minimum length =  1 Maximum length =  127
- `plt_bw_available` (Number) Platinum Bandwidth Available.
- `plt_bw_config` (Number) Platinum Bandwidth configured.
- `plt_bw_total` (Number) Total Platinum Bandwidth.
- `plugin_ip_address` (String) Signaling IP Address. Minimum length =  1 Maximum length =  64
- `plugin_netmask` (String) Signaling Netmask. Minimum length =  1 Maximum length =  64
- `pps` (Number) Assign packets per seconds to NetScaler Instance.
- `profile_name` (String) Device Profile Name that is attached with this managed device. Minimum length =  1 Maximum length =  128
- `profile_password` (String) Password specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128
- `profile_username` (String) User Name specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128
- `reboot_vm_on_cpu_change` (Boolean) Reboot VMs on CPU change during resource allocation.
- `save_config` (Boolean) Should config be saved first in case instance is rebooted while modify.
- `state` (String) Node State. Minimum length =  1 Maximum length =  32
- `std_bw_available` (Number) Standard Bandwidth Available.
- `std_bw_config` (Number) Standard Bandwidth running.
- `std_bw_total` (Number) Standard Bandwidth.
- `template_name` (String) Template Name, This parameter is used while provisioning VM Instance with template, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128
- `throughput` (Number) Assign throughput in Mbps to VM Instance.
- `throughput_allocation_mode` (Number) Throughput Allocation Mode: 0-Fixed, 1-Burst-able.
- `throughput_limit` (Number) Throughput Limit in Mbps set for VM Instance.
- `type` (String) Type of device, (Xen | NS). Minimum length =  1 Maximum length =  64
- `username` (String) User Name (except nsroot) to be configured on NetScaler Instance. Minimum length =  1 Maximum length =  127
- `vcpu_config` (Number) Number of vCPU allocated for the device.
- `vlan_id_0_1` (Number) VLAN id for the management interface 0/1. This VLAN ID is used to filter management traffic on 0/1 at hypervisor layer.
- `vlan_id_0_2` (Number) VLAN id for the management interface 0/2. This VLAN ID is used to filter management traffic on 0/2 at hypervisor layer.
- `vlan_type` (Number) VLAN Type, NetScaler or L2 VLAN. Select 0 for NetScaler VLAN or 1 for L2 VLAN.
- `vm_memory_total` (Number) Total Memory of VM Instance in MB. 2048MB, 5120MB.

<a id="nestedatt--network_interfaces"></a>
### Nested Schema for `network_interfaces`

Required:

- `port_name` (String) Port name of the interface on the host machine.

Optional:

- `gateway` (String) Gateway
- `interface_name` (String) Interface Name
- `ip_address` (String) IP Address
- `is_member_ifc` (Boolean) `true` if this interface is member of a channel.
- `is_mgmt_ifc` (Boolean) `true` if this is the management interface.
- `is_vlan_applied` (Boolean) Is VLAN added on NetworkInterface of VM Instance.
- `l2_enabled` (Boolean) L2 mode status of Interface.
- `mac_address` (String) MAC Address
- `mac_mode` (String) MAC Mode, The method according to which MAC Address is assigned to Interface. Possible values: [default, generated, custom] default: XenServer assigns a MAC Address. custom: SDX Administrator assigns a MAC address. generated: Generate a MAC address by using the base MAC address set at System Level.
- `managed_device_id` (String) Managed Device Id
- `name_server` (String) Name Server
- `netmask` (String) Netmask
- `network_interface_id` (String) Id
- `parent_id` (String) Parent Id
- `parent_name` (String) Parent Name
- `receiveuntagged` (Boolean) Receive Untagged Packets on Interface/Channel. Allow Untagged Traffic.
- `sdx_formation_network_id` (String) Sdx Formation Network Id
- `vlan` (Number) VLAN.
- `vlan_whitelist` (String) VLAN Whitelist.
- `vlan_whitelist_array` (List of String) Allowed VLANs. Range of VLANs can be provided using hyphen '-' separater and separated VLANs can also be provided. (e.g., ["100-110","142","143","151-155"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VLANs if they are not in sequence or if the sequence is 2 or fewer (e.g., ["100","101","4000","4001"]). If the VLANs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., ["100-103","4000-4002"]). Also, maintain the order as well (Ascending order) (e.g., ["100-103","200","4000-4002"])
- `vrid_list_ipv4` (String) VRID List for Interface/Channel for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., "100-110,142,143,151-155").
- `vrid_list_ipv4_array` (List of String) VRID List for Interface for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., ["100-110","142","143","151-155"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., ["100","101","4000","4001"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., ["100-103","4000-4002"]). Also, maintain the order as well (Ascending order) (e.g., ["100-103","200","4000-4002"])
- `vrid_list_ipv6` (String) VRID List for Interface/Channel for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., "100-110,142,143,151-155").
- `vrid_list_ipv6_array` (List of String) VRID List for Interface for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., ["100-110","142","143","151-155"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., ["100","101","4000","4001"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., ["100-103","4000-4002"]). Also, maintain the order as well (Ascending order) (e.g., ["100-103","200","4000-4002"])
