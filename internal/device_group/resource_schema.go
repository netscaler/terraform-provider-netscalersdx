package device_group

import (
	"context"

	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func deviceGroupResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Configuration for Device Group resource.",
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				Optional:            true,
				Description:         "Device group category. Will be default/upgrade.. Maximum length =  255",
				MarkdownDescription: "Device group category. Will be default/upgrade.. Maximum length =  255",
			},
			"criteria_condn": schema.StringAttribute{
				Optional:            true,
				Description:         "Tenant. Maximum length =  255",
				MarkdownDescription: "Tenant. Maximum length =  255",
			},
			"criteria_type": schema.StringAttribute{
				Optional:            true,
				Description:         "Device Group Criteria. Maximum length =  255",
				MarkdownDescription: "Device Group Criteria. Maximum length =  255",
			},
			"criteria_value": schema.StringAttribute{
				Optional:            true,
				Description:         "Criteria Value. Maximum length =  255",
				MarkdownDescription: "Criteria Value. Maximum length =  255",
			},
			"device_family": schema.StringAttribute{
				Optional:            true,
				Description:         "Device Family. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Device Family. Minimum length =  1 Maximum length =  64",
			},
			"disable_upgrade": schema.BoolAttribute{
				Optional:            true,
				Description:         "Setting to disable agent upgrades.",
				MarkdownDescription: "Setting to disable agent upgrades.",
			},
			"duration": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description:         "Duration of Maintenance window for groups of category upgrade. Minimum value =  1 Maximum value =  ",
				MarkdownDescription: "Duration of Maintenance window for groups of category upgrade. Minimum value =  1 Maximum value =  ",
			},
			"lock_acquire_time": schema.StringAttribute{
				Optional:            true,
				Description:         "Upgrade Lock acquiring time. Maximum length =  255",
				MarkdownDescription: "Upgrade Lock acquiring time. Maximum length =  255",
			},
			"lock_acquiring_device": schema.StringAttribute{
				Optional:            true,
				Description:         "Upgrade Lock acquiring device. Maximum length =  255",
				MarkdownDescription: "Upgrade Lock acquiring device. Maximum length =  255",
			},
			"maintenance_window_start": schema.StringAttribute{
				Optional:            true,
				Description:         "Maintenance window start time for groups of category upgrade. Maximum length =  255",
				MarkdownDescription: "Maintenance window start time for groups of category upgrade. Maximum length =  255",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Device Group Name. Maximum length =  255",
				MarkdownDescription: "Device Group Name. Maximum length =  255",
			},
			"static_device_list": schema.StringAttribute{
				Optional:            true,
				Description:         "Devices in the group.",
				MarkdownDescription: "Devices in the group.",
			},
			"static_device_list_arr": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Static Device List.",
				MarkdownDescription: "Static Device List.",
			},
			"upgrade_lock": schema.BoolAttribute{
				Optional:            true,
				Description:         "Lock to be acquired before upgrading device group.",
				MarkdownDescription: "Lock to be acquired before upgrading device group.",
			},
			"upgrade_version": schema.StringAttribute{
				Optional:            true,
				Description:         "New Available upgrade version for devices. Maximum length =  255",
				MarkdownDescription: "New Available upgrade version for devices. Maximum length =  255",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type deviceGroupModel struct {
	Category               types.String `tfsdk:"category"`
	CriteriaCondn          types.String `tfsdk:"criteria_condn"`
	CriteriaType           types.String `tfsdk:"criteria_type"`
	CriteriaValue          types.String `tfsdk:"criteria_value"`
	DeviceFamily           types.String `tfsdk:"device_family"`
	DisableUpgrade         types.Bool   `tfsdk:"disable_upgrade"`
	Duration               types.Int64  `tfsdk:"duration"`
	LockAcquireTime        types.String `tfsdk:"lock_acquire_time"`
	LockAcquiringDevice    types.String `tfsdk:"lock_acquiring_device"`
	MaintenanceWindowStart types.String `tfsdk:"maintenance_window_start"`
	Name                   types.String `tfsdk:"name"`
	StaticDeviceList       types.String `tfsdk:"static_device_list"`
	StaticDeviceListArr    types.List   `tfsdk:"static_device_list_arr"`
	UpgradeLock            types.Bool   `tfsdk:"upgrade_lock"`
	UpgradeVersion         types.String `tfsdk:"upgrade_version"`
	Id                     types.String `tfsdk:"id"`
}

func deviceGroupGetThePayloadFromtheConfig(ctx context.Context, data *deviceGroupModel) deviceGroupReq {
	tflog.Debug(ctx, "In deviceGroupGetThePayloadFromtheConfig Function")
	deviceGroupReqPayload := deviceGroupReq{
		Category:               data.Category.ValueString(),
		CriteriaCondn:          data.CriteriaCondn.ValueString(),
		CriteriaType:           data.CriteriaType.ValueString(),
		CriteriaValue:          data.CriteriaValue.ValueString(),
		DeviceFamily:           data.DeviceFamily.ValueString(),
		DisableUpgrade:         data.DisableUpgrade.ValueBool(),
		Duration:               data.Duration.ValueInt64(),
		LockAcquireTime:        data.LockAcquireTime.ValueString(),
		LockAcquiringDevice:    data.LockAcquiringDevice.ValueString(),
		MaintenanceWindowStart: data.MaintenanceWindowStart.ValueString(),
		Name:                   data.Name.ValueString(),
		StaticDeviceList:       data.StaticDeviceList.ValueString(),
		StaticDeviceListArr:    utils.TypeListToUnmarshalStringList(data.StaticDeviceListArr),
		UpgradeLock:            data.UpgradeLock.ValueBool(),
		UpgradeVersion:         data.UpgradeVersion.ValueString(),
	}
	return deviceGroupReqPayload
}

type deviceGroupReq struct {
	Category               string   `json:"category,omitempty"`
	CriteriaCondn          string   `json:"criteria_condn,omitempty"`
	CriteriaType           string   `json:"criteria_type,omitempty"`
	CriteriaValue          string   `json:"criteria_value,omitempty"`
	DeviceFamily           string   `json:"device_family,omitempty"`
	DisableUpgrade         bool     `json:"disable_upgrade,omitempty"`
	Duration               int64    `json:"duration,omitempty"`
	LockAcquireTime        string   `json:"lock_acquire_time,omitempty"`
	LockAcquiringDevice    string   `json:"lock_acquiring_device,omitempty"`
	MaintenanceWindowStart string   `json:"maintenance_window_start,omitempty"`
	Name                   string   `json:"name,omitempty"`
	StaticDeviceList       string   `json:"static_device_list,omitempty"`
	StaticDeviceListArr    []string `json:"static_device_list_arr,omitempty"`
	UpgradeLock            bool     `json:"upgrade_lock,omitempty"`
	UpgradeVersion         string   `json:"upgrade_version,omitempty"`
}
