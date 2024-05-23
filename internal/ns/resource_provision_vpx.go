package ns

import (
	"context"
	"fmt"

	"strconv"
	"time"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &provisionVpxResource{}
	_ resource.ResourceWithConfigure = &provisionVpxResource{}
)

// provisionVpxResource defines the resource implementation.
type provisionVpxResource struct {
	client *service.NitroClient
}

func ProvisionVpxResource() resource.Resource {
	return &provisionVpxResource{}
}

// provisionVpxResourceModel describes the resource data model.
type provisionVpxResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	IPAddress                  types.String `tfsdk:"ip_address"`
	StdBwConfig                types.String `tfsdk:"std_bw_config"`
	NsIPAddress                types.String `tfsdk:"ns_ip_address"`
	Password                   types.String `tfsdk:"password"`
	GatewayIPv6                types.String `tfsdk:"gateway_ipv6"`
	ThroughputLimit            types.String `tfsdk:"throughput_limit"`
	VlanID01                   types.String `tfsdk:"vlan_id_0_1"`
	Name                       types.String `tfsdk:"name"`
	MastoolsVersion            types.String `tfsdk:"mastools_version"`
	PluginIPAddress            types.String `tfsdk:"plugin_ip_address"`
	VlanType                   types.String `tfsdk:"vlan_type"`
	EntBwTotal                 types.String `tfsdk:"ent_bw_total"`
	VcpuConfig                 types.String `tfsdk:"vcpu_config"`
	NsvlanTagged               types.String `tfsdk:"nsvlan_tagged"`
	Netmask                    types.String `tfsdk:"netmask"`
	EntBwConfig                types.String `tfsdk:"ent_bw_config"`
	DatacenterID               types.String `tfsdk:"datacenter_id"`
	InstanceMode               types.String `tfsdk:"instance_mode"`
	NumberOfSslCoresUp         types.String `tfsdk:"number_of_ssl_cores_up"`
	StdBwAvailable             types.String `tfsdk:"std_bw_available"`
	InternalIPAddress          types.String `tfsdk:"internal_ip_address"`
	If01                       types.String `tfsdk:"if_0_1"`
	PltBwTotal                 types.String `tfsdk:"plt_bw_total"`
	HostIPAddress              types.String `tfsdk:"host_ip_address"`
	VpxID                      types.String `tfsdk:"vpx_id"`
	IPv6Address                types.String `tfsdk:"ipv6_address"`
	MgmtIPAddress              types.String `tfsdk:"mgmt_ip_address"`
	NumberOfAcu                types.String `tfsdk:"number_of_acu"`
	PltBwAvailable             types.String `tfsdk:"plt_bw_available"`
	IsClip                     types.String `tfsdk:"is_clip"`
	DeviceFamily               types.String `tfsdk:"device_family"`
	Type                       types.String `tfsdk:"type"`
	Throughput                 types.String `tfsdk:"throughput"`
	TemplateName               types.String `tfsdk:"template_name"`
	Gateway                    types.String `tfsdk:"gateway"`
	Iscco                      types.String `tfsdk:"iscco"`
	NumberOfScu                types.String `tfsdk:"number_of_scu"`
	License                    types.String `tfsdk:"license"`
	DomainName                 types.String `tfsdk:"domain_name"`
	If02                       types.String `tfsdk:"if_0_2"`
	ImageName                  types.String `tfsdk:"image_name"`
	Hostname                   types.String `tfsdk:"hostname"`
	VmMemoryTotal              types.String `tfsdk:"vm_memory_total"`
	EntBwAvailable             types.String `tfsdk:"ent_bw_available"`
	Description                types.String `tfsdk:"description"`
	Username                   types.String `tfsdk:"username"`
	FipsPartitionName          types.String `tfsdk:"fips_partition_name"`
	NsvlanID                   types.String `tfsdk:"nsvlan_id"`
	NumPes                     types.String `tfsdk:"num_pes"`
	BurstPriority              types.String `tfsdk:"burst_priority"`
	MetricsCollection          types.String `tfsdk:"metrics_collection"`
	IsManaged                  types.String `tfsdk:"is_managed"`
	NexthopV6                  types.String `tfsdk:"nexthop_v6"`
	IPv4Address                types.String `tfsdk:"ipv4_address"`
	ProfileName                types.String `tfsdk:"profile_name"`
	Backplane                  types.String `tfsdk:"backplane"`
	NetworkInterfaces          types.List   `tfsdk:"network_interfaces"`
	State                      types.String `tfsdk:"state"`
	LastUpdatedTime            types.String `tfsdk:"last_updated_time"`
	LicenseEdition             types.String `tfsdk:"license_edition"`
	Customid                   types.String `tfsdk:"customid"`
	LicenseGraceTime           types.String `tfsdk:"license_grace_time"`
	LaMgmt                     types.String `tfsdk:"la_mgmt"`
	VlanID02                   types.String `tfsdk:"vlan_id_0_2"`
	NumberOfSslCores           types.String `tfsdk:"number_of_ssl_cores"`
	NumberOfsslCards           types.String `tfsdk:"number_of_ssl_cards"`
	MaxBurstThroughput         types.String `tfsdk:"max_burst_throughput"`
	ConfigType                 types.String `tfsdk:"config_type"`
	IfInternalIpEnabled        types.String `tfsdk:"if_internal_ip_enabled"`
	CmdPolicy                  types.String `tfsdk:"cmd_policy"`
	NodeId                     types.String `tfsdk:"node_id"`
	ThroughputAllocationMode   types.String `tfsdk:"throughput_allocation_mode"`
	NumberOfCores              types.String `tfsdk:"number_of_cores"`
	Pps                        types.String `tfsdk:"pps"`
	PluginNetmask              types.String `tfsdk:"plugin_netmask"`
	DisplayName                types.String `tfsdk:"display_name"`
	StdBwTotal                 types.String `tfsdk:"std_bw_total"`
	Nexthop                    types.String `tfsdk:"nexthop"`
	PltBwConfig                types.String `tfsdk:"plt_bw_config"`
	SaveConfig                 types.String `tfsdk:"save_config"`
	IsNewCrypto                types.String `tfsdk:"is_new_crypto"`
	ProfileUsername            types.String `tfsdk:"profile_username"`
	ProfilePassword            types.String `tfsdk:"profile_password"`
	L2Enabled                  types.String `tfsdk:"l2_enabled"`
	RebootVmOnCpuChange        types.String `tfsdk:"reboot_vm_on_cpu_change"`
	NsvlanInterfaces           types.List   `tfsdk:"nsvlan_interfaces"`
	CryptoChangeRequiresReboot types.String `tfsdk:"crypto_change_requires_reboot"`
}

