package ns

import (
	"context"
	"fmt"

	"time"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &provisionVpxResource{}
	_ resource.ResourceWithConfigure   = &provisionVpxResource{}
	_ resource.ResourceWithImportState = &provisionVpxResource{}
)

// ImportState implements resource.ResourceWithImportState.
func (r *provisionVpxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// provisionVpxResource defines the resource implementation.
type provisionVpxResource struct {
	client *service.NitroClient
}

func ProvisionVpxResource() resource.Resource {
	return &provisionVpxResource{}
}

// provisionVpxResourceModel describes the resource data model.

type provisionVpxResourceModel struct {
	Backplane                  types.String `tfsdk:"backplane"`
	BurstPriority              types.Int64  `tfsdk:"burst_priority"`
	CmdPolicy                  types.String `tfsdk:"cmd_policy"`
	ConfigType                 types.Int64  `tfsdk:"config_type"`
	CryptoChangeRequiresReboot types.Bool   `tfsdk:"crypto_change_requires_reboot"`
	Customid                   types.String `tfsdk:"customid"`
	DatacenterId               types.String `tfsdk:"datacenter_id"`
	Description                types.String `tfsdk:"description"`
	DeviceFamily               types.String `tfsdk:"device_family"`
	DisplayName                types.String `tfsdk:"display_name"`
	DomainName                 types.String `tfsdk:"domain_name"`
	EntBwAvailable             types.Int64  `tfsdk:"ent_bw_available"`
	EntBwConfig                types.Int64  `tfsdk:"ent_bw_config"`
	EntBwTotal                 types.Int64  `tfsdk:"ent_bw_total"`
	FipsPartitionName          types.String `tfsdk:"fips_partition_name"`
	Gateway                    types.String `tfsdk:"gateway"`
	GatewayIpv6                types.String `tfsdk:"gateway_ipv6"`
	HostIpAddress              types.String `tfsdk:"host_ip_address"`
	Hostname                   types.String `tfsdk:"hostname"`
	If01                       types.Bool   `tfsdk:"if_0_1"`
	If02                       types.Bool   `tfsdk:"if_0_2"`
	IfInternalIpEnabled        types.Bool   `tfsdk:"if_internal_ip_enabled"`
	ImageName                  types.String `tfsdk:"image_name"`
	InstanceMode               types.String `tfsdk:"instance_mode"`
	InternalIpAddress          types.String `tfsdk:"internal_ip_address"`
	IpAddress                  types.String `tfsdk:"ip_address"`
	Ipv4Address                types.String `tfsdk:"ipv4_address"`
	Ipv6Address                types.String `tfsdk:"ipv6_address"`
	IsClip                     types.Bool   `tfsdk:"is_clip"`
	IsManaged                  types.Bool   `tfsdk:"is_managed"`
	IsNewCrypto                types.Bool   `tfsdk:"is_new_crypto"`
	Iscco                      types.Bool   `tfsdk:"iscco"`
	L2Enabled                  types.Bool   `tfsdk:"l2_enabled"`
	LaMgmt                     types.Bool   `tfsdk:"la_mgmt"`
	LastUpdatedTime            types.Int64  `tfsdk:"last_updated_time"`
	License                    types.String `tfsdk:"license"`
	LicenseEdition             types.String `tfsdk:"license_edition"`
	LicenseGraceTime           types.Int64  `tfsdk:"license_grace_time"`
	MastoolsVersion            types.String `tfsdk:"mastools_version"`
	MaxBurstThroughput         types.Int64  `tfsdk:"max_burst_throughput"`
	MetricsCollection          types.Bool   `tfsdk:"metrics_collection"`
	MgmtIpAddress              types.String `tfsdk:"mgmt_ip_address"`
	Name                       types.String `tfsdk:"name"`
	Netmask                    types.String `tfsdk:"netmask"`
	NetworkInterfaces          types.List   `tfsdk:"network_interfaces"`
	Nexthop                    types.String `tfsdk:"nexthop"`
	NexthopV6                  types.String `tfsdk:"nexthop_v6"`
	NodeId                     types.String `tfsdk:"node_id"`
	NsIpAddress                types.String `tfsdk:"ns_ip_address"`
	NsvlanId                   types.Int64  `tfsdk:"nsvlan_id"`
	NsvlanInterfaces           types.List   `tfsdk:"nsvlan_interfaces"`
	NsvlanTagged               types.Bool   `tfsdk:"nsvlan_tagged"`
	NumPes                     types.Int64  `tfsdk:"num_pes"`
	NumberOfAcu                types.Int64  `tfsdk:"number_of_acu"`
	NumberOfCores              types.Int64  `tfsdk:"number_of_cores"`
	NumberOfScu                types.Int64  `tfsdk:"number_of_scu"`
	NumberOfSslCards           types.Int64  `tfsdk:"number_of_ssl_cards"`
	NumberOfSslCores           types.Int64  `tfsdk:"number_of_ssl_cores"`
	NumberOfSslCoresUp         types.Int64  `tfsdk:"number_of_ssl_cores_up"`
	Password                   types.String `tfsdk:"password"`
	PltBwAvailable             types.Int64  `tfsdk:"plt_bw_available"`
	PltBwConfig                types.Int64  `tfsdk:"plt_bw_config"`
	PltBwTotal                 types.Int64  `tfsdk:"plt_bw_total"`
	PluginIpAddress            types.String `tfsdk:"plugin_ip_address"`
	PluginNetmask              types.String `tfsdk:"plugin_netmask"`
	Pps                        types.Int64  `tfsdk:"pps"`
	ProfileName                types.String `tfsdk:"profile_name"`
	ProfilePassword            types.String `tfsdk:"profile_password"`
	ProfileUsername            types.String `tfsdk:"profile_username"`
	RebootVmOnCpuChange        types.Bool   `tfsdk:"reboot_vm_on_cpu_change"`
	SaveConfig                 types.Bool   `tfsdk:"save_config"`
	State                      types.String `tfsdk:"state"`
	StdBwAvailable             types.Int64  `tfsdk:"std_bw_available"`
	StdBwConfig                types.Int64  `tfsdk:"std_bw_config"`
	StdBwTotal                 types.Int64  `tfsdk:"std_bw_total"`
	TemplateName               types.String `tfsdk:"template_name"`
	Throughput                 types.Int64  `tfsdk:"throughput"`
	ThroughputAllocationMode   types.Int64  `tfsdk:"throughput_allocation_mode"`
	ThroughputLimit            types.Int64  `tfsdk:"throughput_limit"`
	Type                       types.String `tfsdk:"type"`
	Username                   types.String `tfsdk:"username"`
	VcpuConfig                 types.Int64  `tfsdk:"vcpu_config"`
	VlanId01                   types.Int64  `tfsdk:"vlan_id_0_1"`
	VlanId02                   types.Int64  `tfsdk:"vlan_id_0_2"`
	VlanType                   types.Int64  `tfsdk:"vlan_type"`
	VmMemoryTotal              types.Int64  `tfsdk:"vm_memory_total"`
	Id                         types.String `tfsdk:"id"`
}

type provisionVpxResourceReq struct {
	Backplane                  string                   `json:"backplane,omitempty"`
	BurstPriority              *int64                   `json:"burst_priority,omitempty"`
	CmdPolicy                  string                   `json:"cmd_policy,omitempty"`
	ConfigType                 *int64                   `json:"config_type,omitempty"`
	CryptoChangeRequiresReboot *bool                    `json:"crypto_change_requires_reboot,omitempty"`
	Customid                   string                   `json:"customid,omitempty"`
	DatacenterId               string                   `json:"datacenter_id,omitempty"`
	Description                string                   `json:"description,omitempty"`
	DeviceFamily               string                   `json:"device_family,omitempty"`
	DisplayName                string                   `json:"display_name,omitempty"`
	DomainName                 string                   `json:"domain_name,omitempty"`
	EntBwAvailable             *int64                   `json:"ent_bw_available,omitempty"`
	EntBwConfig                *int64                   `json:"ent_bw_config,omitempty"`
	EntBwTotal                 *int64                   `json:"ent_bw_total,omitempty"`
	FipsPartitionName          string                   `json:"fips_partition_name,omitempty"`
	Gateway                    string                   `json:"gateway,omitempty"`
	GatewayIpv6                string                   `json:"gateway_ipv6,omitempty"`
	HostIpAddress              string                   `json:"host_ip_address,omitempty"`
	Hostname                   string                   `json:"hostname,omitempty"`
	If01                       *bool                    `json:"if_0_1,omitempty"`
	If02                       *bool                    `json:"if_0_2,omitempty"`
	IfInternalIpEnabled        *bool                    `json:"if_internal_ip_enabled,omitempty"`
	ImageName                  string                   `json:"image_name,omitempty"`
	InstanceMode               string                   `json:"instance_mode,omitempty"`
	InternalIpAddress          string                   `json:"internal_ip_address,omitempty"`
	IpAddress                  string                   `json:"ip_address,omitempty"`
	Ipv4Address                string                   `json:"ipv4_address,omitempty"`
	Ipv6Address                string                   `json:"ipv6_address,omitempty"`
	IsClip                     *bool                    `json:"is_clip,omitempty"`
	IsManaged                  *bool                    `json:"is_managed,omitempty"`
	IsNewCrypto                *bool                    `json:"is_new_crypto,omitempty"`
	Iscco                      *bool                    `json:"iscco,omitempty"`
	L2Enabled                  *bool                    `json:"l2_enabled,omitempty"`
	LaMgmt                     *bool                    `json:"la_mgmt,omitempty"`
	LastUpdatedTime            *int64                   `json:"last_updated_time,omitempty"`
	License                    string                   `json:"license,omitempty"`
	LicenseEdition             string                   `json:"license_edition,omitempty"`
	LicenseGraceTime           *int64                   `json:"license_grace_time,omitempty"`
	MastoolsVersion            string                   `json:"mastools_version,omitempty"`
	MaxBurstThroughput         *int64                   `json:"max_burst_throughput,omitempty"`
	MetricsCollection          *bool                    `json:"metrics_collection,omitempty"`
	MgmtIpAddress              string                   `json:"mgmt_ip_address,omitempty"`
	Name                       string                   `json:"name,omitempty"`
	Netmask                    string                   `json:"netmask,omitempty"`
	NetworkInterfaces          []map[string]interface{} `json:"network_interfaces,omitempty"`
	Nexthop                    string                   `json:"nexthop,omitempty"`
	NexthopV6                  string                   `json:"nexthop_v6,omitempty"`
	NodeId                     string                   `json:"node_id,omitempty"`
	NsIpAddress                string                   `json:"ns_ip_address,omitempty"`
	NsvlanId                   *int64                   `json:"nsvlan_id,omitempty"`
	NsvlanInterfaces           []string                 `json:"nsvlan_interfaces,omitempty"`
	NsvlanTagged               *bool                    `json:"nsvlan_tagged,omitempty"`
	NumPes                     *int64                   `json:"num_pes,omitempty"`
	NumberOfAcu                *int64                   `json:"number_of_acu,omitempty"`
	NumberOfCores              *int64                   `json:"number_of_cores,omitempty"`
	NumberOfScu                *int64                   `json:"number_of_scu,omitempty"`
	NumberOfSslCards           *int64                   `json:"number_of_ssl_cards,omitempty"`
	NumberOfSslCores           *int64                   `json:"number_of_ssl_cores,omitempty"`
	NumberOfSslCoresUp         *int64                   `json:"number_of_ssl_cores_up,omitempty"`
	Password                   string                   `json:"password,omitempty"`
	PltBwAvailable             *int64                   `json:"plt_bw_available,omitempty"`
	PltBwConfig                *int64                   `json:"plt_bw_config,omitempty"`
	PltBwTotal                 *int64                   `json:"plt_bw_total,omitempty"`
	PluginIpAddress            string                   `json:"plugin_ip_address,omitempty"`
	PluginNetmask              string                   `json:"plugin_netmask,omitempty"`
	Pps                        *int64                   `json:"pps,omitempty"`
	ProfileName                string                   `json:"profile_name,omitempty"`
	ProfilePassword            string                   `json:"profile_password,omitempty"`
	ProfileUsername            string                   `json:"profile_username,omitempty"`
	RebootVmOnCpuChange        *bool                    `json:"reboot_vm_on_cpu_change,omitempty"`
	SaveConfig                 *bool                    `json:"save_config,omitempty"`
	State                      string                   `json:"state,omitempty"`
	StdBwAvailable             *int64                   `json:"std_bw_available,omitempty"`
	StdBwConfig                *int64                   `json:"std_bw_config,omitempty"`
	StdBwTotal                 *int64                   `json:"std_bw_total,omitempty"`
	TemplateName               string                   `json:"template_name,omitempty"`
	Throughput                 *int64                   `json:"throughput,omitempty"`
	ThroughputAllocationMode   *int64                   `json:"throughput_allocation_mode,omitempty"`
	ThroughputLimit            *int64                   `json:"throughput_limit,omitempty"`
	Type                       string                   `json:"type,omitempty"`
	Username                   string                   `json:"username,omitempty"`
	VcpuConfig                 *int64                   `json:"vcpu_config,omitempty"`
	VlanId01                   *int64                   `json:"vlan_id_0_1,omitempty"`
	VlanId02                   *int64                   `json:"vlan_id_0_2,omitempty"`
	VlanType                   *int64                   `json:"vlan_type,omitempty"`
	VmMemoryTotal              *int64                   `json:"vm_memory_total,omitempty"`
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
			"backplane": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Backplane Interface. Minimum length =  1",
				MarkdownDescription: "Backplane Interface. Minimum length =  1",
			},
			"burst_priority": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Burst Priority of the VM Instance between 1 and 4.",
				MarkdownDescription: "Burst Priority of the VM Instance between 1 and 4.",
			},
			"cmd_policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "true if you want to allow shell/sftp/scp access to NetScaler Instance administrator. Minimum length =  1 Maximum length =  1024",
				MarkdownDescription: "true if you want to allow shell/sftp/scp access to NetScaler Instance administrator. Minimum length =  1 Maximum length =  1024",
			},
			"config_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Configuration Type. Values: 0: IPv4, 1: IPv6, 2: Both.",
				MarkdownDescription: "Configuration Type. Values: 0: IPv4, 1: IPv6, 2: Both.",
			},
			"crypto_change_requires_reboot": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "`true` if the current changes made by user requires a reboot of the VM else `false`.",
				MarkdownDescription: "`true` if the current changes made by user requires a reboot of the VM else `false`.",
			},
			"customid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Custom ID.",
				MarkdownDescription: "Custom ID.",
			},
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "ID of the NetScaler ADC Instance.",
				MarkdownDescription: "ID of the NetScaler ADC Instance.",
			},
			"datacenter_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Datacenter Id is system generated key for data center.",
				MarkdownDescription: "Datacenter Id is system generated key for data center.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description of managed device. Minimum length =  1 Maximum length =  512",
				MarkdownDescription: "Description of managed device. Minimum length =  1 Maximum length =  512",
			},
			"device_family": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Device Family. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Device Family. Minimum length =  1 Maximum length =  64",
			},
			"display_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Display Name for this managed device. For HA pair it will be A-B, and for Cluster it will be CLIP. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Display Name for this managed device. For HA pair it will be A-B, and for Cluster it will be CLIP. Minimum length =  1 Maximum length =  128",
			},
			"domain_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Domain name of VM Device. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Domain name of VM Device. Minimum length =  1 Maximum length =  128",
			},
			"ent_bw_available": schema.Int64Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "Enterprise Bandwidth configured.",
				MarkdownDescription: "Enterprise Bandwidth configured.",
			},
			"ent_bw_config": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enterprise Bandwidth configured.",
				MarkdownDescription: "Enterprise Bandwidth configured.",
			},
			"ent_bw_total": schema.Int64Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "Enterprise Bandwidth Total.",
				MarkdownDescription: "Enterprise Bandwidth Total.",
			},
			"fips_partition_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "FIPS Partition Name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "FIPS Partition Name. Minimum length =  1 Maximum length =  128",
			},
			"gateway": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Default Gateway of managed device. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Default Gateway of managed device. Minimum length =  1 Maximum length =  64",
			},
			"gateway_ipv6": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Gateway IPv6 Address.",
				MarkdownDescription: "Gateway IPv6 Address.",
			},
			"host_ip_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Host IPAddress where VM is provisioned. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Host IPAddress where VM is provisioned. Minimum length =  1 Maximum length =  64",
			},
			"hostname": schema.StringAttribute{
				Optional: true,
				// Computed:            true,
				Description:         "Assign hostname to managed device, if this is not provided, name will be set as host name . Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "Assign hostname to managed device, if this is not provided, name will be set as host name . Minimum length =  1 Maximum length =  256",
			},
			"if_0_1": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Network 0/1 on VM Instance, Select this option to assign 0/1 Interface",
				MarkdownDescription: "Network 0/1 on VM Instance, Select this option to assign 0/1 Interface",
			},
			"if_0_2": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Network 0/2 on VM Instance, Select this option to assign 0/2 Interface",
				MarkdownDescription: "Network 0/2 on VM Instance, Select this option to assign 0/2 Interface",
			},
			"if_internal_ip_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Set as true if VPX is managed by internal network (not required to be set for SDWAN).",
				MarkdownDescription: "Set as true if VPX is managed by internal network (not required to be set for SDWAN).",
			},
			"image_name": schema.StringAttribute{
				Optional: true,
				// Computed:            true,
				Description:         "Image Name, This parameter is used while provisioning VM Instance with XVA image, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Image Name, This parameter is used while provisioning VM Instance with XVA image, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
			},
			"instance_mode": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Denotes state- primary,secondary,clip,clusternode.",
				MarkdownDescription: "Denotes state- primary,secondary,clip,clusternode.",
			},
			"internal_ip_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Internal IP Address for this managed device. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Internal IP Address for this managed device. Minimum length =  1 Maximum length =  64",
			},
			"ip_address": schema.StringAttribute{
				Required: true,
				// We have below code insted of ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "IP Address for this managed device. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IP Address for this managed device. Minimum length =  1 Maximum length =  64",
			},
			"ipv4_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "IPv4 Address. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "IPv4 Address. Minimum length =  1 Maximum length =  64",
			},
			"ipv6_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "IPv6 Address.",
				MarkdownDescription: "IPv6 Address.",
			},
			"is_clip": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Is Clip.",
				MarkdownDescription: "Is Clip.",
			},
			"is_managed": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Is Managed.",
				MarkdownDescription: "Is Managed.",
			},
			"is_new_crypto": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "`true` if number_of_acu/number_of_scu are used, `false` if number_of_ssl_cores is used.",
				MarkdownDescription: "`true` if number_of_acu/number_of_scu are used, `false` if number_of_ssl_cores is used.",
			},
			"iscco": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Is CCO.",
				MarkdownDescription: "Is CCO.",
			},
			"l2_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "L2mode status of VM Instance. Select this option to allow L2 mode on all the Data Interfaces on this NetScaler ADC Instance",
				MarkdownDescription: "L2mode status of VM Instance. Select this option to allow L2 mode on all the Data Interfaces on this NetScaler ADC Instance",
			},
			"la_mgmt": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Bond consisting of management ports on VM Instance. When Management Channel created for interfaces, this will be set to `true`",
				MarkdownDescription: "Bond consisting of management ports on VM Instance. When Management Channel created for interfaces, this will be set to `true`",
			},
			"last_updated_time": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Last Updated Time.",
				MarkdownDescription: "Last Updated Time.",
			},
			"license": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Feature License for NetScaler ADC Instance, needs to be set while provisioning [Possible values: Standard, Enterprise, Platinum].",
				Validators: []validator.String{
					stringvalidator.OneOf("Enterprise", "Platinum", "Standard"),
				},
			},
			"license_edition": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Edition of instance.",
				MarkdownDescription: "Edition of instance.",
			},
			"license_grace_time": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Grace for this NetScaler Instance..",
				MarkdownDescription: "Grace for this NetScaler Instance..",
			},
			"mastools_version": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Mastools version if the device is embedded agent.",
				MarkdownDescription: "Mastools version if the device is embedded agent.",
			},
			"max_burst_throughput": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Maximum burst throughput in Mbps of VM Instance.",
				MarkdownDescription: "Maximum burst throughput in Mbps of VM Instance.",
			},
			"metrics_collection": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Flag to check if metrics collection is enabled or disabled..",
				MarkdownDescription: "Flag to check if metrics collection is enabled or disabled..",
			},
			"mgmt_ip_address": schema.StringAttribute{
				Optional: true,
				// Computed:            true,
				Description:         "Management IP Address for this Managed Device. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Management IP Address for this Managed Device. Minimum length =  1 Maximum length =  64",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of managed device. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Name of managed device. Minimum length =  1 Maximum length =  128",
			},
			"netmask": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Netmask of managed device. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Netmask of managed device. Minimum length =  1 Maximum length =  64",
			},
			"network_interfaces": schema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network Interfaces.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_channel_name": schema.StringAttribute{
							Optional:            true,
							Description:         "Device channel name of the interface on the host machine.",
							MarkdownDescription: "Device channel name of the interface on the host machine.",
						},
						"gateway": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Gateway",
							MarkdownDescription: "Gateway",
						},
						// Terraform does not expect `id` to be used as an attribute name.
						// In Terraform, the `id` attribute is typically reserved for the unique identifier of a resource.
						// Using `id` as an attribute name can lead to conflicts or unexpected behavior.
						// Therefore, it is recommended to use a different attribute name to avoid such issues.
						"network_interface_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Id",
							MarkdownDescription: "Id",
						},
						"interface_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Name of this interface.",
							MarkdownDescription: "Interface Name",
						},
						"ip_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "IP Address",
							MarkdownDescription: "IP Address",
						},
						"is_member_ifc": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "`true` if this interface is member of a channel.",
							MarkdownDescription: "`true` if this interface is member of a channel.",
						},
						"is_mgmt_ifc": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "`true` if this is the management interface.",
							MarkdownDescription: "`true` if this is the management interface.",
						},
						"is_vlan_applied": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Is VLAN added on NetworkInterface of VM Instance.",
							MarkdownDescription: "Is VLAN added on NetworkInterface of VM Instance.",
						},
						"l2_enabled": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "L2 mode status of Interface.",
							MarkdownDescription: "L2 mode status of Interface.",
						},
						"mac_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "MAC Address",
							MarkdownDescription: "MAC Address",
						},
						"mac_mode": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "MAC Mode, The method according to which MAC Address is assigned to Interface. Possible values: [default, generated, custom] default: XenServer assigns a MAC Address. custom: SDX Administrator assigns a MAC address. generated: Generate a MAC address by using the base MAC address set at System Level.",
							MarkdownDescription: "MAC Mode, The method according to which MAC Address is assigned to Interface. Possible values: [default, generated, custom] default: XenServer assigns a MAC Address. custom: SDX Administrator assigns a MAC address. generated: Generate a MAC address by using the base MAC address set at System Level.",
						},
						"managed_device_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Managed Device Id",
							MarkdownDescription: "Managed Device Id",
						},
						"name_server": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Name Server",
							MarkdownDescription: "Name Server",
						},
						"netmask": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Netmask",
							MarkdownDescription: "Netmask",
						},
						"parent_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Parent Id",
							MarkdownDescription: "Parent Id",
						},
						"parent_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Parent Name",
							MarkdownDescription: "Parent Name",
						},
						"port_name": schema.StringAttribute{
							Required:            true,
							Description:         "Port name of the interface on the host machine.",
							MarkdownDescription: "Port name of the interface on the host machine.",
						},
						"receiveuntagged": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Receive Untagged Packets on Interface/Channel. Allow Untagged Traffic.",
							MarkdownDescription: "Receive Untagged Packets on Interface/Channel. Allow Untagged Traffic.",
						},
						"sdx_formation_network_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "Sdx Formation Network Id",
							MarkdownDescription: "Sdx Formation Network Id",
						},
						"vlan": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							Description:         "VLAN.",
							MarkdownDescription: "VLAN.",
						},
						"vlan_whitelist": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "VLAN Whitelist.",
							MarkdownDescription: "VLAN Whitelist.",
						},
						"vlan_whitelist_array": schema.ListAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							Computed:            true,
							Description:         "Allowed VLANs. Range of VLANs can be provided using hyphen '-' separater and separated VLANs can also be provided. (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VLANs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VLANs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
							MarkdownDescription: "Allowed VLANs. Range of VLANs can be provided using hyphen '-' separater and separated VLANs can also be provided. (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VLANs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VLANs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
						},
						"vrid_list_ipv4": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "VRID List for Interface/Channel for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., \"100-110,142,143,151-155\").",
							MarkdownDescription: "VRID List for Interface/Channel for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., \"100-110,142,143,151-155\").",
						},
						"vrid_list_ipv4_array": schema.ListAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							Computed:            true,
							Description:         "VRID List for Interface for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
							MarkdownDescription: "VRID List for Interface for IPV4 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
						},
						"vrid_list_ipv6": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "VRID List for Interface/Channel for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., \"100-110,142,143,151-155\").",
							MarkdownDescription: "VRID List for Interface/Channel for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and multiple comma ',' separated VRIDs can also be provided, (e.g., \"100-110,142,143,151-155\").",
						},
						"vrid_list_ipv6_array": schema.ListAttribute{
							ElementType:         types.StringType,
							Optional:            true,
							Computed:            true,
							Description:         "VRID List for Interface for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
							MarkdownDescription: "VRID List for Interface for IPV6 VMAC Generation. Range of VRIDs can be provided using hyphen '-' separater and separated VRIDs can also be provided, (e.g., [\"100-110\",\"142\",\"143\",\"151-155\"]). Providing in the suggested format is necessary as SDX internally try to convert in this form, so due to that, if the format provided is wrong it may cause error from terraform. To list individual VRIDs if they are not in sequence or if the sequence is 2 or fewer (e.g., [\"100\",\"101\",\"4000\",\"4001\"]). If the VRIDs are in sequence of 3 or more, use the range format with hypen '-' seperated like (e.g., [\"100-103\",\"4000-4002\"]). Also, maintain the order as well (Ascending order) (e.g., [\"100-103\",\"200\",\"4000-4002\"])",
						},
					},
				},
			},
			"nexthop": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Next Hop IP address. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Next Hop IP address. Minimum length =  1 Maximum length =  64",
			},
			"nexthop_v6": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Next Hop IPv6 Address.",
				MarkdownDescription: "Next Hop IPv6 Address.",
			},
			"node_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Node identification of a device.",
				MarkdownDescription: "Node identification of a device.",
			},
			"ns_ip_address": schema.StringAttribute{
				// Computed:            true,
				Optional:            true,
				Description:         "NetScaler IP Address for this managed device. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "NetScaler IP Address for this managed device. Minimum length =  1 Maximum length =  128",
			},
			"nsvlan_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "VLAN for Management Traffic.",
				MarkdownDescription: "VLAN for Management Traffic.",
			},
			"nsvlan_interfaces": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "VLAN Interfaces. Minimum length =  1 Maximum length =  50",
				MarkdownDescription: "VLAN Interfaces. Minimum length =  1 Maximum length =  50",
			},
			"nsvlan_tagged": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "When this option is selected, selected interfaces are added as tagged members of Management VLAN",
				MarkdownDescription: "When this option is selected, selected interfaces are added as tagged members of Management VLAN",
			},
			"num_pes": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Total number of PEs.",
				MarkdownDescription: "Total number of PEs.",
			},
			"number_of_acu": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Assign number of asymmetric crypto units to VM Instance.",
				MarkdownDescription: "Assign number of asymmetric crypto units to VM Instance.",
			},
			"number_of_cores": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Number of cores that are assigned to VM Instance.",
				MarkdownDescription: "Number of cores that are assigned to VM Instance.",
			},
			"number_of_scu": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Assign number of asymmetric crypto units to VM Instance.",
				MarkdownDescription: "Assign number of asymmetric crypto units to VM Instance.",
			},
			"number_of_ssl_cards": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Number of SSL Cards.",
				MarkdownDescription: "Number of SSL Cards.",
			},
			"number_of_ssl_cores": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Assign number of ssl virtual functions to VM Instance.",
				MarkdownDescription: "Assign number of ssl virtual functions to VM Instance.",
			},
			"number_of_ssl_cores_up": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Number of SSL Cores Up.",
				MarkdownDescription: "Number of SSL Cores Up.",
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Description:         "Password for specified user on NetScaler Instance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Password for specified user on NetScaler Instance. Minimum length =  1 Maximum length =  127",
			},
			"plt_bw_available": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Platinum Bandwidth Available.",
				MarkdownDescription: "Platinum Bandwidth Available.",
			},
			"plt_bw_config": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Platinum Bandwidth configured.",
				MarkdownDescription: "Platinum Bandwidth configured.",
			},
			"plt_bw_total": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Total Platinum Bandwidth.",
				MarkdownDescription: "Total Platinum Bandwidth.",
			},
			"plugin_ip_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Signaling IP Address. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Signaling IP Address. Minimum length =  1 Maximum length =  64",
			},
			"plugin_netmask": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Signaling Netmask. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Signaling Netmask. Minimum length =  1 Maximum length =  64",
			},
			"pps": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Assign packets per seconds to NetScaler Instance.",
				MarkdownDescription: "Assign packets per seconds to NetScaler Instance.",
			},
			"profile_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Device Profile Name that is attached with this managed device. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Device Profile Name that is attached with this managed device. Minimum length =  1 Maximum length =  128",
			},
			"profile_password": schema.StringAttribute{
				Optional:            true,
				Description:         "Password specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Password specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128",
			},
			"profile_username": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User Name specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "User Name specified by the user for this NetScaler Instance.. Minimum length =  1 Maximum length =  128",
			},
			"reboot_vm_on_cpu_change": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Reboot VMs on CPU change during resource allocation.",
				MarkdownDescription: "Reboot VMs on CPU change during resource allocation.",
			},
			"save_config": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Should config be saved first in case instance is rebooted while modify.",
				MarkdownDescription: "Should config be saved first in case instance is rebooted while modify.",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Node State. Minimum length =  1 Maximum length =  32",
				MarkdownDescription: "Node State. Minimum length =  1 Maximum length =  32",
			},
			"std_bw_available": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Standard Bandwidth Available.",
				MarkdownDescription: "Standard Bandwidth Available.",
			},
			"std_bw_config": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Standard Bandwidth running.",
				MarkdownDescription: "Standard Bandwidth running.",
			},
			"std_bw_total": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Standard Bandwidth.",
				MarkdownDescription: "Standard Bandwidth.",
			},
			"template_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Template Name, This parameter is used while provisioning VM Instance with template, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Template Name, This parameter is used while provisioning VM Instance with template, template_name is given priority if provided along with image_name. Minimum length =  1 Maximum length =  128",
			},
			"throughput": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Assign throughput in Mbps to VM Instance.",
				MarkdownDescription: "Assign throughput in Mbps to VM Instance.",
			},
			"throughput_allocation_mode": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Throughput Allocation Mode: 0-Fixed, 1-Burst-able.",
				MarkdownDescription: "Throughput Allocation Mode: 0-Fixed, 1-Burst-able.",
			},
			"throughput_limit": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Throughput Limit in Mbps set for VM Instance.",
				MarkdownDescription: "Throughput Limit in Mbps set for VM Instance.",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Type of device, (Xen | NS). Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Type of device, (Xen | NS). Minimum length =  1 Maximum length =  64",
			},
			"username": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User Name (except nsroot) to be configured on NetScaler Instance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "User Name (except nsroot) to be configured on NetScaler Instance. Minimum length =  1 Maximum length =  127",
			},
			"vcpu_config": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Number of vCPU allocated for the device.",
				MarkdownDescription: "Number of vCPU allocated for the device.",
			},
			"vlan_id_0_1": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "VLAN id for the management interface 0/1. This VLAN ID is used to filter management traffic on 0/1 at hypervisor layer.",
				MarkdownDescription: "VLAN id for the management interface 0/1. This VLAN ID is used to filter management traffic on 0/1 at hypervisor layer.",
			},
			"vlan_id_0_2": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "VLAN id for the management interface 0/2. This VLAN ID is used to filter management traffic on 0/2 at hypervisor layer.",
				MarkdownDescription: "VLAN id for the management interface 0/2. This VLAN ID is used to filter management traffic on 0/2 at hypervisor layer.",
			},
			"vlan_type": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "VLAN Type, NetScaler or L2 VLAN. Select 0 for NetScaler VLAN or 1 for L2 VLAN.",
				MarkdownDescription: "VLAN Type, NetScaler or L2 VLAN. Select 0 for NetScaler VLAN or 1 for L2 VLAN.",
			},
			"vm_memory_total": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Total Memory of VM Instance in MB. 2048MB, 5120MB.",
				MarkdownDescription: "Total Memory of VM Instance in MB. 2048MB, 5120MB.",
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

	provisionVpxReq := nsGetThePayloadFromtheConfig(ctx, data)

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

