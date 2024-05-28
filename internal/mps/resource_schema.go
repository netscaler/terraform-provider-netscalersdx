package mps

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func mpsResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for MPS System Status resource.",
		Attributes: map[string]schema.Attribute{
			"config_motd": schema.BoolAttribute{
				Optional:            true,
				Description:         "Configure Message of the Day (contents of motd file), this needs to be set true if user wants to configure message if the day.",
				MarkdownDescription: "Configure Message of the Day (contents of motd file), this needs to be set true if user wants to configure message if the day.",
			},
			"deployment_type": schema.StringAttribute{
				Optional:            true,
				Description:         "Indicates the type of deployment of NetScaler ADM: standalone or ha or scaleout..",
				MarkdownDescription: "Indicates the type of deployment of NetScaler ADM: standalone or ha or scaleout..",
			},
			"hist_mig_inprog": schema.BoolAttribute{
				Optional:            true,
				Description:         "Indicates whether the historical tables database migration is in progress or not..",
				MarkdownDescription: "Indicates whether the historical tables database migration is in progress or not..",
			},
			"is_cloud": schema.BoolAttribute{
				Optional:            true,
				Description:         "True if its a cloud deployment.",
				MarkdownDescription: "True if its a cloud deployment.",
			},
			"is_container": schema.BoolAttribute{
				Optional:            true,
				Description:         "True if its a container deployment.",
				MarkdownDescription: "True if its a container deployment.",
			},
			"is_member_of_default_group": schema.BoolAttribute{
				Optional:            true,
				Description:         "Is Member Of Default Group.",
				MarkdownDescription: "Is Member Of Default Group.",
			},
			"is_passive": schema.BoolAttribute{
				Optional:            true,
				Description:         "Indicates the node's state: ACTIVE or PASSIVE in a HA deployment..",
				MarkdownDescription: "Indicates the node's state: ACTIVE or PASSIVE in a HA deployment..",
			},
			"is_thirdparty_vm_supported": schema.BoolAttribute{
				Optional:            true,
				Description:         "True, if third party VM is supported.",
				MarkdownDescription: "True, if third party VM is supported.",
			},
			"maps_apikey": schema.StringAttribute{
				Optional:            true,
				Description:         "Maps API Key.",
				MarkdownDescription: "Maps API Key.",
			},
			"motd": schema.StringAttribute{
				Optional:            true,
				Description:         "Message of the Day (contents of motd file) that can be shown on UI after successful login. Maximum length =  255",
				MarkdownDescription: "Message of the Day (contents of motd file) that can be shown on UI after successful login. Maximum length =  255",
			},
			"privilege_feature": schema.StringAttribute{
				Optional:            true,
				Description:         "privilege_feature.",
				MarkdownDescription: "privilege_feature.",
			},
			"upgrade_agent_version": schema.StringAttribute{
				Optional:            true,
				Description:         "Indicates the next version the agent needs to upgrade to..",
				MarkdownDescription: "Indicates the next version the agent needs to upgrade to..",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type mpsModel struct {
	ConfigMotd              types.Bool   `tfsdk:"config_motd"`
	DeploymentType          types.String `tfsdk:"deployment_type"`
	HistMigInprog           types.Bool   `tfsdk:"hist_mig_inprog"`
	IsCloud                 types.Bool   `tfsdk:"is_cloud"`
	IsContainer             types.Bool   `tfsdk:"is_container"`
	IsMemberOfDefaultGroup  types.Bool   `tfsdk:"is_member_of_default_group"`
	IsPassive               types.Bool   `tfsdk:"is_passive"`
	IsThirdpartyVmSupported types.Bool   `tfsdk:"is_thirdparty_vm_supported"`
	MapsApikey              types.String `tfsdk:"maps_apikey"`
	Motd                    types.String `tfsdk:"motd"`
	PrivilegeFeature        types.String `tfsdk:"privilege_feature"`
	UpgradeAgentVersion     types.String `tfsdk:"upgrade_agent_version"`
	Id                      types.String `tfsdk:"id"`
}

func mpsGetThePayloadFromtheConfig(ctx context.Context, data *mpsModel) mpsReq {
	tflog.Debug(ctx, "In mpsGetThePayloadFromtheConfig Function")
	mpsReqPayload := mpsReq{
		ConfigMotd:              data.ConfigMotd.ValueBoolPointer(),
		DeploymentType:          data.DeploymentType.ValueString(),
		HistMigInprog:           data.HistMigInprog.ValueBoolPointer(),
		IsCloud:                 data.IsCloud.ValueBoolPointer(),
		IsContainer:             data.IsContainer.ValueBoolPointer(),
		IsMemberOfDefaultGroup:  data.IsMemberOfDefaultGroup.ValueBoolPointer(),
		IsPassive:               data.IsPassive.ValueBoolPointer(),
		IsThirdpartyVmSupported: data.IsThirdpartyVmSupported.ValueBoolPointer(),
		MapsApikey:              data.MapsApikey.ValueString(),
		Motd:                    data.Motd.ValueString(),
		PrivilegeFeature:        data.PrivilegeFeature.ValueString(),
		UpgradeAgentVersion:     data.UpgradeAgentVersion.ValueString(),
	}
	return mpsReqPayload
}
func mpsSetAttrFromGet(ctx context.Context, data *mpsModel, getResponseData map[string]interface{}) *mpsModel {
	tflog.Debug(ctx, "In mpsSetAttrFromGet Function")
	if !data.ConfigMotd.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["config_motd"].(string))
		data.ConfigMotd = types.BoolValue(val)
	}
	if !data.DeploymentType.IsNull() {
		data.DeploymentType = types.StringValue(getResponseData["deployment_type"].(string))
	}
	if !data.HistMigInprog.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["hist_mig_inprog"].(string))
		data.HistMigInprog = types.BoolValue(val)
	}
	if !data.IsCloud.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_cloud"].(string))
		data.IsCloud = types.BoolValue(val)
	}
	if !data.IsContainer.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_container"].(string))
		data.IsContainer = types.BoolValue(val)
	}
	if !data.IsMemberOfDefaultGroup.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_member_of_default_group"].(string))
		data.IsMemberOfDefaultGroup = types.BoolValue(val)
	}
	if !data.IsPassive.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_passive"].(string))
		data.IsPassive = types.BoolValue(val)
	}
	if !data.IsThirdpartyVmSupported.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["is_thirdparty_vm_supported"].(string))
		data.IsThirdpartyVmSupported = types.BoolValue(val)
	}
	if !data.MapsApikey.IsNull() {
		data.MapsApikey = types.StringValue(getResponseData["maps_apikey"].(string))
	}
	if !data.Motd.IsNull() {
		data.Motd = types.StringValue(getResponseData["motd"].(string))
	}
	if !data.PrivilegeFeature.IsNull() {
		data.PrivilegeFeature = types.StringValue(getResponseData["privilege_feature"].(string))
	}
	if !data.UpgradeAgentVersion.IsNull() {
		data.UpgradeAgentVersion = types.StringValue(getResponseData["upgrade_agent_version"].(string))
	}
	return data
}

type mpsReq struct {
	ConfigMotd              *bool  `json:"config_motd,omitempty"`
	DeploymentType          string `json:"deployment_type,omitempty"`
	HistMigInprog           *bool  `json:"hist_mig_inprog,omitempty"`
	IsCloud                 *bool  `json:"is_cloud,omitempty"`
	IsContainer             *bool  `json:"is_container,omitempty"`
	IsMemberOfDefaultGroup  *bool  `json:"is_member_of_default_group,omitempty"`
	IsPassive               *bool  `json:"is_passive,omitempty"`
	IsThirdpartyVmSupported *bool  `json:"is_thirdparty_vm_supported,omitempty"`
	MapsApikey              string `json:"maps_apikey,omitempty"`
	Motd                    string `json:"motd,omitempty"`
	PrivilegeFeature        string `json:"privilege_feature,omitempty"`
	UpgradeAgentVersion     string `json:"upgrade_agent_version,omitempty"`
}