type provisionVpxResourceReq struct {
	Id                         string                   `json:"id,omitempty"`
	IPAddress                  string                   `json:"ip_address,omitempty"`
	StdBwConfig                string                   `json:"std_bw_config,omitempty"`
	NsIPAddress                string                   `json:"ns_ip_address,omitempty"`
	Password                   string                   `json:"password,omitempty"`
	GatewayIPv6                string                   `json:"gateway_ipv6,omitempty"`
	ThroughputLimit            string                   `json:"throughput_limit,omitempty"`
	VlanID01                   string                   `json:"vlan_id_0_1,omitempty"`
	Name                       string                   `json:"name,omitempty"`
	MastoolsVersion            string                   `json:"mastools_version,omitempty"`
	PluginIPAddress            string                   `json:"plugin_ip_address,omitempty"`
	VlanType                   string                   `json:"vlan_type,omitempty"`
	EntBwTotal                 string                   `json:"ent_bw_total,omitempty"`
	VcpuConfig                 string                   `json:"vcpu_config,omitempty"`
	NsvlanTagged               string                   `json:"nsvlan_tagged,omitempty"`
	Netmask                    string                   `json:"netmask,omitempty"`
	EntBwConfig                string                   `json:"ent_bw_config,omitempty"`
	DatacenterID               string                   `json:"datacenter_id,omitempty"`
	InstanceMode               string                   `json:"instance_mode,omitempty"`
	NumberOfSslCoresUp         string                   `json:"number_of_ssl_cores_up,omitempty"`
	StdBwAvailable             string                   `json:"std_bw_available,omitempty"`
	InternalIPAddress          string                   `json:"internal_ip_address,omitempty"`
	If01                       string                   `json:"if_0_1,omitempty"`
	PltBwTotal                 string                   `json:"plt_bw_total,omitempty"`
	HostIPAddress              string                   `json:"host_ip_address,omitempty"`
	VpxID                      string                   `json:"vpx_id,omitempty"`
	IPv6Address                string                   `json:"ipv6_address,omitempty"`
	MgmtIPAddress              string                   `json:"mgmt_ip_address,omitempty"`
	NumberOfAcu                string                   `json:"number_of_acu,omitempty"`
	PltBwAvailable             string                   `json:"plt_bw_available,omitempty"`
	IsClip                     string                   `json:"is_clip,omitempty"`
	DeviceFamily               string                   `json:"device_family,omitempty"`
	Type                       string                   `json:"type,omitempty"`
	Throughput                 string                   `json:"throughput,omitempty"`
	TemplateName               string                   `json:"template_name,omitempty"`
	Gateway                    string                   `json:"gateway,omitempty"`
	Iscco                      string                   `json:"iscco,omitempty"`
	NumberOfScu                string                   `json:"number_of_scu,omitempty"`
	License                    string                   `json:"license,omitempty"`
	DomainName                 string                   `json:"domain_name,omitempty"`
	If02                       string                   `json:"if_0_2,omitempty"`
	ImageName                  string                   `json:"image_name,omitempty"`
	Hostname                   string                   `json:"hostname,omitempty"`
	VmMemoryTotal              string                   `json:"vm_memory_total,omitempty"`
	EntBwAvailable             string                   `json:"ent_bw_available,omitempty"`
	Description                string                   `json:"description,omitempty"`
	Username                   string                   `json:"username,omitempty"`
	FipsPartitionName          string                   `json:"fips_partition_name,omitempty"`
	NsvlanID                   string                   `json:"nsvlan_id,omitempty"`
	NumPes                     string                   `json:"num_pes,omitempty"`
	BurstPriority              string                   `json:"burst_priority,omitempty"`
	MetricsCollection          string                   `json:"metrics_collection,omitempty"`
	IsManaged                  string                   `json:"is_managed,omitempty"`
	NexthopV6                  string                   `json:"nexthop_v6,omitempty"`
	IPv4Address                string                   `json:"ipv4_address,omitempty"`
	ProfileName                string                   `json:"profile_name,omitempty"`
	Backplane                  string                   `json:"backplane,omitempty"`
	NetworkInterfaces          []map[string]interface{} `json:"network_interfaces,omitempty"`
	State                      string                   `json:"state,omitempty"`
	LastUpdatedTime            string                   `json:"last_updated_time,omitempty"`
	LicenseEdition             string                   `json:"license_edition,omitempty"`
	Customid                   string                   `json:"customid,omitempty"`
	LicenseGraceTime           string                   `json:"license_grace_time,omitempty"`
	LaMgmt                     string                   `json:"la_mgmt,omitempty"`
	VlanID02                   string                   `json:"vlan_id_0_2,omitempty"`
	NumberOfSslCores           string                   `json:"number_of_ssl_cores,omitempty"`
	NumberOfsslCards           string                   `json:"number_of_ssl_cards,omitempty"`
	MaxBurstThroughput         string                   `json:"max_burst_throughput,omitempty"`
	ConfigType                 string                   `json:"config_type,omitempty"`
	IfInternalIpEnabled        string                   `json:"if_internal_ip_enabled,omitempty"`
	CmdPolicy                  string                   `json:"cmd_policy,omitempty"`
	NodeId                     string                   `json:"node_id,omitempty"`
	ThroughputAllocationMode   string                   `json:"throughput_allocation_mode,omitempty"`
	NumberOfCores              string                   `json:"number_of_cores,omitempty"`
	Pps                        string                   `json:"pps,omitempty"`
	PluginNetmask              string                   `json:"plugin_netmask,omitempty"`
	DisplayName                string                   `json:"display_name,omitempty"`
	StdBwTotal                 string                   `json:"std_bw_total,omitempty"`
	Nexthop                    string                   `json:"nexthop,omitempty"`
	PltBwConfig                string                   `json:"plt_bw_config,omitempty"`
	SaveConfig                 string                   `json:"save_config,omitempty"`
	IsNewCrypto                string                   `json:"is_new_crypto,omitempty"`
	ProfileUsername            string                   `json:"profile_username,omitempty"`
	ProfilePassword            string                   `json:"profile_password,omitempty"`
	L2Enabled                  string                   `json:"l2_enabled,omitempty"`
	RebootVmOnCpuChange        string                   `json:"reboot_vm_on_cpu_change,omitempty"`
	NsvlanInterfaces           []string                 `json:"nsvlan_interfaces,omitempty"`
	CryptoChangeRequiresReboot string                   `json:"crypto_change_requires_reboot,omitempty"`
}