func nsGetThePayloadFromtheConfig(ctx context.Context, data *provisionVpxResourceModel) provisionVpxResourceReq {
	tflog.Debug(ctx, "In nsGetThePayloadFromtheConfig Function")

	nsReqPayload := provisionVpxResourceReq{
		Backplane:                  data.Backplane.ValueString(),
		BurstPriority:              utils.ToIntValue(data.BurstPriority),
		CmdPolicy:                  data.CmdPolicy.ValueString(),
		ConfigType:                 utils.ToIntValue(data.ConfigType),
		CryptoChangeRequiresReboot: utils.ToBoolValue(data.CryptoChangeRequiresReboot),
		Customid:                   data.Customid.ValueString(),
		DatacenterId:               data.DatacenterId.ValueString(),
		Description:                data.Description.ValueString(),
		DeviceFamily:               data.DeviceFamily.ValueString(),
		DisplayName:                data.DisplayName.ValueString(),
		DomainName:                 data.DomainName.ValueString(),
		EntBwAvailable:             utils.ToIntValue(data.EntBwAvailable),
		EntBwConfig:                utils.ToIntValue(data.EntBwConfig),
		EntBwTotal:                 utils.ToIntValue(data.EntBwTotal),
		FipsPartitionName:          data.FipsPartitionName.ValueString(),
		Gateway:                    data.Gateway.ValueString(),
		GatewayIpv6:                data.GatewayIpv6.ValueString(),
		HostIpAddress:              data.HostIpAddress.ValueString(),
		Hostname:                   data.Hostname.ValueString(),
		If01:                       utils.ToBoolValue(data.If01),
		If02:                       utils.ToBoolValue(data.If02),
		IfInternalIpEnabled:        utils.ToBoolValue(data.IfInternalIpEnabled),
		ImageName:                  data.ImageName.ValueString(),
		InstanceMode:               data.InstanceMode.ValueString(),
		InternalIpAddress:          data.InternalIpAddress.ValueString(),
		IpAddress:                  data.IpAddress.ValueString(),
		Ipv4Address:                data.Ipv4Address.ValueString(),
		Ipv6Address:                data.Ipv6Address.ValueString(),
		IsClip:                     utils.ToBoolValue(data.IsClip),
		IsManaged:                  utils.ToBoolValue(data.IsManaged),
		IsNewCrypto:                utils.ToBoolValue(data.IsNewCrypto),
		Iscco:                      utils.ToBoolValue(data.Iscco),
		L2Enabled:                  utils.ToBoolValue(data.L2Enabled),
		LaMgmt:                     utils.ToBoolValue(data.LaMgmt),
		LastUpdatedTime:            utils.ToIntValue(data.LastUpdatedTime),
		License:                    data.License.ValueString(),
		LicenseEdition:             data.LicenseEdition.ValueString(),
		LicenseGraceTime:           utils.ToIntValue(data.LicenseGraceTime),
		MastoolsVersion:            data.MastoolsVersion.ValueString(),
		MaxBurstThroughput:         utils.ToIntValue(data.MaxBurstThroughput),
		MetricsCollection:          utils.ToBoolValue(data.MetricsCollection),
		MgmtIpAddress:              data.MgmtIpAddress.ValueString(),
		Name:                       data.Name.ValueString(),
		Netmask:                    data.Netmask.ValueString(),
		NetworkInterfaces:          networkInterfaceFromConfigToSchema(ctx, data.NetworkInterfaces),
		Nexthop:                    data.Nexthop.ValueString(),
		NexthopV6:                  data.NexthopV6.ValueString(),
		NodeId:                     data.NodeId.ValueString(),
		NsIpAddress:                data.NsIpAddress.ValueString(),
		NsvlanId:                   utils.ToIntValue(data.NsvlanId),
		NsvlanInterfaces:           utils.TypeListToUnmarshalStringList(data.NsvlanInterfaces),
		NsvlanTagged:               utils.ToBoolValue(data.NsvlanTagged),
		NumPes:                     utils.ToIntValue(data.NumPes),
		NumberOfAcu:                utils.ToIntValue(data.NumberOfAcu),
		NumberOfCores:              utils.ToIntValue(data.NumberOfCores),
		NumberOfScu:                utils.ToIntValue(data.NumberOfScu),
		NumberOfSslCards:           utils.ToIntValue(data.NumberOfSslCards),
		NumberOfSslCores:           utils.ToIntValue(data.NumberOfSslCores),
		NumberOfSslCoresUp:         utils.ToIntValue(data.NumberOfSslCoresUp),
		Password:                   data.Password.ValueString(),
		PltBwAvailable:             utils.ToIntValue(data.PltBwAvailable),
		PltBwConfig:                utils.ToIntValue(data.PltBwConfig),
		PltBwTotal:                 utils.ToIntValue(data.PltBwTotal),
		PluginIpAddress:            data.PluginIpAddress.ValueString(),
		PluginNetmask:              data.PluginNetmask.ValueString(),
		Pps:                        utils.ToIntValue(data.Pps),
		ProfileName:                data.ProfileName.ValueString(),
		ProfilePassword:            data.ProfilePassword.ValueString(),
		ProfileUsername:            data.ProfileUsername.ValueString(),
		RebootVmOnCpuChange:        utils.ToBoolValue(data.RebootVmOnCpuChange),
		SaveConfig:                 utils.ToBoolValue(data.SaveConfig),
		State:                      data.State.ValueString(),
		StdBwAvailable:             utils.ToIntValue(data.StdBwAvailable),
		StdBwConfig:                utils.ToIntValue(data.StdBwConfig),
		StdBwTotal:                 utils.ToIntValue(data.StdBwTotal),
		TemplateName:               data.TemplateName.ValueString(),
		Throughput:                 utils.ToIntValue(data.Throughput),
		ThroughputAllocationMode:   utils.ToIntValue(data.ThroughputAllocationMode),
		ThroughputLimit:            utils.ToIntValue(data.ThroughputLimit),
		Type:                       data.Type.ValueString(),
		Username:                   data.Username.ValueString(),
		VcpuConfig:                 utils.ToIntValue(data.VcpuConfig),
		VlanId01:                   utils.ToIntValue(data.VlanId01),
		VlanId02:                   utils.ToIntValue(data.VlanId02),
		VlanType:                   utils.ToIntValue(data.VlanType),
		VmMemoryTotal:              utils.ToIntValue(data.VmMemoryTotal),
	}

	return nsReqPayload
}

