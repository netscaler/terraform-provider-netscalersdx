package citrixsdx

import (
	"context"
	"errors"
	"log"
	"time"

	"terraform-provider-citrixsdx/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProvisionVpx() *schema.Resource {
	return &schema.Resource{
		Description:   "Configuration forProvision VPX resource",
		CreateContext: resourceProvisionVpxCreate,
		ReadContext:   resourceProvisionVpxRead,
		UpdateContext: resourceProvisionVpxUpdate,
		DeleteContext: resourceProvisionVpxDelete,
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Description: "IP Address.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"std_bw_config": {
				Description: "Standard Bandwidth running.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ns_ip_address": {
				Description: "Citrix ADC IP Address for this provision VPX. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password for specified user on Citrix ADC Instance. Minimum length =  1 Maximum length =  127",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"gateway_ipv6": {
				Description: "Gateway IPv6 Address.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"throughput_limit": {
				Description: "Throughput Limit in Mbps set for VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_id_0_1": {
				Description: "VLAN id for the management interface 0/1. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of provision VPX. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mastools_version": {
				Description: "Mastools version if the device is embedded agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"plugin_ip_address": {
				Description: "Signaling IP Address. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_type": {
				Description: "VLAN Type, Citrix ADC or L2 VLAN. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ent_bw_total": {
				Description: "Enterprise Bandwidth Total.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vcpu_config": {
				Description: "Number of vCPU allocated for the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nsvlan_tagged": {
				Description: "NSVLAN Tagged.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"netmask": {
				Description: "Netmask of provision VPX. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ent_bw_config": {
				Description: "Enterprise Bandwidth configured.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"datacenter_id": {
				Description: "Datacenter Id is system generated key for data center.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"instance_mode": {
				Description: "Denotes state- primary,secondary,clip,clusternode.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_ssl_cores_up": {
				Description: "Number of SSL Cores Up.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"std_bw_available": {
				Description: "Standard Bandwidth Available.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"internal_ip_address": {
				Description: "Internal IP Address for this provision VPX. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_0_1": {
				Description: "Network 0/1 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"plt_bw_total": {
				Description: "Total Platinum Bandwidth.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_ip_address": {
				Description: "Host IPAddress where VM is provisioned. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vpx_id": {
				Description: "Id is system generated key for all the provision VPXs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv6_address": {
				Description: "IPv6 Address.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mgmt_ip_address": {
				Description: "Management IP Address for thisProvision VPX. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_acu": {
				Description: "Assign number of asymmetric crypto units to VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"plt_bw_available": {
				Description: "Platinum Bandwidth Available.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_clip": {
				Description: "Is Clip.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_family": {
				Description: "Device Family. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Type of device, (Xen | NS). Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"throughput": {
				Description: "Assign throughput in Mbps to VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"template_name": {
				Description: "Template Name, This parameter is used while provisioning VM Instance with template, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Description: "Default Gateway of provision VPX. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"iscco": {
				Description: "Is CCO.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_scu": {
				Description: "Assign number of asymmetric crypto units to VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license": {
				Description: "Feature License for Citrix ADC Instance, needs to be set while provisioning (standard, enterprise, platinum). Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"domain_name": {
				Description: "Domain name of VM Device. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_0_2": {
				Description: "Network 0/2 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"image_name": {
				Description: "Image Name, This parameter is used while provisioning VM Instance with XVA image, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hostname": {
				Description: "Assign hostname to provision VPX, if this is not provided, name will be set as host name . Minimum length =  1 Maximum length =  256",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vm_memory_total": {
				Description: "Total Memory of VM Instance in MB. 2048MB, 5120MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ent_bw_available": {
				Description: "Enterprise Bandwidth configured.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of provision VPX. Minimum length =  1 Maximum length =  512",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Description: "User Name (except nsroot) to be configured on Citrix ADC Instance. Minimum length =  1 Maximum length =  127",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fips_partition_name": {
				Description: "FIPS Partition Name. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nsvlan_id": {
				Description: "VLAN Id. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_pes": {
				Description: "Total number of PEs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"burst_priority": {
				Description: "Burst Priority of the VM Instance between 1 and 4.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"metrics_collection": {
				Description: "Flag to check if metrics collection is enabled or disabled..",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_managed": {
				Description: "Is Managed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nexthop_v6": {
				Description: "Next Hop IPv6 Address.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "IPv4 Address. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"profile_name": {
				Description: "Device Profile Name that is attached with this provision VPX. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backplane": {
				Description: "Backplane Interface. Minimum length =  1",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"network_interfaces": {
				Description: "Network Interfaces.",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_name": {
							Description: "Port Name.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name_server": {
							Description: "Name Server.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_mgmt_ifc": {
							Description: "Is Mgmt Ifc.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"gateway": {
							Description: "Gateway.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vrid_list_ipv6": {
							Description: "VRID List IPv6.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"parent_id": {
							Description: "Parent ID.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vrid_list_ipv4": {
							Description: "VRID List IPv4.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_member_ifc": {
							Description: "Is Member Ifc.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"mac_address": {
							Description: "Mac Address.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"netmask": {
							Description: "Netmask.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"l2_enabled": {
							Description: "L2 Enabled.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "Id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"interface_name": {
							Description: "Interface Name.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"parent_name": {
							Description: "Parent Name.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vlan_whitelist_array": {
							Description: "Vlan Whitelist Array.",
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Computed:    true,
						},
						"mac_mode": {
							Description: "Mac Mode.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"managed_device_id": {
							Description: "Managed Device Id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Description: "Vlan.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vrid_list_ipv4_array": {
							Description: "VRID List IPv4 Array.",
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Computed:    true,
						},
						"receiveuntagged": {
							Description: "Receiveuntagged.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vrid_list_ipv6_array": {
							Description: "VRID List IPv6 Array.",
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Computed:    true,
						},
						"is_vlan_applied": {
							Description: "Is Vlan Applied.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"vlan_whitelist": {
							Description: "Vlan Whitelist.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"state": {
				Description: "Node State. Minimum length =  1 Maximum length =  32",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_updated_time": {
				Description: "Last Updated Time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_edition": {
				Description: "Edition of instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"customid": {
				Description: "Custom ID.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_grace_time": {
				Description: "Grace for this Citrix Instance..",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"la_mgmt": {
				Description: "Bond consisting of management ports on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_id_0_2": {
				Description: "VLAN id for the management interface 0/2. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_ssl_cores": {
				Description: "Assign number of ssl virtual functions to VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_ssl_cards": {
				Description: "Number of SSL Cards.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"max_burst_throughput": {
				Description: "Maximum burst throughput in Mbps of VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"config_type": {
				Description: "Configuration Type. Values: 0: IPv4, 1: IPv6, 2: Both. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_internal_ip_enabled": {
				Description: "Set as true if VPX is managed by internal network (not required to be set for SDWAN).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cmd_policy": {
				Description: "true if you want to allow shell/sftp/scp access to Citrix ADC Instance administrator. Minimum length =  1 Maximum length =  1024",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"node_id": {
				Description: "Node identification of a device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"throughput_allocation_mode": {
				Description: "Throughput Allocation Mode: 0-Fixed, 1-Burst-able.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_cores": {
				Description: "Number of cores that are assigned to VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pps": {
				Description: "Assign packets per seconds to Citrix ADC Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"plugin_netmask": {
				Description: "Signaling Netmask. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"display_name": {
				Description: "Display Name for this provision VPX. For HA pair it will be A-B, and for Cluster it will be CLIP. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"std_bw_total": {
				Description: "Standard Bandwidth.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nexthop": {
				Description: "Next Hop IP address. Minimum length =  1 Maximum length =  64",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"plt_bw_config": {
				Description: "Platinum Bandwidth configured.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_1_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/2 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_1_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/2 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_1_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/7 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_1_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/2 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv4_1_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/4 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_10_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/6 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_10_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/4 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/5 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_10_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/3 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_1_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/1 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_1_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/5 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_1_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/4 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_10_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/4 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_1_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/3for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_10_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/3 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_10_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/5 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_10_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/4 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_1_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/8 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_10_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/6 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_10_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/5 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"save_config": {
				Description: "Should config be saved first in case instance is rebooted while modify.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_new_crypto": {
				Description: "True if number_of_acu/number_of_scu are used, false if number_of_ssl_cores is used.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_1_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/6 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"profile_username": {
				Description: "User Name specified by the user for this Citrix ADC Instance.. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_10_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/7 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_10_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/8 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_1_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/4 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_1_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/6 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_1_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/4 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_1_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/2 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_1_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/5 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_10_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/6 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_10_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/1 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_1_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/6 for IPV4 VMAC Generation.",
				// list of string
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_1_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/5 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_10_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/8 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv4_1_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/7 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_10_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/2 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"profile_password": {
				Description: "Password specified by the user for this Citrix ADC Instance.. Minimum length =  1 Maximum length =  128",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_1_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/5 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_1_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/7 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/4 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_1_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/7 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_1_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/1 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_1_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/7 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_10_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/1 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_10_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/8 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"l2_enabled": {
				Description: "L2mode status of VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/7 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/1 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_1_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/2 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"reboot_vm_on_cpu_change": {
				Description: "Reboot VMs on CPU change during resource allocation.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_1_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/6 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_1_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/3for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_10_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/7 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_1_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/1 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_1_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/1 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_10_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/7 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_1_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/8 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/3 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_10_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/5 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/2 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/6 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv4_1_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 1/8 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"receiveuntagged_1_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/8 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/4 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_1_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/3 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_10_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 10/3 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_10_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/6 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_1_6": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/6 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_10_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/3 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_1_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 1/5 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"if_1_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/3 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nsvlan_interfaces": {
				Description: "VLAN Interfaces. Minimum length =  1 Maximum length =  50",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv4_10_5": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/5 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_1_4": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/4 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/8 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_1_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/1 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_7": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/7 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vlan_10_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVLAN for Network 10/8 on VM Instance. Maximum value =  ",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"receiveuntagged_1_3": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtReceive Untagged Packets on 1/3 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"crypto_change_requires_reboot": {
				Description: "True if the current changes made by user requires a reboot of the VM else false.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/2 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_1_8": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 1/8 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv4_10_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/1 for IPV4 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"vrid_list_ipv6_10_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/2 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"if_10_2": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtNetwork 10/2 on VM Instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vrid_list_ipv6_10_1": {
				Description: "This property is deprecated by network_interfaces&ltbr&gtVRID List for Interface 10/1 for IPV6 VMAC Generation.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func getProvisionVpxPayload(d *schema.ResourceData) interface{} {
	data := make(map[string]interface{})

	data["ip_address"] = d.Get("ip_address").(string)

	if v, ok := d.GetOk("std_bw_config"); ok {
		data["std_bw_config"] = v.(string)
	}

	if v, ok := d.GetOk("ns_ip_address"); ok {
		data["ns_ip_address"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		data["password"] = v.(string)
	}

	if v, ok := d.GetOk("gateway_ipv6"); ok {
		data["gateway_ipv6"] = v.(string)
	}

	if v, ok := d.GetOk("throughput_limit"); ok {
		data["throughput_limit"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_id_0_1"); ok {
		data["vlan_id_0_1"] = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		data["name"] = v.(string)
	}

	if v, ok := d.GetOk("mastools_version"); ok {
		data["mastools_version"] = v.(string)
	}

	if v, ok := d.GetOk("plugin_ip_address"); ok {
		data["plugin_ip_address"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_type"); ok {
		data["vlan_type"] = v.(string)
	}

	if v, ok := d.GetOk("ent_bw_total"); ok {
		data["ent_bw_total"] = v.(string)
	}

	if v, ok := d.GetOk("vcpu_config"); ok {
		data["vcpu_config"] = v.(string)
	}

	if v, ok := d.GetOk("nsvlan_tagged"); ok {
		data["nsvlan_tagged"] = v.(string)
	}

	if v, ok := d.GetOk("netmask"); ok {
		data["netmask"] = v.(string)
	}

	if v, ok := d.GetOk("ent_bw_config"); ok {
		data["ent_bw_config"] = v.(string)
	}

	if v, ok := d.GetOk("datacenter_id"); ok {
		data["datacenter_id"] = v.(string)
	}

	if v, ok := d.GetOk("instance_mode"); ok {
		data["instance_mode"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_ssl_cores_up"); ok {
		data["number_of_ssl_cores_up"] = v.(string)
	}

	if v, ok := d.GetOk("std_bw_available"); ok {
		data["std_bw_available"] = v.(string)
	}

	if v, ok := d.GetOk("internal_ip_address"); ok {
		data["internal_ip_address"] = v.(string)
	}

	if v, ok := d.GetOk("if_0_1"); ok {
		data["if_0_1"] = v.(string)
	}

	if v, ok := d.GetOk("plt_bw_total"); ok {
		data["plt_bw_total"] = v.(string)
	}

	if v, ok := d.GetOk("host_ip_address"); ok {
		data["host_ip_address"] = v.(string)
	}

	// if v, ok := d.GetOk("id"); ok {
	// 	data["id"] = v.(string)
	// }

	if v, ok := d.GetOk("ipv6_address"); ok {
		data["ipv6_address"] = v.(string)
	}

	if v, ok := d.GetOk("mgmt_ip_address"); ok {
		data["mgmt_ip_address"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_acu"); ok {
		data["number_of_acu"] = v.(string)
	}

	if v, ok := d.GetOk("plt_bw_available"); ok {
		data["plt_bw_available"] = v.(string)
	}

	if v, ok := d.GetOk("is_clip"); ok {
		data["is_clip"] = v.(string)
	}

	if v, ok := d.GetOk("device_family"); ok {
		data["device_family"] = v.(string)
	}

	if v, ok := d.GetOk("type"); ok {
		data["type"] = v.(string)
	}

	if v, ok := d.GetOk("throughput"); ok {
		data["throughput"] = v.(string)
	}

	if v, ok := d.GetOk("template_name"); ok {
		data["template_name"] = v.(string)
	}

	if v, ok := d.GetOk("gateway"); ok {
		data["gateway"] = v.(string)
	}

	if v, ok := d.GetOk("iscco"); ok {
		data["iscco"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_scu"); ok {
		data["number_of_scu"] = v.(string)
	}

	if v, ok := d.GetOk("license"); ok {
		data["license"] = v.(string)
	}

	if v, ok := d.GetOk("domain_name"); ok {
		data["domain_name"] = v.(string)
	}

	if v, ok := d.GetOk("if_0_2"); ok {
		data["if_0_2"] = v.(string)
	}

	if v, ok := d.GetOk("image_name"); ok {
		data["image_name"] = v.(string)
	}

	if v, ok := d.GetOk("hostname"); ok {
		data["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("vm_memory_total"); ok {
		data["vm_memory_total"] = v.(string)
	}

	if v, ok := d.GetOk("ent_bw_available"); ok {
		data["ent_bw_available"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		data["description"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		data["username"] = v.(string)
	}

	if v, ok := d.GetOk("fips_partition_name"); ok {
		data["fips_partition_name"] = v.(string)
	}

	if v, ok := d.GetOk("nsvlan_id"); ok {
		data["nsvlan_id"] = v.(string)
	}

	if v, ok := d.GetOk("num_pes"); ok {
		data["num_pes"] = v.(string)
	}

	if v, ok := d.GetOk("burst_priority"); ok {
		data["burst_priority"] = v.(string)
	}

	if v, ok := d.GetOk("metrics_collection"); ok {
		data["metrics_collection"] = v.(string)
	}

	if v, ok := d.GetOk("is_managed"); ok {
		data["is_managed"] = v.(string)
	}

	if v, ok := d.GetOk("nexthop_v6"); ok {
		data["nexthop_v6"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		data["ipv4_address"] = v.(string)
	}

	if v, ok := d.GetOk("profile_name"); ok {
		data["profile_name"] = v.(string)
	}

	if v, ok := d.GetOk("backplane"); ok {
		data["backplane"] = v.(string)
	}

	if v, ok := d.GetOk("network_interfaces"); ok {
		data["network_interfaces"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("state"); ok {
		data["state"] = v.(string)
	}

	if v, ok := d.GetOk("last_updated_time"); ok {
		data["last_updated_time"] = v.(string)
	}

	if v, ok := d.GetOk("license_edition"); ok {
		data["license_edition"] = v.(string)
	}

	if v, ok := d.GetOk("customid"); ok {
		data["customid"] = v.(string)
	}

	if v, ok := d.GetOk("license_grace_time"); ok {
		data["license_grace_time"] = v.(string)
	}

	if v, ok := d.GetOk("la_mgmt"); ok {
		data["la_mgmt"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_id_0_2"); ok {
		data["vlan_id_0_2"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_ssl_cores"); ok {
		data["number_of_ssl_cores"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_ssl_cards"); ok {
		data["number_of_ssl_cards"] = v.(string)
	}

	if v, ok := d.GetOk("max_burst_throughput"); ok {
		data["max_burst_throughput"] = v.(string)
	}

	if v, ok := d.GetOk("config_type"); ok {
		data["config_type"] = v.(string)
	}

	if v, ok := d.GetOk("if_internal_ip_enabled"); ok {
		data["if_internal_ip_enabled"] = v.(string)
	}

	if v, ok := d.GetOk("cmd_policy"); ok {
		data["cmd_policy"] = v.(string)
	}

	if v, ok := d.GetOk("node_id"); ok {
		data["node_id"] = v.(string)
	}

	if v, ok := d.GetOk("throughput_allocation_mode"); ok {
		data["throughput_allocation_mode"] = v.(string)
	}

	if v, ok := d.GetOk("number_of_cores"); ok {
		data["number_of_cores"] = v.(string)
	}

	if v, ok := d.GetOk("pps"); ok {
		data["pps"] = v.(string)
	}

	if v, ok := d.GetOk("plugin_netmask"); ok {
		data["plugin_netmask"] = v.(string)
	}

	if v, ok := d.GetOk("display_name"); ok {
		data["display_name"] = v.(string)
	}

	if v, ok := d.GetOk("std_bw_total"); ok {
		data["std_bw_total"] = v.(string)
	}

	if v, ok := d.GetOk("nexthop"); ok {
		data["nexthop"] = v.(string)
	}

	if v, ok := d.GetOk("plt_bw_config"); ok {
		data["plt_bw_config"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_1_2"); ok {
		data["vlan_1_2"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_2"); ok {
		data["vrid_list_ipv4_1_2"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_1_7"); ok {
		data["vlan_1_7"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_2"); ok {
		data["vrid_list_ipv6_1_2"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_4"); ok {
		data["vrid_list_ipv4_1_4"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_6"); ok {
		data["vrid_list_ipv6_10_6"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_10_4"); ok {
		data["if_10_4"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_5"); ok {
		data["vlan_10_5"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_3"); ok {
		data["vrid_list_ipv6_10_3"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_1"); ok {
		data["vrid_list_ipv6_1_1"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_1_5"); ok {
		data["if_1_5"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_1_4"); ok {
		data["receiveuntagged_1_4"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_4"); ok {
		data["vrid_list_ipv6_10_4"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_3"); ok {
		data["vrid_list_ipv6_1_3"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_10_3"); ok {
		data["if_10_3"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_5"); ok {
		data["vrid_list_ipv6_10_5"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_10_4"); ok {
		data["receiveuntagged_10_4"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_8"); ok {
		data["vrid_list_ipv6_1_8"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_10_6"); ok {
		data["vlan_10_6"] = v.(string)
	}

	if v, ok := d.GetOk("if_10_5"); ok {
		data["if_10_5"] = v.(string)
	}

	if v, ok := d.GetOk("save_config"); ok {
		data["save_config"] = v.(string)
	}

	if v, ok := d.GetOk("is_new_crypto"); ok {
		data["is_new_crypto"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_6"); ok {
		data["vrid_list_ipv6_1_6"] = v.([]interface{})
	}

	if v, ok := d.GetOk("profile_username"); ok {
		data["profile_username"] = v.(string)
	}

	if v, ok := d.GetOk("if_10_7"); ok {
		data["if_10_7"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_10_8"); ok {
		data["receiveuntagged_10_8"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_4"); ok {
		data["vrid_list_ipv6_1_4"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_1_6"); ok {
		data["receiveuntagged_1_6"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_1_4"); ok {
		data["vlan_1_4"] = v.(string)
	}

	if v, ok := d.GetOk("if_1_2"); ok {
		data["if_1_2"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_1_5"); ok {
		data["receiveuntagged_1_5"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_10_6"); ok {
		data["receiveuntagged_10_6"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_10_1"); ok {
		data["receiveuntagged_10_1"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_6"); ok {
		data["vrid_list_ipv4_1_6"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_5"); ok {
		data["vrid_list_ipv6_1_5"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_8"); ok {
		data["vrid_list_ipv6_10_8"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_7"); ok {
		data["vrid_list_ipv4_1_7"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_10_2"); ok {
		data["receiveuntagged_10_2"] = v.(string)
	}

	if v, ok := d.GetOk("profile_password"); ok {
		data["profile_password"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_5"); ok {
		data["vrid_list_ipv4_1_5"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_1_7"); ok {
		data["receiveuntagged_1_7"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_4"); ok {
		data["vlan_10_4"] = v.(string)
	}

	if v, ok := d.GetOk("if_1_7"); ok {
		data["if_1_7"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_1_1"); ok {
		data["receiveuntagged_1_1"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_1_7"); ok {
		data["vrid_list_ipv6_1_7"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_10_1"); ok {
		data["if_10_1"] = v.(string)
	}

	if v, ok := d.GetOk("if_10_8"); ok {
		data["if_10_8"] = v.(string)
	}

	if v, ok := d.GetOk("l2_enabled"); ok {
		data["l2_enabled"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_7"); ok {
		data["vlan_10_7"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_1"); ok {
		data["vlan_10_1"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_1_2"); ok {
		data["receiveuntagged_1_2"] = v.(string)
	}

	if v, ok := d.GetOk("reboot_vm_on_cpu_change"); ok {
		data["reboot_vm_on_cpu_change"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_1_6"); ok {
		data["vlan_1_6"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_3"); ok {
		data["vrid_list_ipv4_1_3"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_10_7"); ok {
		data["receiveuntagged_10_7"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_1"); ok {
		data["vrid_list_ipv4_1_1"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_1_1"); ok {
		data["vlan_1_1"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_7"); ok {
		data["vrid_list_ipv6_10_7"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_1_8"); ok {
		data["vlan_1_8"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_3"); ok {
		data["vrid_list_ipv4_10_3"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_10_5"); ok {
		data["receiveuntagged_10_5"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_2"); ok {
		data["vlan_10_2"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_6"); ok {
		data["vrid_list_ipv4_10_6"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv4_1_8"); ok {
		data["vrid_list_ipv4_1_8"] = v.([]interface{})
	}

	if v, ok := d.GetOk("receiveuntagged_1_8"); ok {
		data["receiveuntagged_1_8"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_4"); ok {
		data["vrid_list_ipv4_10_4"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_1_3"); ok {
		data["vlan_1_3"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_10_3"); ok {
		data["receiveuntagged_10_3"] = v.(string)
	}

	if v, ok := d.GetOk("if_10_6"); ok {
		data["if_10_6"] = v.(string)
	}

	if v, ok := d.GetOk("if_1_6"); ok {
		data["if_1_6"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_10_3"); ok {
		data["vlan_10_3"] = v.(string)
	}

	if v, ok := d.GetOk("vlan_1_5"); ok {
		data["vlan_1_5"] = v.(string)
	}

	if v, ok := d.GetOk("if_1_3"); ok {
		data["if_1_3"] = v.(string)
	}

	if v, ok := d.GetOk("nsvlan_interfaces"); ok {
		data["nsvlan_interfaces"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_5"); ok {
		data["vrid_list_ipv4_10_5"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_1_4"); ok {
		data["if_1_4"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_8"); ok {
		data["vrid_list_ipv4_10_8"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_1_1"); ok {
		data["if_1_1"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_7"); ok {
		data["vrid_list_ipv4_10_7"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vlan_10_8"); ok {
		data["vlan_10_8"] = v.(string)
	}

	if v, ok := d.GetOk("receiveuntagged_1_3"); ok {
		data["receiveuntagged_1_3"] = v.(string)
	}

	if v, ok := d.GetOk("crypto_change_requires_reboot"); ok {
		data["crypto_change_requires_reboot"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_2"); ok {
		data["vrid_list_ipv4_10_2"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_1_8"); ok {
		data["if_1_8"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv4_10_1"); ok {
		data["vrid_list_ipv4_10_1"] = v.([]interface{})
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_2"); ok {
		data["vrid_list_ipv6_10_2"] = v.([]interface{})
	}

	if v, ok := d.GetOk("if_10_2"); ok {
		data["if_10_2"] = v.(string)
	}

	if v, ok := d.GetOk("vrid_list_ipv6_10_1"); ok {
		data["vrid_list_ipv6_10_1"] = v.([]interface{})
	}

	return data
}

func resourceProvisionVpxCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In resourceProvisionVpxCreate")

	c := m.(*service.NitroClient)

	endpoint := "ns"

	returnData, err := c.AddResource(endpoint, getProvisionVpxPayload(d))

	if err != nil {
		return diag.Errorf("unable to createProvision VPX: %s", err.Error())
	}

	vpxID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// wait for VPX instance_state to be Up
	log.Printf("Wait for VPX instance_state to be Up")

	for {
		time.Sleep(5 * time.Second)

		returnData, err := c.GetResource(endpoint, vpxID)
		if err != nil {
			return diag.Errorf("unable to get VPX: %s", err.Error())
		}
		instanceState := returnData[endpoint].([]interface{})[0].(map[string]interface{})["instance_state"].(string)
		if instanceState == "Up" {
			break
		}
		log.Printf("VPX instance_state is %s", instanceState)
	}

	d.SetId(vpxID)
	return resourceProvisionVpxRead(ctx, d, m)
}

func getProvisionVpxID(c *service.NitroClient, ipAddress string) (string, error) {
	endpoint := "ns"
	returnData, err := c.GetAllResource(endpoint)
	if err != nil {
		return "", err
	}

	for _, v := range returnData[endpoint].([]interface{}) {
		if v.(map[string]interface{})["ip_address"].(string) == ipAddress {
			return v.(map[string]interface{})["id"].(string), nil
		}
	}
	return "", errors.New("Failed to find provision VPX resource ID with IP: " + ipAddress)
}

func resourceProvisionVpxRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In resourceProvisionVpxRead")
	var diags diag.Diagnostics
	c := m.(*service.NitroClient)

	resourceID := d.Id()
	endpoint := "ns"

	returnData, err := c.GetResource(endpoint, resourceID)
	if err != nil {
		return diag.FromErr(err)
	}

	getResponseData := returnData[endpoint].([]interface{})[0].(map[string]interface{})

	d.Set("std_bw_config", getResponseData["std_bw_config"].(string))
	d.Set("ns_ip_address", getResponseData["ns_ip_address"].(string))
	d.Set("gateway_ipv6", getResponseData["gateway_ipv6"].(string))
	d.Set("throughput_limit", getResponseData["throughput_limit"].(string))
	d.Set("vlan_id_0_1", getResponseData["vlan_id_0_1"].(string))
	d.Set("name", getResponseData["name"].(string))
	d.Set("mastools_version", getResponseData["mastools_version"].(string))
	d.Set("plugin_ip_address", getResponseData["plugin_ip_address"].(string))
	d.Set("vlan_type", getResponseData["vlan_type"].(string))
	d.Set("ent_bw_total", getResponseData["ent_bw_total"].(string))
	d.Set("vcpu_config", getResponseData["vcpu_config"].(string))
	d.Set("nsvlan_tagged", getResponseData["nsvlan_tagged"].(string))
	d.Set("netmask", getResponseData["netmask"].(string))
	d.Set("ent_bw_config", getResponseData["ent_bw_config"].(string))
	d.Set("datacenter_id", getResponseData["datacenter_id"].(string))
	d.Set("instance_mode", getResponseData["instance_mode"].(string))
	d.Set("number_of_ssl_cores_up", getResponseData["number_of_ssl_cores_up"].(string))
	d.Set("std_bw_available", getResponseData["std_bw_available"].(string))
	d.Set("internal_ip_address", getResponseData["internal_ip_address"].(string))
	d.Set("if_0_1", getResponseData["if_0_1"].(string))
	d.Set("plt_bw_total", getResponseData["plt_bw_total"].(string))
	d.Set("host_ip_address", getResponseData["host_ip_address"].(string))
	d.Set("vpx_id", getResponseData["id"].(string))
	d.Set("ipv6_address", getResponseData["ipv6_address"].(string))
	d.Set("mgmt_ip_address", getResponseData["mgmt_ip_address"].(string))
	d.Set("number_of_acu", getResponseData["number_of_acu"].(string))
	d.Set("plt_bw_available", getResponseData["plt_bw_available"].(string))
	d.Set("is_clip", getResponseData["is_clip"].(string))
	d.Set("device_family", getResponseData["device_family"].(string))
	d.Set("type", getResponseData["type"].(string))
	d.Set("throughput", getResponseData["throughput"].(string))
	d.Set("template_name", getResponseData["template_name"].(string))
	d.Set("gateway", getResponseData["gateway"].(string))
	d.Set("iscco", getResponseData["iscco"].(string))
	d.Set("number_of_scu", getResponseData["number_of_scu"].(string))
	d.Set("license", getResponseData["license"].(string))
	d.Set("domain_name", getResponseData["domain_name"].(string))
	d.Set("if_0_2", getResponseData["if_0_2"].(string))
	// d.Set("image_name", getResponseData["image_name"].(string)) // FIXME: API Problem. image_name is empty after Update operataion
	d.Set("hostname", getResponseData["hostname"].(string))
	d.Set("vm_memory_total", getResponseData["vm_memory_total"].(string))
	d.Set("ent_bw_available", getResponseData["ent_bw_available"].(string))
	d.Set("description", getResponseData["description"].(string))
	d.Set("username", getResponseData["username"].(string))
	d.Set("fips_partition_name", getResponseData["fips_partition_name"].(string))
	d.Set("nsvlan_id", getResponseData["nsvlan_id"].(string))
	d.Set("num_pes", getResponseData["num_pes"].(string))
	d.Set("burst_priority", getResponseData["burst_priority"].(string))
	d.Set("metrics_collection", getResponseData["metrics_collection"].(string))
	d.Set("is_managed", getResponseData["is_managed"].(string))
	d.Set("nexthop_v6", getResponseData["nexthop_v6"].(string))
	d.Set("ipv4_address", getResponseData["ipv4_address"].(string))
	d.Set("profile_name", getResponseData["profile_name"].(string))
	d.Set("backplane", getResponseData["backplane"].(string))
	if err := d.Set("network_interfaces", parseNetworkInterface(d, getResponseData["network_interfaces"].([]interface{}))); err != nil {
		return diag.FromErr(err)
	}
	d.Set("state", getResponseData["state"].(string))
	d.Set("last_updated_time", getResponseData["last_updated_time"].(string))
	d.Set("license_edition", getResponseData["license_edition"].(string))
	d.Set("customid", getResponseData["customid"].(string))
	d.Set("license_grace_time", getResponseData["license_grace_time"].(string))
	d.Set("la_mgmt", getResponseData["la_mgmt"].(string))
	d.Set("vlan_id_0_2", getResponseData["vlan_id_0_2"].(string))
	d.Set("number_of_ssl_cores", getResponseData["number_of_ssl_cores"].(string))
	d.Set("number_of_ssl_cards", getResponseData["number_of_ssl_cards"].(string))
	d.Set("max_burst_throughput", getResponseData["max_burst_throughput"].(string))
	d.Set("config_type", getResponseData["config_type"].(string))
	d.Set("if_internal_ip_enabled", getResponseData["if_internal_ip_enabled"].(string))
	d.Set("cmd_policy", getResponseData["cmd_policy"].(string))
	d.Set("node_id", getResponseData["node_id"].(string))
	d.Set("ip_address", getResponseData["ip_address"].(string))
	d.Set("throughput_allocation_mode", getResponseData["throughput_allocation_mode"].(string))
	d.Set("number_of_cores", getResponseData["number_of_cores"].(string))
	d.Set("pps", getResponseData["pps"].(string))
	d.Set("plugin_netmask", getResponseData["plugin_netmask"].(string))
	d.Set("display_name", getResponseData["display_name"].(string))
	d.Set("std_bw_total", getResponseData["std_bw_total"].(string))
	d.Set("nexthop", getResponseData["nexthop"].(string))
	d.Set("plt_bw_config", getResponseData["plt_bw_config"].(string))
	d.Set("vlan_1_2", getResponseData["vlan_1_2"].(string))
	d.Set("vrid_list_ipv4_1_2", getResponseData["vrid_list_ipv4_1_2"].([]interface{}))
	d.Set("vlan_1_7", getResponseData["vlan_1_7"].(string))
	d.Set("vrid_list_ipv6_1_2", getResponseData["vrid_list_ipv6_1_2"].([]interface{}))
	d.Set("vrid_list_ipv4_1_4", getResponseData["vrid_list_ipv4_1_4"].([]interface{}))
	d.Set("vrid_list_ipv6_10_6", getResponseData["vrid_list_ipv6_10_6"].([]interface{}))
	d.Set("if_10_4", getResponseData["if_10_4"].(string))
	d.Set("vlan_10_5", getResponseData["vlan_10_5"].(string))
	d.Set("vrid_list_ipv6_10_3", getResponseData["vrid_list_ipv6_10_3"].([]interface{}))
	d.Set("vrid_list_ipv6_1_1", getResponseData["vrid_list_ipv6_1_1"].([]interface{}))
	d.Set("if_1_5", getResponseData["if_1_5"].(string))
	d.Set("receiveuntagged_1_4", getResponseData["receiveuntagged_1_4"].(string))
	d.Set("vrid_list_ipv6_10_4", getResponseData["vrid_list_ipv6_10_4"].([]interface{}))
	d.Set("vrid_list_ipv6_1_3", getResponseData["vrid_list_ipv6_1_3"].([]interface{}))
	d.Set("if_10_3", getResponseData["if_10_3"].(string))
	d.Set("vrid_list_ipv6_10_5", getResponseData["vrid_list_ipv6_10_5"].([]interface{}))
	d.Set("receiveuntagged_10_4", getResponseData["receiveuntagged_10_4"].(string))
	d.Set("vrid_list_ipv6_1_8", getResponseData["vrid_list_ipv6_1_8"].([]interface{}))
	d.Set("vlan_10_6", getResponseData["vlan_10_6"].(string))
	d.Set("if_10_5", getResponseData["if_10_5"].(string))
	d.Set("save_config", getResponseData["save_config"].(string))
	d.Set("is_new_crypto", getResponseData["is_new_crypto"].(string))
	d.Set("vrid_list_ipv6_1_6", getResponseData["vrid_list_ipv6_1_6"].([]interface{}))
	d.Set("profile_username", getResponseData["profile_username"].(string))
	d.Set("if_10_7", getResponseData["if_10_7"].(string))
	d.Set("receiveuntagged_10_8", getResponseData["receiveuntagged_10_8"].(string))
	d.Set("vrid_list_ipv6_1_4", getResponseData["vrid_list_ipv6_1_4"].([]interface{}))
	d.Set("receiveuntagged_1_6", getResponseData["receiveuntagged_1_6"].(string))
	d.Set("vlan_1_4", getResponseData["vlan_1_4"].(string))
	d.Set("if_1_2", getResponseData["if_1_2"].(string))
	d.Set("receiveuntagged_1_5", getResponseData["receiveuntagged_1_5"].(string))
	d.Set("receiveuntagged_10_6", getResponseData["receiveuntagged_10_6"].(string))
	d.Set("receiveuntagged_10_1", getResponseData["receiveuntagged_10_1"].(string))
	d.Set("vrid_list_ipv4_1_6", getResponseData["vrid_list_ipv4_1_6"].([]interface{}))
	d.Set("vrid_list_ipv6_1_5", getResponseData["vrid_list_ipv6_1_5"].([]interface{}))
	d.Set("vrid_list_ipv6_10_8", getResponseData["vrid_list_ipv6_10_8"].([]interface{}))
	d.Set("vrid_list_ipv4_1_7", getResponseData["vrid_list_ipv4_1_7"].([]interface{}))
	d.Set("receiveuntagged_10_2", getResponseData["receiveuntagged_10_2"].(string))
	d.Set("vrid_list_ipv4_1_5", getResponseData["vrid_list_ipv4_1_5"].([]interface{}))
	d.Set("receiveuntagged_1_7", getResponseData["receiveuntagged_1_7"].(string))
	d.Set("vlan_10_4", getResponseData["vlan_10_4"].(string))
	d.Set("if_1_7", getResponseData["if_1_7"].(string))
	d.Set("receiveuntagged_1_1", getResponseData["receiveuntagged_1_1"].(string))
	d.Set("vrid_list_ipv6_1_7", getResponseData["vrid_list_ipv6_1_7"].([]interface{}))
	d.Set("if_10_1", getResponseData["if_10_1"].(string))
	d.Set("if_10_8", getResponseData["if_10_8"].(string))
	d.Set("l2_enabled", getResponseData["l2_enabled"].(string))
	d.Set("vlan_10_7", getResponseData["vlan_10_7"].(string))
	d.Set("vlan_10_1", getResponseData["vlan_10_1"].(string))
	d.Set("receiveuntagged_1_2", getResponseData["receiveuntagged_1_2"].(string))
	d.Set("reboot_vm_on_cpu_change", getResponseData["reboot_vm_on_cpu_change"].(string))
	d.Set("vlan_1_6", getResponseData["vlan_1_6"].(string))
	d.Set("vrid_list_ipv4_1_3", getResponseData["vrid_list_ipv4_1_3"].([]interface{}))
	d.Set("receiveuntagged_10_7", getResponseData["receiveuntagged_10_7"].(string))
	d.Set("vrid_list_ipv4_1_1", getResponseData["vrid_list_ipv4_1_1"].([]interface{}))
	d.Set("vlan_1_1", getResponseData["vlan_1_1"].(string))
	d.Set("vrid_list_ipv6_10_7", getResponseData["vrid_list_ipv6_10_7"].([]interface{}))
	d.Set("vlan_1_8", getResponseData["vlan_1_8"].(string))
	d.Set("vrid_list_ipv4_10_3", getResponseData["vrid_list_ipv4_10_3"].([]interface{}))
	d.Set("receiveuntagged_10_5", getResponseData["receiveuntagged_10_5"].(string))
	d.Set("vlan_10_2", getResponseData["vlan_10_2"].(string))
	d.Set("vrid_list_ipv4_10_6", getResponseData["vrid_list_ipv4_10_6"].([]interface{}))
	d.Set("vrid_list_ipv4_1_8", getResponseData["vrid_list_ipv4_1_8"].([]interface{}))
	d.Set("receiveuntagged_1_8", getResponseData["receiveuntagged_1_8"].(string))
	d.Set("vrid_list_ipv4_10_4", getResponseData["vrid_list_ipv4_10_4"].([]interface{}))
	d.Set("vlan_1_3", getResponseData["vlan_1_3"].(string))
	d.Set("receiveuntagged_10_3", getResponseData["receiveuntagged_10_3"].(string))
	d.Set("if_10_6", getResponseData["if_10_6"].(string))
	d.Set("if_1_6", getResponseData["if_1_6"].(string))
	d.Set("vlan_10_3", getResponseData["vlan_10_3"].(string))
	d.Set("vlan_1_5", getResponseData["vlan_1_5"].(string))
	d.Set("if_1_3", getResponseData["if_1_3"].(string))
	d.Set("nsvlan_interfaces", getResponseData["nsvlan_interfaces"].([]interface{}))
	d.Set("vrid_list_ipv4_10_5", getResponseData["vrid_list_ipv4_10_5"].([]interface{}))
	d.Set("if_1_4", getResponseData["if_1_4"].(string))
	d.Set("vrid_list_ipv4_10_8", getResponseData["vrid_list_ipv4_10_8"].([]interface{}))
	d.Set("if_1_1", getResponseData["if_1_1"].(string))
	d.Set("vrid_list_ipv4_10_7", getResponseData["vrid_list_ipv4_10_7"].([]interface{}))
	d.Set("vlan_10_8", getResponseData["vlan_10_8"].(string))
	d.Set("receiveuntagged_1_3", getResponseData["receiveuntagged_1_3"].(string))
	d.Set("crypto_change_requires_reboot", getResponseData["crypto_change_requires_reboot"].(string))
	d.Set("vrid_list_ipv4_10_2", getResponseData["vrid_list_ipv4_10_2"].([]interface{}))
	d.Set("if_1_8", getResponseData["if_1_8"].(string))
	d.Set("vrid_list_ipv4_10_1", getResponseData["vrid_list_ipv4_10_1"].([]interface{}))
	d.Set("vrid_list_ipv6_10_2", getResponseData["vrid_list_ipv6_10_2"].([]interface{}))
	d.Set("if_10_2", getResponseData["if_10_2"].(string))
	d.Set("vrid_list_ipv6_10_1", getResponseData["vrid_list_ipv6_10_1"].([]interface{}))

	return diags
}

func parseNetworkInterface(d *schema.ResourceData, nif []interface{}) []map[string]interface{} {
	var nifSchemaAttributes = []string{
		"port_name",
		"name_server",
		"is_mgmt_ifc",
		"gateway",
		"vrid_list_ipv6",
		"parent_id",
		"vrid_list_ipv4",
		"is_member_ifc",
		"mac_address",
		"netmask",
		"l2_enabled",
		"id",
		"interface_name",
		"parent_name",
		"vlan_whitelist_array",
		"mac_mode",
		"managed_device_id",
		"vlan",
		"vrid_list_ipv4_array",
		"receiveuntagged",
		"vrid_list_ipv6_array",
		"is_vlan_applied",
		"vlan_whitelist",
	}
	var nifs []map[string]interface{}

	if v, ok := d.GetOk("network_interfaces"); ok {
		inputNifs := v.(*schema.Set).List()

		// get the portnames of all the inputNifs
		var inputNifPortNames []string
		for _, inputNif := range inputNifs {
			inputNif := inputNif.(map[string]interface{})
			inputNifPortNames = append(inputNifPortNames, inputNif["port_name"].(string))
		}

		for _, nif := range nif {
			nifMap := nif.(map[string]interface{})
			if len(nifMap) == 0 {
				continue
			}
			if !service.Contains(inputNifPortNames, nifMap["port_name"].(string)) {
				continue
			}
			// iterate through nifMap and only keep the keys that are in the nifSchemaAttributes
			var nifMap2 map[string]interface{}
			nifMap2 = make(map[string]interface{})
			for k, v := range nifMap {
				if service.Contains(nifSchemaAttributes, k) {
					nifMap2[k] = v
				}
			}
			nifs = append(nifs, nifMap2)
		}
	}
	return nifs
}
func resourceProvisionVpxUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In resourceProvisionVpxUpdate")
	c := m.(*service.NitroClient)

	resourceID := d.Id()
	endpoint := "ns"

	_, err := c.UpdateResource(endpoint, getProvisionVpxPayload(d), resourceID)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProvisionVpxRead(ctx, d, m)
}

func resourceProvisionVpxDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In resourceProvisionVpxDelete")
	var diags diag.Diagnostics

	c := m.(*service.NitroClient)

	endpoint := "ns"
	resourceID := d.Id()

	_, err := c.DeleteResource(endpoint, resourceID)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