// Metadata returns the resource type name.
func (r *provisionVpxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_provision_vpx"
}

// Configure configures the client resource.
func (r *provisionVpxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *provisionVpxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Configuration for Provision VPX resource",
		Attributes: map[string]schema.Attribute{
			"ip_address": schema.StringAttribute{
				Required: true,
				// We have below code insted of ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP Address.",
			},
			"std_bw_config": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Standard Bandwidth running.",
			},
			"ns_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "NetScaler ADC IP Address for this provision VPX. Minimum length =  1 Maximum length =  128",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Password for specified user on NetScaler ADC Instance. Minimum length =  1 Maximum length =  127",
			},
			"gateway_ipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Gateway IPv6 Address.",
			},
			"throughput_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Throughput Limit in Mbps set for VM Instance.",
			},
			"vlan_id_0_1": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "VLAN id for the management interface 0/1. Maximum value =  ",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Name of provision VPX. Minimum length =  1 Maximum length =  128",
			},
			"mastools_version": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Mastools version if the device is embedded agent.",
			},
			"plugin_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Signaling IP Address. Minimum length =  1 Maximum length =  64",
			},
			"vlan_type": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "VLAN Type, NetScaler ADC or L2 VLAN. Maximum value =  ",
			},
			"ent_bw_total": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Enterprise Bandwidth Total.",
			},
			"vcpu_config": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of vCPU allocated for the device.",
			},
			"nsvlan_tagged": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "NSVLAN Tagged.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Netmask of provision VPX. Minimum length =  1 Maximum length =  64",
			},
			"ent_bw_config": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Enterprise Bandwidth configured.",
			},
			"datacenter_id": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Datacenter Id is system generated key for data center.",
			},
			"instance_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Denotes state- primary,secondary,clip,clusternode.",
			},
			"number_of_ssl_cores_up": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of SSL Cores Up.",
			},
			"std_bw_available": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Standard Bandwidth Available.",
			},
			"internal_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Internal IP Address for this provision VPX. Minimum length =  1 Maximum length =  64",
			},
			"if_0_1": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Network 0/1 on VM Instance.",
			},
			"plt_bw_total": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Total Platinum Bandwidth.",
			},
			"host_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Host IPAddress where VM is provisioned. Minimum length =  1 Maximum length =  64",
			},
			"vpx_id": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Id is system generated key for all the provision VPXs.",
			},
			"ipv6_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "IPv6 Address.",
			},
			"mgmt_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Management IP Address for thisProvision VPX. Minimum length =  1 Maximum length =  64",
			},
			"number_of_acu": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of ACU.",
			},
			"plt_bw_available": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Platinum Bandwidth Available.",
			},
			"is_clip": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Is Clip.",
			},
			"device_family": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Device Family. Minimum length =  1 Maximum length =  64",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Type of VPX. Minimum length =  1 Maximum length =  64",
			},
			"throughput": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Throughput.",
			},
			"template_name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Template Name.",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Default Gateway of provision VPX. Minimum length =  1 Maximum length =  64",
			},
			"iscco": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Is CCO.",
			},
			"number_of_scu": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of SCU.",
			},
			"license": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "License.",
			},
			"domain_name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Domain name of VM Device. Minimum length =  1 Maximum length =  128",
			},
			"if_0_2": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Network 0/2 on VM Instance.",
			},
			"image_name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Image Name.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Hostname.",
			},
			"vm_memory_total": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Total Memory of VM Instance.",
			},
			"ent_bw_available": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Enterprise Bandwidth Available.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Description of provision VPX. Minimum length =  1 Maximum length =  256",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Username for specified user on NetScaler ADC Instance. Minimum length =  1 Maximum length =  127",
			},
			"fips_partition_name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "FIPS Partition Name.",
			},
			"nsvlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "VLAN Id. Maximum value =  ",
			},
			"num_pes": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Total number of PEs.",
			},
			"burst_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Burst Priority.",
			},
			"metrics_collection": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Metrics Collection.",
			},
			"is_managed": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Is Managed.",
			},
			"nexthop_v6": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Nexthop V6.",
			},
			"ipv4_address": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "IPv4 Address.",
			},
			"profile_name": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Profile Name.",
			},
			"backplane": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Backplane.",
			},
			"network_interfaces": schema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network Interfaces.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_name": schema.StringAttribute{
							Required:    true,
							Computed:    false,
							Description: "Port Name.",
						},
						"name_server": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Name Server.",
						},
						"is_mgmt_ifc": schema.BoolAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Is Management Interface.",
						},
						"gateway": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Gateway.",
						},
						"vrid_list_ipv6": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Vrid List IPv6.",
						},
						"parent_id": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Parent Id.",
						},
						"vrid_list_ipv4": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Vrid List IPv4.",
						},
						"is_member_ifc": schema.BoolAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Is Member Interface.",
						},
						"mac_address": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "MAC Address.",
						},
						"netmask": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Netmask.",
						},
						"ip_address": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "IP Address.",
						},
						"l2_enabled": schema.BoolAttribute{
							Optional:    true,
							Computed:    false,
							Description: "L2 Enabled.",
						},
						"id": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Id.",
						},
						"interface_name": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Interface Name.",
						},
						"parent_name": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Parent Name.",
						},
						"vlan_whitelist_array": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    false,
							Description: "VLAN Whitelist Array.",
						},
						"mac_mode": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "MAC Mode.",
						},
						"managed_device_id": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Managed Device Id.",
						},
						"vlan": schema.Int64Attribute{
							Optional:    true,
							Computed:    false,
							Description: "VLAN.",
						},
						"vrid_list_ipv4_array": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    false,
							Description: "Vrid List IPv4 Array.",
						},
						"receiveuntagged": schema.BoolAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Receiveuntagged.",
						},
						"sdx_formation_network_id": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Sdx Formation Network Id.",
						},
						"vrid_list_ipv6_array": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    false,
							Description: "Vrid List IPv6 Array.",
						},
						"is_vlan_applied": schema.BoolAttribute{
							Optional:    true,
							Computed:    false,
							Description: "Is VLAN Applied.",
						},
						"vlan_whitelist": schema.StringAttribute{
							Optional:    true,
							Computed:    false,
							Description: "VLAN Whitelist.",
						},
					},
				},
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "State.",
			},
			"last_updated_time": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Last Updated Time.",
			},
			"license_edition": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "License Edition.",
			},
			"customid": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Custom ID.",
			},
			"license_grace_time": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Grace for this NetScaler Instance.",
			},
			"la_mgmt": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "LA Mgmt.",
			},
			"vlan_id_0_2": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "VLAN id for the management interface 0/2. Maximum value =  ",
			},
			"number_of_ssl_cores": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Assign number of ssl virtual functions to VM Instance.",
			},
			"number_of_ssl_cards": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of SSL Cards.",
			},
			"max_burst_throughput": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Maximum burst throughput in Mbps of VM Instance.",
			},
			"config_type": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Configuration Type. Values: 0: IPv4, 1: IPv6, 2: Both. Maximum value =  ",
			},
			"if_internal_ip_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Set as true if VPX is managed by internal network (not required to be set for SDWAN).",
			},
			"cmd_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "true if you want to allow shell/sftp/scp access to NetScaler ADC Instance administrator. Minimum length =  1 Maximum length =  1024.",
			},
			"node_id": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Node identification of a device.",
			},
			"throughput_allocation_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Throughput Allocation Mode: 0-Fixed, 1-Burst-able.",
			},
			"number_of_cores": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Number of Cores.",
			},
			"pps": schema.StringAttribute{
				Optional:    true,
				Computed:    false,
				Description: "Assign packets per seconds to NetScaler ADC Instance. Maximum value =  ",
			},
			"plugin_netmask": schema.StringAttribute{
				Description: "Signaling Netmask. Minimum length =  1 Maximum length =  64",
				Optional:    true,
				Computed:    false,
			},
			"display_name": schema.StringAttribute{
				Description: "Display Name for this provision VPX. For HA pair it will be A-B, and for Cluster it will be CLIP. Minimum length =  1 Maximum length =  128",
				Optional:    true,
				Computed:    false,
			},
			"std_bw_total": schema.StringAttribute{
				Description: "Standard Bandwidth.",
				Optional:    true,
				Computed:    false,
			},
			"nexthop": schema.StringAttribute{
				Description: "Next Hop IP address. Minimum length =  1 Maximum length =  64",
				Optional:    true,
				Computed:    false,
			},
			"plt_bw_config": schema.StringAttribute{
				Description: "Platinum Bandwidth configured.",
				Optional:    true,
				Computed:    false,
			},
			"save_config": schema.StringAttribute{
				Description: "Should config be saved first in case instance is rebooted while modify.",
				Optional:    true,
				Computed:    false,
			},
			"is_new_crypto": schema.StringAttribute{
				Description: "True if number_of_acu/number_of_scu are used, false if number_of_ssl_cores is used.",
				Optional:    true,
				Computed:    false,
			},
			"profile_username": schema.StringAttribute{
				Description: "User Name specified by the user for this NetScaler ADC Instance.. Minimum length =  1 Maximum length =  128",
				Optional:    true,
				Computed:    false,
			},
			"profile_password": schema.StringAttribute{
				Description: "Password specified by the user for this NetScaler ADC Instance.. Minimum length =  1 Maximum length =  128",
				Optional:    true,
				Computed:    false,
			},
			"l2_enabled": schema.StringAttribute{
				Description: "L2mode status of VM Instance.",
				Optional:    true,
				Computed:    false,
			},
			"reboot_vm_on_cpu_change": schema.StringAttribute{
				Description: "Reboot VMs on CPU change during resource allocation.",
				Optional:    true,
				Computed:    false,
			},
			"nsvlan_interfaces": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "VLAN Interfaces. Minimum length =  1 Maximum length =  50",
				Optional:    true,
				Computed:    false,
			},
			"crypto_change_requires_reboot": schema.StringAttribute{
				Description: "True if the current changes made by user requires a reboot of the VM else false.",
				Optional:    true,
				Computed:    false,
			},
			"id": schema.StringAttribute{
				Description: "ID of the NetScaler ADC Instance.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

// Create creates a new resources and adds it into the Terraform state.
func (r *provisionVpxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "[DEBUG] In Create Method of provisionVpxResource")

	var data *provisionVpxResourceModel

	// Read Terraform plan data into the model( proVpxResourceModel )
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	provisionVpxReq := getThePayloadFromtheConfig(ctx, data)

	endpoint := "ns"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, provisionVpxReq)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Resource",
			fmt.Sprintf("Error creating resource: %s", err.Error()),
		)
		return
	}

	vpxID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Wait for the VPX to be up
	for {
		time.Sleep(5 * time.Second)

		returnData, err := r.client.GetResource(endpoint, vpxID)
		if err != nil {
			// return diag.Errorf("unable to get VPX: %s", err.Error())
			resp.Diagnostics.AddError(
				"Error Getting Resource",
				fmt.Sprintf("Error getting resource: %s", err.Error()),
			)
			return
		}
		instanceState := returnData[endpoint].([]interface{})[0].(map[string]interface{})["instance_state"].(string)
		if instanceState == "Up" {
			break
		}
		tflog.Debug(ctx, fmt.Sprintf("VPX instance_state is %s", instanceState))
	}

	// Set the ID of the created resource
	data.Id = types.StringValue(vpxID)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	rreq := resource.ReadRequest{
		State:        resp.State,
		ProviderMeta: req.ProviderMeta,
	}
	rresp := resource.ReadResponse{
		State:       resp.State,
		Diagnostics: resp.Diagnostics,
	}
	r.Read(ctx, rreq, &rresp)

	*resp = resource.CreateResponse{
		State:       rresp.State,
		Diagnostics: rresp.Diagnostics,
	}

}