func networkInterfaceFromConfigToSchema(ctx context.Context, nidata basetypes.ListValue) []map[string]interface{} {
	tflog.Debug(ctx, "In networkInterfaceFromConfigToSchema func of provisionVpxResource")

	var networkInterfacesMapList []map[string]interface{}
	// Get the network interfaces from the model
	for _, internalObjectAttr := range nidata.Elements() {
		internalObjectMap := internalObjectAttr.(basetypes.ObjectValue)

		mapNi := make(map[string]interface{}, 0)

		for key, val := range internalObjectMap.Attributes() {
			if !val.IsNull() && !val.IsUnknown() {
				switch val.Type(ctx) {
				case types.StringType:
					if key == "network_interface_id" {
						mapNi["id"] = val.(basetypes.StringValue).ValueString()
					} else {
						mapNi[key] = val.(basetypes.StringValue).ValueString()
					}
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

	var data *provisionVpxResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	endpoint := "ns"

	dataArr, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Resource",
			fmt.Sprintf("Error reading resource: %s", err.Error()),
		)
		return
	}

	getResponseData := dataArr[endpoint].([]interface{})[0].(map[string]interface{})

	data.Backplane = utils.StringValueToFramework(getResponseData["backplane"])
	data.BurstPriority = utils.Int64ValueToFramework(getResponseData["burst_priority"])
	data.CmdPolicy = utils.StringValueToFramework(getResponseData["cmd_policy"])
	data.ConfigType = utils.Int64ValueToFramework(getResponseData["config_type"])
	data.CryptoChangeRequiresReboot = utils.BoolValueToFramework(getResponseData["crypto_change_requires_reboot"])
	data.Customid = utils.StringValueToFramework(getResponseData["customid"])
	data.DatacenterId = utils.StringValueToFramework(getResponseData["datacenter_id"])
	data.Description = utils.StringValueToFramework(getResponseData["description"])
	data.DeviceFamily = utils.StringValueToFramework(getResponseData["device_family"])
	data.DisplayName = utils.StringValueToFramework(getResponseData["display_name"])
	data.DomainName = utils.StringValueToFramework(getResponseData["domain_name"])
	data.EntBwAvailable = utils.Int64ValueToFramework(getResponseData["ent_bw_available"])
	data.EntBwConfig = utils.Int64ValueToFramework(getResponseData["ent_bw_config"])
	data.EntBwTotal = utils.Int64ValueToFramework(getResponseData["ent_bw_total"])
	data.FipsPartitionName = utils.StringValueToFramework(getResponseData["fips_partition_name"])
	data.Gateway = utils.StringValueToFramework(getResponseData["gateway"])
	data.GatewayIpv6 = utils.StringValueToFramework(getResponseData["gateway_ipv6"])
	data.HostIpAddress = utils.StringValueToFramework(getResponseData["host_ip_address"])
	// data.Hostname = utils.StringValueToFramework(getResponseData["hostname"])
	data.If01 = utils.BoolValueToFramework(getResponseData["if_0_1"])
	data.If02 = utils.BoolValueToFramework(getResponseData["if_0_2"])
	data.IfInternalIpEnabled = utils.BoolValueToFramework(getResponseData["if_internal_ip_enabled"])
	// data.ImageName = utils.StringValueToFramework(getResponseData["image_name"])
	data.InstanceMode = utils.StringValueToFramework(getResponseData["instance_mode"])
	data.InternalIpAddress = utils.StringValueToFramework(getResponseData["internal_ip_address"])
	data.IpAddress = utils.StringValueToFramework(getResponseData["ip_address"])
	data.Ipv4Address = utils.StringValueToFramework(getResponseData["ipv4_address"])
	data.Ipv6Address = utils.StringValueToFramework(getResponseData["ipv6_address"])
	data.IsClip = utils.BoolValueToFramework(getResponseData["is_clip"])
	data.IsManaged = utils.BoolValueToFramework(getResponseData["is_managed"])
	data.IsNewCrypto = utils.BoolValueToFramework(getResponseData["is_new_crypto"])
	data.Iscco = utils.BoolValueToFramework(getResponseData["iscco"])
	data.L2Enabled = utils.BoolValueToFramework(getResponseData["l2_enabled"])
	data.LaMgmt = utils.BoolValueToFramework(getResponseData["la_mgmt"])
	data.LastUpdatedTime = utils.Int64ValueToFramework(getResponseData["last_updated_time"])
	data.License = utils.StringValueToFramework(getResponseData["license"])
	data.LicenseEdition = utils.StringValueToFramework(getResponseData["license_edition"])
	data.LicenseGraceTime = utils.Int64ValueToFramework(getResponseData["license_grace_time"])
	data.MastoolsVersion = utils.StringValueToFramework(getResponseData["mastools_version"])
	data.MaxBurstThroughput = utils.Int64ValueToFramework(getResponseData["max_burst_throughput"])
	data.MetricsCollection = utils.BoolValueToFramework(getResponseData["metrics_collection"])
	// data.MgmtIpAddress = utils.StringValueToFramework(getResponseData["mgmt_ip_address"])
	data.Name = utils.StringValueToFramework(getResponseData["name"])
	data.Netmask = utils.StringValueToFramework(getResponseData["netmask"])
	data.NetworkInterfaces = networkinterfaceTonetworkinterfaceTF(getResponseData["network_interfaces"].([]interface{}), data.NetworkInterfaces, ctx)
	data.Nexthop = utils.StringValueToFramework(getResponseData["nexthop"])
	data.NexthopV6 = utils.StringValueToFramework(getResponseData["nexthop_v6"])
	data.NodeId = utils.StringValueToFramework(getResponseData["node_id"])
	// data.NsIpAddress = utils.StringValueToFramework(getResponseData["ns_ip_address"])
	data.NsvlanId = utils.Int64ValueToFramework(getResponseData["nsvlan_id"])
	data.NsvlanInterfaces = utils.StringListToTypeList(utils.ToStringList(getResponseData["nsvlan_interfaces"].([]interface{})))
	data.NsvlanTagged = utils.BoolValueToFramework(getResponseData["nsvlan_tagged"])
	data.NumPes = utils.Int64ValueToFramework(getResponseData["num_pes"])
	data.NumberOfAcu = utils.Int64ValueToFramework(getResponseData["number_of_acu"])
	data.NumberOfCores = utils.Int64ValueToFramework(getResponseData["number_of_cores"])
	data.NumberOfScu = utils.Int64ValueToFramework(getResponseData["number_of_scu"])
	data.NumberOfSslCards = utils.Int64ValueToFramework(getResponseData["number_of_ssl_cards"])
	data.NumberOfSslCores = utils.Int64ValueToFramework(getResponseData["number_of_ssl_cores"])
	data.NumberOfSslCoresUp = utils.Int64ValueToFramework(getResponseData["number_of_ssl_cores_up"])
	// data.Password = utils.StringValueToFramework(getResponseData["password"])
	data.PltBwAvailable = utils.Int64ValueToFramework(getResponseData["plt_bw_available"])
	data.PltBwConfig = utils.Int64ValueToFramework(getResponseData["plt_bw_config"])
	data.PltBwTotal = utils.Int64ValueToFramework(getResponseData["plt_bw_total"])
	data.PluginIpAddress = utils.StringValueToFramework(getResponseData["plugin_ip_address"])
	data.PluginNetmask = utils.StringValueToFramework(getResponseData["plugin_netmask"])
	data.Pps = utils.Int64ValueToFramework(getResponseData["pps"])
	data.ProfileName = utils.StringValueToFramework(getResponseData["profile_name"])
	// data.ProfilePassword = utils.StringValueToFramework(getResponseData["profile_password"])
	data.ProfileUsername = utils.StringValueToFramework(getResponseData["profile_username"])
	data.RebootVmOnCpuChange = utils.BoolValueToFramework(getResponseData["reboot_vm_on_cpu_change"])
	data.SaveConfig = utils.BoolValueToFramework(getResponseData["save_config"])
	data.State = utils.StringValueToFramework(getResponseData["state"])
	data.StdBwAvailable = utils.Int64ValueToFramework(getResponseData["std_bw_available"])
	data.StdBwConfig = utils.Int64ValueToFramework(getResponseData["std_bw_config"])
	data.StdBwTotal = utils.Int64ValueToFramework(getResponseData["std_bw_total"])
	data.TemplateName = utils.StringValueToFramework(getResponseData["template_name"])
	data.Throughput = utils.Int64ValueToFramework(getResponseData["throughput"])
	data.ThroughputAllocationMode = utils.Int64ValueToFramework(getResponseData["throughput_allocation_mode"])
	data.ThroughputLimit = utils.Int64ValueToFramework(getResponseData["throughput_limit"])
	data.Type = utils.StringValueToFramework(getResponseData["type"])
	data.Username = utils.StringValueToFramework(getResponseData["username"])
	data.VcpuConfig = utils.Int64ValueToFramework(getResponseData["vcpu_config"])
	data.VlanId01 = utils.Int64ValueToFramework(getResponseData["vlan_id_0_1"])
	data.VlanId02 = utils.Int64ValueToFramework(getResponseData["vlan_id_0_2"])
	data.VlanType = utils.Int64ValueToFramework(getResponseData["vlan_type"])
	data.VmMemoryTotal = utils.Int64ValueToFramework(getResponseData["vm_memory_total"])

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func networkinterfaceTonetworkinterfaceTF(nifGetResponse []interface{}, stateNetworkInterface basetypes.ListValue, ctx context.Context) basetypes.ListValue {
	var nifSchemaAttributes = []string{
		"port_name",
		"device_channel_name",
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

	var nifListFilterFromGetResponse []map[string]interface{}

	for _, v := range nifGetResponse {

		nifMapGetResponse := v.(map[string]interface{})
		nifMapFilterFromGetResponse := make(map[string]interface{})
		if nifMapGetResponse["port_name"].(string) != "" && nifMapGetResponse["parent_channel_id"].(string) == "" {
			for key, inputNif := range nifMapGetResponse {
				if service.Contains(nifSchemaAttributes, key) {
					nifMapFilterFromGetResponse[key] = inputNif
				}
			}

			nifListFilterFromGetResponse = append(nifListFilterFromGetResponse, nifMapFilterFromGetResponse)
		}
	}

	var nifList []attr.Value
	for _, nif := range nifListFilterFromGetResponse {

		var nifObjectValue basetypes.ObjectValue
		nifmap := make(map[string]attr.Value)
		nifMapTypes := make(map[string]attr.Type)

		attributeTypes := stateNetworkInterface.ElementType(ctx).(types.ObjectType).AttributeTypes()

		for k, v := range nif {

			switch attributeTypes[k].(type) {
			case basetypes.BoolType:
				nifmap[k] = utils.BoolValueToFramework(v)
				nifMapTypes[k] = types.BoolType
			case basetypes.Int64Type:
				nifmap[k] = utils.Int64ValueToFramework(v)
				nifMapTypes[k] = types.Int64Type
			case basetypes.ListType:
				if v == nil {
					nifmap[k] = basetypes.NewListNull(types.StringType)
					nifMapTypes[k] = types.ListType{ElemType: types.StringType}
				} else if len(v.([]interface{})) == 0 {
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
			default:
				if k == "id" {
					nifmap["network_interface_id"] = utils.StringValueToFramework(v)
					nifMapTypes["network_interface_id"] = types.StringType
				} else {
					nifmap[k] = utils.StringValueToFramework(v)
					nifMapTypes[k] = types.StringType
				}
			}
		}

		nifObjectValue, _ = basetypes.NewObjectValue(nifMapTypes, nifmap)
		nifList = append(nifList, nifObjectValue)

	}
	newNifList, _ := basetypes.NewListValueFrom(ctx, stateNetworkInterface.ElementType(ctx), nifList)

	return newNifList

}

func (r *provisionVpxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] In Update Method of provisionVpxResource with Id: %s", resId))

	var data provisionVpxResourceModel
	diags := req.Plan.Get(ctx, &data)
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
	requestPayload := nsGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, resourceId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Resource",
			fmt.Sprintf("Error updating resource: %s", err.Error()),
		)
		return
	}

	diags = resp.State.Set(ctx, &data)
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

	*resp = resource.UpdateResponse{
		State:       rresp.State,
		Diagnostics: rresp.Diagnostics,
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