func getThePayloadFromtheConfig(ctx context.Context, data *provisionVpxResourceModel) provisionVpxResourceReq {
	tflog.Debug(ctx, "In getThePayloadFromtheConfig Function of provisionVpxResource")

	// Create a struct for the request body
	provisionVpxReq := provisionVpxResourceReq{
		IPAddress:                  data.IPAddress.ValueString(),
		StdBwConfig:                data.StdBwConfig.ValueString(),
		Name:                       data.Name.ValueString(),
		MastoolsVersion:            data.MastoolsVersion.ValueString(),
		PluginIPAddress:            data.PluginIPAddress.ValueString(),
		VlanType:                   data.VlanType.ValueString(),
		EntBwTotal:                 data.EntBwTotal.ValueString(),
		VcpuConfig:                 data.VcpuConfig.ValueString(),
		NsvlanTagged:               data.NsvlanTagged.ValueString(),
		Netmask:                    data.Netmask.ValueString(),
		EntBwConfig:                data.EntBwConfig.ValueString(),
		DatacenterID:               data.DatacenterID.ValueString(),
		InstanceMode:               data.InstanceMode.ValueString(),
		NumberOfSslCoresUp:         data.NumberOfSslCoresUp.ValueString(),
		StdBwAvailable:             data.StdBwAvailable.ValueString(),
		InternalIPAddress:          data.InternalIPAddress.ValueString(),
		If01:                       data.If01.ValueString(),
		PltBwTotal:                 data.PltBwTotal.ValueString(),
		HostIPAddress:              data.HostIPAddress.ValueString(),
		VpxID:                      data.VpxID.ValueString(),
		IPv6Address:                data.IPv6Address.ValueString(),
		MgmtIPAddress:              data.MgmtIPAddress.ValueString(),
		NumberOfAcu:                data.NumberOfAcu.ValueString(),
		PltBwAvailable:             data.PltBwAvailable.ValueString(),
		IsClip:                     data.IsClip.ValueString(),
		DeviceFamily:               data.DeviceFamily.ValueString(),
		Type:                       data.Type.ValueString(),
		Throughput:                 data.Throughput.ValueString(),
		TemplateName:               data.TemplateName.ValueString(),
		Gateway:                    data.Gateway.ValueString(),
		Iscco:                      data.Iscco.ValueString(),
		NumberOfScu:                data.NumberOfScu.ValueString(),
		License:                    data.License.ValueString(),
		DomainName:                 data.DomainName.ValueString(),
		If02:                       data.If02.ValueString(),
		ImageName:                  data.ImageName.ValueString(),
		Hostname:                   data.Hostname.ValueString(),
		VmMemoryTotal:              data.VmMemoryTotal.ValueString(),
		EntBwAvailable:             data.EntBwAvailable.ValueString(),
		Description:                data.Description.ValueString(),
		Username:                   data.Username.ValueString(),
		FipsPartitionName:          data.FipsPartitionName.ValueString(),
		NsvlanID:                   data.NsvlanID.ValueString(),
		NumPes:                     data.NumPes.ValueString(),
		BurstPriority:              data.BurstPriority.ValueString(),
		MetricsCollection:          data.MetricsCollection.ValueString(),
		IsManaged:                  data.IsManaged.ValueString(),
		NexthopV6:                  data.NexthopV6.ValueString(),
		IPv4Address:                data.IPv4Address.ValueString(),
		ProfileName:                data.ProfileName.ValueString(),
		Backplane:                  data.Backplane.ValueString(),
		NetworkInterfaces:          networkInterfaceFromConfigToSchema(ctx, data.NetworkInterfaces),
		State:                      data.State.ValueString(),
		LastUpdatedTime:            data.LastUpdatedTime.ValueString(),
		LicenseEdition:             data.LicenseEdition.ValueString(),
		Customid:                   data.Customid.ValueString(),
		LicenseGraceTime:           data.LicenseGraceTime.ValueString(),
		LaMgmt:                     data.LaMgmt.ValueString(),
		VlanID02:                   data.VlanID02.ValueString(),
		NumberOfSslCores:           data.NumberOfSslCores.ValueString(),
		NumberOfsslCards:           data.NumberOfsslCards.ValueString(),
		MaxBurstThroughput:         data.MaxBurstThroughput.ValueString(),
		ConfigType:                 data.ConfigType.ValueString(),
		IfInternalIpEnabled:        data.IfInternalIpEnabled.ValueString(),
		CmdPolicy:                  data.CmdPolicy.ValueString(),
		NodeId:                     data.NodeId.ValueString(),
		ThroughputAllocationMode:   data.ThroughputAllocationMode.ValueString(),
		NumberOfCores:              data.NumberOfCores.ValueString(),
		Pps:                        data.Pps.ValueString(),
		PluginNetmask:              data.PluginNetmask.ValueString(),
		DisplayName:                data.DisplayName.ValueString(),
		StdBwTotal:                 data.StdBwTotal.ValueString(),
		Nexthop:                    data.Nexthop.ValueString(),
		PltBwConfig:                data.PltBwConfig.ValueString(),
		CryptoChangeRequiresReboot: data.CryptoChangeRequiresReboot.ValueString(),
		NsvlanInterfaces:           utils.TypeListToUnmarshalStringList(data.NsvlanInterfaces),
		RebootVmOnCpuChange:        data.RebootVmOnCpuChange.ValueString(),
		L2Enabled:                  data.L2Enabled.ValueString(),
		ProfilePassword:            data.ProfilePassword.ValueString(),
		ProfileUsername:            data.ProfileUsername.ValueString(),
		IsNewCrypto:                data.IsNewCrypto.ValueString(),
		SaveConfig:                 data.SaveConfig.ValueString(),
		GatewayIPv6:                data.GatewayIPv6.ValueString(),
		VlanID01:                   data.VlanID01.ValueString(),
		Password:                   data.Password.ValueString(),
		ThroughputLimit:            data.ThroughputLimit.ValueString(),
		NsIPAddress:                data.NsIPAddress.ValueString(),
	}

	return provisionVpxReq
}

func networkInterfaceFromConfigToSchema(ctx context.Context, nidata basetypes.ListValue) []map[string]interface{} {
	tflog.Debug(ctx, "In networkInterfaceFromConfigToSchema func of provisionVpxResource")

	var networkInterfacesMapList []map[string]interface{}
	// Get the network interfaces from the model
	for _, internalObjectAttr := range nidata.Elements() {
		var internalObjectMap basetypes.ObjectValue
		internalObjectMap = internalObjectAttr.(basetypes.ObjectValue)

		mapNi := make(map[string]interface{}, 0)

		for key, val := range internalObjectMap.Attributes() {
			if !val.IsNull() {
				switch val.Type(ctx) {
				case types.StringType:
					mapNi[key] = val.(basetypes.StringValue).ValueString()
				case types.Int64Type:
					mapNi[key] = val.(basetypes.Int64Value).ValueInt64()
				case types.BoolType:
					mapNi[key] = val.(basetypes.BoolValue).ValueBool()
				case types.ListType{ElemType: types.StringType}:
					if len(val.(basetypes.ListValue).Elements()) == 0 {
						mapNi[key] = []string{}
						continue
					}
					var listVal []string
					for _, vals := range val.(basetypes.ListValue).Elements() {
						listVal = append(listVal, vals.(basetypes.StringValue).ValueString())
					}
					mapNi[key] = listVal
				}
			}
		}
		networkInterfacesMapList = append(networkInterfacesMapList, mapNi)

	}
	return networkInterfacesMapList
}

func (r *provisionVpxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Read Method of provisionVpxResource with Id: %s", resId))

	var state *provisionVpxResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "ns"

	data, err := r.client.GetResource(endpoint, state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Resource",
			fmt.Sprintf("Error reading resource: %s", err.Error()),
		)
		return
	}

	getResponseData := data[endpoint].([]interface{})[0].(map[string]interface{})

	if !state.IPAddress.IsNull() {
		state.IPAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !state.StdBwConfig.IsNull() {
		state.StdBwConfig = types.StringValue(getResponseData["std_bw_config"].(string))
	}
	if !state.NsIPAddress.IsNull() {
		state.NsIPAddress = types.StringValue(getResponseData["ns_ip_address"].(string))
	}
	// if !state.Password.IsNull() {
	// 	state.Password = types.StringValue(getResponseData["password"].(string))
	// }
	if !state.GatewayIPv6.IsNull() {
		state.GatewayIPv6 = types.StringValue(getResponseData["gateway_ipv6"].(string))
	}
	if !state.ThroughputLimit.IsNull() {
		state.ThroughputLimit = types.StringValue(getResponseData["throughput_limit"].(string))
	}
	if !state.VlanID01.IsNull() {
		state.VlanID01 = types.StringValue(getResponseData["vlan_id_0_1"].(string))
	}
	if !state.Name.IsNull() {
		state.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !state.MastoolsVersion.IsNull() {
		state.MastoolsVersion = types.StringValue(getResponseData["mastools_version"].(string))
	}
	if !state.PluginIPAddress.IsNull() {
		state.PluginIPAddress = types.StringValue(getResponseData["plugin_ip_address"].(string))
	}
	if !state.VlanType.IsNull() {
		state.VlanType = types.StringValue(getResponseData["vlan_type"].(string))
	}
	if !state.EntBwTotal.IsNull() {
		state.EntBwTotal = types.StringValue(getResponseData["ent_bw_total"].(string))
	}
	if !state.VcpuConfig.IsNull() {
		state.VcpuConfig = types.StringValue(getResponseData["vcpu_config"].(string))
	}
	if !state.NsvlanTagged.IsNull() {
		state.NsvlanTagged = types.StringValue(getResponseData["nsvlan_tagged"].(string))
	}
	if !state.Netmask.IsNull() {
		state.Netmask = types.StringValue(getResponseData["netmask"].(string))
	}
	if !state.EntBwConfig.IsNull() {
		state.EntBwConfig = types.StringValue(getResponseData["ent_bw_config"].(string))
	}
	if !state.DatacenterID.IsNull() {
		state.DatacenterID = types.StringValue(getResponseData["datacenter_id"].(string))
	}
	if !state.InstanceMode.IsNull() {
		state.InstanceMode = types.StringValue(getResponseData["instance_mode"].(string))
	}
	if !state.NumberOfSslCoresUp.IsNull() {
		state.NumberOfSslCoresUp = types.StringValue(getResponseData["number_of_ssl_cores_up"].(string))
	}
	if !state.StdBwAvailable.IsNull() {
		state.StdBwAvailable = types.StringValue(getResponseData["std_bw_available"].(string))
	}
	if !state.InternalIPAddress.IsNull() {
		state.InternalIPAddress = types.StringValue(getResponseData["internal_ip_address"].(string))
	}
	if !state.If01.IsNull() {
		state.If01 = types.StringValue(getResponseData["if_0_1"].(string))
	}
	if !state.PltBwTotal.IsNull() {
		state.PltBwTotal = types.StringValue(getResponseData["plt_bw_total"].(string))
	}
	if !state.HostIPAddress.IsNull() {
		state.HostIPAddress = types.StringValue(getResponseData["host_ip_address"].(string))
	}
	// if !state.VpxID.IsNull() {
	// 	state.VpxID = types.StringValue(getResponseData["vpx_id"].(string))
	// }
	if !state.IPv6Address.IsNull() {
		state.IPv6Address = types.StringValue(getResponseData["ipv6_address"].(string))
	}
	if !state.MgmtIPAddress.IsNull() {
		state.MgmtIPAddress = types.StringValue(getResponseData["mgmt_ip_address"].(string))
	}
	if !state.NumberOfAcu.IsNull() {
		state.NumberOfAcu = types.StringValue(getResponseData["number_of_acu"].(string))
	}
	if !state.PltBwAvailable.IsNull() {
		state.PltBwAvailable = types.StringValue(getResponseData["plt_bw_available"].(string))
	}
	if !state.IsClip.IsNull() {
		state.IsClip = types.StringValue(getResponseData["is_clip"].(string))
	}
	if !state.DeviceFamily.IsNull() {
		state.DeviceFamily = types.StringValue(getResponseData["device_family"].(string))
	}
	if !state.Type.IsNull() {
		state.Type = types.StringValue(getResponseData["type"].(string))
	}
	if !state.Throughput.IsNull() {
		state.Throughput = types.StringValue(getResponseData["throughput"].(string))
	}
	if !state.TemplateName.IsNull() {
		state.TemplateName = types.StringValue(getResponseData["template_name"].(string))
	}
	if !state.Gateway.IsNull() {
		state.Gateway = types.StringValue(getResponseData["gateway"].(string))
	}
	if !state.Iscco.IsNull() {
		state.Iscco = types.StringValue(getResponseData["iscco"].(string))
	}
	if !state.NumberOfScu.IsNull() {
		state.NumberOfScu = types.StringValue(getResponseData["number_of_scu"].(string))
	}
	if !state.License.IsNull() {
		state.License = types.StringValue(getResponseData["license"].(string))
	}
	if !state.DomainName.IsNull() {
		state.DomainName = types.StringValue(getResponseData["domain_name"].(string))
	}
	if !state.If02.IsNull() {
		state.If02 = types.StringValue(getResponseData["if_0_2"].(string))
	}
	// if !state.ImageName.IsNull() {
	// 	state.ImageName = types.StringValue(getResponseData["image_name"].(string))	 // FIXME: API Problem. image_name is empty after Update operataion
	// }
	if !state.Hostname.IsNull() {
		state.Hostname = types.StringValue(getResponseData["hostname"].(string))
	}
	if !state.VmMemoryTotal.IsNull() {
		state.VmMemoryTotal = types.StringValue(getResponseData["vm_memory_total"].(string))
	}
	if !state.EntBwAvailable.IsNull() {
		state.EntBwAvailable = types.StringValue(getResponseData["ent_bw_available"].(string))
	}
	if !state.Description.IsNull() {
		state.Description = types.StringValue(getResponseData["description"].(string))
	}
	if !state.Username.IsNull() {
		state.Username = types.StringValue(getResponseData["username"].(string))
	}
	if !state.FipsPartitionName.IsNull() {
		state.FipsPartitionName = types.StringValue(getResponseData["fips_partition_name"].(string))
	}
	if !state.NsvlanID.IsNull() {
		state.NsvlanID = types.StringValue(getResponseData["nsvlan_id"].(string))
	}
	if !state.NumPes.IsNull() {
		state.NumPes = types.StringValue(getResponseData["num_pes"].(string))
	}
	if !state.BurstPriority.IsNull() {
		state.BurstPriority = types.StringValue(getResponseData["burst_priority"].(string))
	}
	if !state.MetricsCollection.IsNull() {
		state.MetricsCollection = types.StringValue(getResponseData["metrics_collection"].(string))
	}
	if !state.IsManaged.IsNull() {
		state.IsManaged = types.StringValue(getResponseData["is_managed"].(string))
	}
	if !state.NexthopV6.IsNull() {
		state.NexthopV6 = types.StringValue(getResponseData["nexthop_v6"].(string))
	}
	if !state.IPv4Address.IsNull() {
		state.IPv4Address = types.StringValue(getResponseData["ipv4_address"].(string))
	}
	if !state.ProfileName.IsNull() {
		state.ProfileName = types.StringValue(getResponseData["profile_name"].(string))
	}
	if !state.Backplane.IsNull() {
		state.Backplane = types.StringValue(getResponseData["backplane"].(string))
	}
	if !state.NetworkInterfaces.IsNull() {
		state.NetworkInterfaces = networkinterfaceTonetworkinterfaceTF(getResponseData["network_interfaces"].([]interface{}), state.NetworkInterfaces, ctx)
	}
	if !state.State.IsNull() {
		state.State = types.StringValue(getResponseData["state"].(string))
	}
	if !state.LastUpdatedTime.IsNull() {
		state.LastUpdatedTime = types.StringValue(getResponseData["last_updated_time"].(string))
	}
	if !state.LicenseEdition.IsNull() {
		state.LicenseEdition = types.StringValue(getResponseData["license_edition"].(string))
	}
	if !state.Customid.IsNull() {
		state.Customid = types.StringValue(getResponseData["customid"].(string))
	}
	if !state.LicenseGraceTime.IsNull() {
		state.LicenseGraceTime = types.StringValue(getResponseData["license_grace_time"].(string))
	}
	if !state.LaMgmt.IsNull() {
		state.LaMgmt = types.StringValue(getResponseData["la_mgmt"].(string))
	}
	if !state.VlanID02.IsNull() {
		state.VlanID02 = types.StringValue(getResponseData["vlan_id_0_2"].(string))
	}
	if !state.NumberOfSslCores.IsNull() {
		state.NumberOfSslCores = types.StringValue(getResponseData["number_of_ssl_cores"].(string))
	}
	if !state.NumberOfsslCards.IsNull() {
		state.NumberOfsslCards = types.StringValue(getResponseData["number_of_ssl_cards"].(string))
	}
	if !state.MaxBurstThroughput.IsNull() {
		state.MaxBurstThroughput = types.StringValue(getResponseData["max_burst_throughput"].(string))
	}
	if !state.ConfigType.IsNull() {
		state.ConfigType = types.StringValue(getResponseData["config_type"].(string))
	}
	if !state.IfInternalIpEnabled.IsNull() {
		state.IfInternalIpEnabled = types.StringValue(getResponseData["if_internal_ip_enabled"].(string))
	}
	if !state.CmdPolicy.IsNull() {
		state.CmdPolicy = types.StringValue(getResponseData["cmd_policy"].(string))
	}
	if !state.NodeId.IsNull() {
		state.NodeId = types.StringValue(getResponseData["node_id"].(string))
	}
	if !state.ThroughputAllocationMode.IsNull() {
		state.ThroughputAllocationMode = types.StringValue(getResponseData["throughput_allocation_mode"].(string))
	}
	if !state.NumberOfCores.IsNull() {
		state.NumberOfCores = types.StringValue(getResponseData["number_of_cores"].(string))
	}
	if !state.Pps.IsNull() {
		state.Pps = types.StringValue(getResponseData["pps"].(string))
	}
	if !state.PluginNetmask.IsNull() {
		state.PluginNetmask = types.StringValue(getResponseData["plugin_netmask"].(string))
	}
	if !state.DisplayName.IsNull() {
		state.DisplayName = types.StringValue(getResponseData["display_name"].(string))
	}
	if !state.StdBwTotal.IsNull() {
		state.StdBwTotal = types.StringValue(getResponseData["std_bw_total"].(string))
	}
	if !state.Nexthop.IsNull() {
		state.Nexthop = types.StringValue(getResponseData["nexthop"].(string))
	}
	if !state.PltBwConfig.IsNull() {
		state.PltBwConfig = types.StringValue(getResponseData["plt_bw_config"].(string))
	}
	if !state.CryptoChangeRequiresReboot.IsNull() {
		state.CryptoChangeRequiresReboot = types.StringValue(getResponseData["crypto_change_requires_reboot"].(string))
	}
	if !state.NsvlanInterfaces.IsNull() {
		state.NsvlanInterfaces = utils.StringListToTypeList(utils.ToStringList(getResponseData["nsvlan_interfaces"].([]interface{})))
	}
	if !state.RebootVmOnCpuChange.IsNull() {
		state.RebootVmOnCpuChange = types.StringValue(getResponseData["reboot_vm_on_cpu_change"].(string))
	}
	if !state.L2Enabled.IsNull() {
		state.L2Enabled = types.StringValue(getResponseData["l2_enabled"].(string))
	}
	if !state.ProfilePassword.IsNull() {
		state.ProfilePassword = types.StringValue(getResponseData["profile_password"].(string))
	}
	if !state.ProfileUsername.IsNull() {
		state.ProfileUsername = types.StringValue(getResponseData["profile_username"].(string))
	}
	if !state.IsNewCrypto.IsNull() {
		state.IsNewCrypto = types.StringValue(getResponseData["is_new_crypto"].(string))
	}
	if !state.SaveConfig.IsNull() {
		state.SaveConfig = types.StringValue(getResponseData["save_config"].(string))
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func networkinterfaceTonetworkinterfaceTF(nif []interface{}, stateNetworkInterface basetypes.ListValue, ctx context.Context) basetypes.ListValue {
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
		"ip_address",
		"sdx_formation_network_id",
	}

	var nifs []map[string]interface{}

	// if network_interfaces is not set in the state file, then return empty list
	if len(stateNetworkInterface.Elements()) == 0 {
		return basetypes.ListValue{}
	}

	for _, v := range stateNetworkInterface.Elements() {

		inputNifs := v.(basetypes.ObjectValue).Attributes()

		// get the portnames of all the inputNifs
		var inputNifPortNames []string
		for key, inputNif := range inputNifs {
			if key == "port_name" {
				inputNif := inputNif.(basetypes.StringValue).ValueString()
				inputNifPortNames = append(inputNifPortNames, inputNif)
			}
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
			// var nifMap2 map[string]interface{}
			nifMap2 := make(map[string]interface{})
			for k, v := range nifMap {
				if service.Contains(nifSchemaAttributes, k) {
					nifMap2[k] = v
				}
			}
			nifs = append(nifs, nifMap2)
		}
	}
	var nifList []attr.Value
	for listindex, nif := range nifs {
		var networkInterfaceElementsInState []string
		stateNetworkInterfaceElements := stateNetworkInterface.Elements()[listindex].(basetypes.ObjectValue)
		for key, val := range stateNetworkInterfaceElements.Attributes() {
			if !val.IsNull() {
				networkInterfaceElementsInState = append(networkInterfaceElementsInState, key)
			}
		}
		var newMap basetypes.ObjectValue
		nifmap := make(map[string]attr.Value)
		nifMapTypes := make(map[string]attr.Type)
		for k, v := range nif {
			if !service.Contains(networkInterfaceElementsInState, k) {
				if k == "is_vlan_applied" || k == "is_mgmt_ifc" || k == "is_member_ifc" || k == "l2_enabled" || k == "receiveuntagged" {
					nifmap[k] = basetypes.NewBoolNull()
					nifMapTypes[k] = types.BoolType
				} else if k == "vlan" {
					nifmap[k] = basetypes.NewInt64Null()
					nifMapTypes[k] = types.Int64Type
				} else if k == "vrid_list_ipv4_array" || k == "vrid_list_ipv6_array" || k == "vlan_whitelist_array" {
					nifmap[k] = basetypes.NewListNull(types.StringType)
					nifMapTypes[k] = types.ListType{ElemType: types.StringType}
				} else {
					nifmap[k] = basetypes.NewStringNull()
					nifMapTypes[k] = types.StringType
				}
			} else {
				if k == "is_vlan_applied" || k == "is_mgmt_ifc" || k == "is_member_ifc" || k == "l2_enabled" || k == "receiveuntagged" {
					val, _ := strconv.ParseBool(v.(string))
					nifmap[k] = basetypes.NewBoolValue(val)
					nifMapTypes[k] = types.BoolType
				} else if k == "vlan" {
					val, _ := strconv.ParseInt(v.(string), 10, 64)
					nifmap[k] = basetypes.NewInt64Value(val)
					nifMapTypes[k] = types.Int64Type
				} else if k == "vrid_list_ipv4_array" || k == "vrid_list_ipv6_array" || k == "vlan_whitelist_array" {
					if len(v.([]interface{})) == 0 {
						val := make([]attr.Value, 0)
						nifmap[k], _ = basetypes.NewListValue(types.StringType, val)
					} else {
						var stringSlice []attr.Value
						for _, item := range v.([]interface{}) {
							if str, ok := item.(string); ok {
								stringSlice = append(stringSlice, basetypes.NewStringValue(str))
							}
							nifmap[k], _ = basetypes.NewListValue(types.StringType, stringSlice)
						}
					}
					nifMapTypes[k] = types.ListType{ElemType: types.StringType}
				} else {
					nifmap[k] = basetypes.NewStringValue(v.(string))
					nifMapTypes[k] = types.StringType
				}
			}
		}

		newMap, _ = basetypes.NewObjectValue(nifMapTypes, nifmap)
		nifList = append(nifList, newMap)

	}
	newNifList, _ := basetypes.NewListValueFrom(ctx, stateNetworkInterface.ElementType(ctx), nifList)

	return newNifList

}

func (r *provisionVpxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Update Method of provisionVpxResource with Id: %s", resId))

	var plan provisionVpxResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state provisionVpxResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "ns"
	requestPayload := getThePayloadFromtheConfig(ctx, &plan)
	plan.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, resourceId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Resource",
			fmt.Sprintf("Error updating resource: %s", err.Error()),
		)
		return
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *provisionVpxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Delete Method of provisionVpxResource with Id: %s", resId))

	var data *provisionVpxResourceModel

	// Read terraform prior state data into the model
	diag := req.State.Get(ctx, &data)

	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "ns"
	// Delete the resource
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Resource",
			fmt.Sprintf("Error deleting resource: %s", err.Error()),
		)
		return
	}
}
