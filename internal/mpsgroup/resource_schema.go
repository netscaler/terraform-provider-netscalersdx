package mpsgroup

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-netscalersdx/internal/utils"
)

func mpsgroupResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for System Groups resource.",
		Attributes: map[string]schema.Attribute{
			"allow_application_only": schema.BoolAttribute{
				Optional:            true,
				Description:         "Checks if only application centic page is needed.",
				MarkdownDescription: "Checks if only application centic page is needed.",
			},
			"application_names": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "All Application names that are part of this group.This includes selected appnames as well as applications which are result of defined regex.",
				MarkdownDescription: "All Application names that are part of this group.This includes selected appnames as well as applications which are result of defined regex.",
			},
			"apply_all_bound_entities": schema.BoolAttribute{
				Optional:            true,
				Description:         "Apply for all bound entities (TRUE|FALSE).",
				MarkdownDescription: "Apply for all bound entities (TRUE|FALSE).",
			},
			"assign_all_apps": schema.BoolAttribute{
				Optional:            true,
				Description:         "Assign All Applications (YES|NO).",
				MarkdownDescription: "Assign All Applications (YES|NO).",
			},
			"assign_all_autoscale_groups": schema.BoolAttribute{
				Optional:            true,
				Description:         "Assign All Autoscale groups (YES|NO).",
				MarkdownDescription: "Assign All Autoscale groups (YES|NO).",
			},
			"assign_all_devices": schema.BoolAttribute{
				Optional:            true,
				Description:         "Assign All Instances (YES|NO).",
				MarkdownDescription: "Assign All Instances (YES|NO).",
			},
			"assign_all_selected_device_apps": schema.BoolAttribute{
				Optional:            true,
				Description:         "Assign All Application from selected instances (YES|NO).",
				MarkdownDescription: "Assign All Application from selected instances (YES|NO).",
			},
			"autoscale_groups_id": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Autoscale groups belong to this groupp.",
				MarkdownDescription: "Autoscale groups belong to this groupp.",
			},
			"bound_entity_selected": schema.Int64Attribute{
				Optional:            true,
				Description:         "Which bound entiy is selected VSERVER(0),SERVICE(1),SERVICEGROUP(2),SERVER(3).",
				MarkdownDescription: "Which bound entiy is selected VSERVER(0),SERVICE(1),SERVICEGROUP(2),SERVER(3).",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Description:         "Description of Group. Minimum length =  1 Maximum length =  1024",
				MarkdownDescription: "Description of Group. Minimum length =  1 Maximum length =  1024",
			},
			"enable_session_timeout": schema.BoolAttribute{
				Optional:            true,
				Description:         "Enables session timeout for group.",
				MarkdownDescription: "Enables session timeout for group.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Group Name. Minimum length =  1 Maximum length =  64",
				MarkdownDescription: "Group Name. Minimum length =  1 Maximum length =  64",
			},
			"permission": schema.StringAttribute{
				Required:            true,
				Description:         "Permission for the group (admin/read-only). Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Permission for the group (admin/read-only). Minimum length =  1 Maximum length =  128",
			},
			"role": schema.StringAttribute{
				Optional:            true,
				Description:         "Role (admin|nonadmin).",
				MarkdownDescription: "Role (admin|nonadmin).",
			},
			"roles": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Roles assigned to the group.",
				MarkdownDescription: "Roles assigned to the group.",
			},
			"select_individual_entity": schema.BoolAttribute{
				Optional:            true,
				Description:         "Select Individual Entity Type.",
				MarkdownDescription: "Select Individual Entity Type.",
			},
			"session_timeout": schema.Int64Attribute{
				Optional:            true,
				Description:         "Session timeout for the Group.",
				MarkdownDescription: "Session timeout for the Group.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Optional:            true,
				Description:         "Session timeout unit for the Group.",
				MarkdownDescription: "Session timeout unit for the Group.",
			},
			"standalone_instances_id": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Stand alone instances belong to this groupp.",
				MarkdownDescription: "Stand alone instances belong to this groupp.",
			},
			"tenant_id": schema.StringAttribute{
				Optional:            true,
				Description:         "Id of the tenant. Minimum length =  1 Maximum length =  128",
				MarkdownDescription: "Id of the tenant. Minimum length =  1 Maximum length =  128",
			},
			"users": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Users belong to the group.",
				MarkdownDescription: "Users belong to the group.",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type mpsgroupModel struct {
	AllowApplicationOnly        types.Bool   `tfsdk:"allow_application_only"`
	ApplicationNames            types.List   `tfsdk:"application_names"`
	ApplyAllBoundEntities       types.Bool   `tfsdk:"apply_all_bound_entities"`
	AssignAllApps               types.Bool   `tfsdk:"assign_all_apps"`
	AssignAllAutoscaleGroups    types.Bool   `tfsdk:"assign_all_autoscale_groups"`
	AssignAllDevices            types.Bool   `tfsdk:"assign_all_devices"`
	AssignAllSelectedDeviceApps types.Bool   `tfsdk:"assign_all_selected_device_apps"`
	AutoscaleGroupsId           types.List   `tfsdk:"autoscale_groups_id"`
	BoundEntitySelected         types.Int64  `tfsdk:"bound_entity_selected"`
	Description                 types.String `tfsdk:"description"`
	EnableSessionTimeout        types.Bool   `tfsdk:"enable_session_timeout"`
	Name                        types.String `tfsdk:"name"`
	Permission                  types.String `tfsdk:"permission"`
	Role                        types.String `tfsdk:"role"`
	Roles                       types.List   `tfsdk:"roles"`
	SelectIndividualEntity      types.Bool   `tfsdk:"select_individual_entity"`
	SessionTimeout              types.Int64  `tfsdk:"session_timeout"`
	SessionTimeoutUnit          types.String `tfsdk:"session_timeout_unit"`
	StandaloneInstancesId       types.List   `tfsdk:"standalone_instances_id"`
	TenantId                    types.String `tfsdk:"tenant_id"`
	Users                       types.List   `tfsdk:"users"`
	Id                          types.String `tfsdk:"id"`
}

func mpsgroupGetThePayloadFromtheConfig(ctx context.Context, data *mpsgroupModel) mpsgroupReq {
	tflog.Debug(ctx, "In mpsgroupGetThePayloadFromtheConfig Function")
	// tflog.Debug(ctx, fmt.Sprintf("Data: %v", *data.ApplyAllBoundEntities.ValueBoolPointer()))
	tflog.Debug(ctx, fmt.Sprintf("Data: %v", data.ApplyAllBoundEntities.ValueBoolPointer()))
	mpsgroupReqPayload := mpsgroupReq{
		AllowApplicationOnly:        data.AllowApplicationOnly.ValueBoolPointer(),
		ApplicationNames:            utils.TypeListToStringList(data.ApplicationNames),
		ApplyAllBoundEntities:       data.ApplyAllBoundEntities.ValueBoolPointer(),
		AssignAllApps:               data.AssignAllApps.ValueBoolPointer(),
		AssignAllAutoscaleGroups:    data.AssignAllAutoscaleGroups.ValueBoolPointer(),
		AssignAllDevices:            data.AssignAllDevices.ValueBoolPointer(),
		AssignAllSelectedDeviceApps: data.AssignAllSelectedDeviceApps.ValueBoolPointer(),
		AutoscaleGroupsId:           utils.TypeListToStringList(data.AutoscaleGroupsId),
		BoundEntitySelected:         data.BoundEntitySelected.ValueInt64Pointer(),
		Description:                 data.Description.ValueStringPointer(),
		EnableSessionTimeout:        data.EnableSessionTimeout.ValueBoolPointer(),
		Name:                        data.Name.ValueString(),
		Permission:                  data.Permission.ValueString(),
		Role:                        data.Role.ValueString(),
		Roles:                       utils.TypeListToStringList(data.Roles),
		SelectIndividualEntity:      data.SelectIndividualEntity.ValueBoolPointer(),
		SessionTimeout:              data.SessionTimeout.ValueInt64Pointer(),
		SessionTimeoutUnit:          data.SessionTimeoutUnit.ValueString(),
		StandaloneInstancesId:       utils.TypeListToStringList(data.StandaloneInstancesId),
		TenantId:                    data.TenantId.ValueString(),
		Users:                       utils.TypeListToStringList(data.Users),
	}

	return mpsgroupReqPayload
}
func mpsgroupSetAttrFromGet(ctx context.Context, data *mpsgroupModel, getResponseData map[string]interface{}) *mpsgroupModel {
	tflog.Debug(ctx, fmt.Sprintf("In mpsgroupSetAttrFromGet Function %v", data))

	if !data.AllowApplicationOnly.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["allow_application_only"].(string))
		data.AllowApplicationOnly = types.BoolValue(val)
	}
	if !data.ApplicationNames.IsNull() {
		data.ApplicationNames = utils.StringListToTypeList(utils.ToStringList(getResponseData["application_names"].([]interface{})))
	}
	if !data.ApplyAllBoundEntities.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["apply_all_bound_entities"].(string))
		data.ApplyAllBoundEntities = types.BoolValue(val)
	}
	if !data.AssignAllApps.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["assign_all_apps"].(string))
		data.AssignAllApps = types.BoolValue(val)
	}
	if !data.AssignAllAutoscaleGroups.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["assign_all_autoscale_groups"].(string))
		data.AssignAllAutoscaleGroups = types.BoolValue(val)
	}
	if !data.AssignAllDevices.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["assign_all_devices"].(string))
		data.AssignAllDevices = types.BoolValue(val)
	}
	if !data.AssignAllSelectedDeviceApps.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["assign_all_selected_device_apps"].(string))
		data.AssignAllSelectedDeviceApps = types.BoolValue(val)
	}
	if !data.AutoscaleGroupsId.IsNull() {
		data.AutoscaleGroupsId = utils.StringListToTypeList(utils.ToStringList(getResponseData["autoscale_groups_id"].([]interface{})))
	}
	if !data.BoundEntitySelected.IsNull() {
		val, _ := strconv.Atoi(getResponseData["bound_entity_selected"].(string))
		data.BoundEntitySelected = types.Int64Value(int64(val))
	}
	if !data.Description.IsNull() {
		data.Description = types.StringValue(getResponseData["description"].(string))
	}
	if !data.EnableSessionTimeout.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_session_timeout"].(string))
		data.EnableSessionTimeout = types.BoolValue(val)
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.Permission.IsNull() {
		data.Permission = types.StringValue(getResponseData["permission"].(string))
	}
	if !data.Role.IsNull() {
		data.Role = types.StringValue(getResponseData["role"].(string))
	}
	if !data.Roles.IsNull() {
		data.Roles = utils.StringListToTypeList(utils.ToStringList(getResponseData["roles"].([]interface{})))
	}
	if !data.SelectIndividualEntity.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["select_individual_entity"].(string))
		data.SelectIndividualEntity = types.BoolValue(val)
	}
	if !data.SessionTimeout.IsNull() {
		val, _ := strconv.Atoi(getResponseData["session_timeout"].(string))
		data.SessionTimeout = types.Int64Value(int64(val))
	}
	if !data.SessionTimeoutUnit.IsNull() {
		data.SessionTimeoutUnit = types.StringValue(getResponseData["session_timeout_unit"].(string))
	}
	if !data.StandaloneInstancesId.IsNull() {
		data.StandaloneInstancesId = utils.StringListToTypeList(utils.ToStringList(getResponseData["standalone_instances_id"].([]interface{})))
	}
	if !data.TenantId.IsNull() {
		data.TenantId = types.StringValue(getResponseData["tenant_id"].(string))
	}
	if !data.Users.IsNull() {
		data.Users = utils.StringListToTypeList(utils.ToStringList(getResponseData["users"].([]interface{})))
	}
	return data
}

type mpsgroupReq struct {
	AllowApplicationOnly        *bool    `json:"allow_application_only,omitempty"`
	ApplicationNames            []string `json:"application_names,omitempty"`
	ApplyAllBoundEntities       *bool    `json:"apply_all_bound_entities,omitempty"`
	AssignAllApps               *bool    `json:"assign_all_apps,omitempty"`
	AssignAllAutoscaleGroups    *bool    `json:"assign_all_autoscale_groups,omitempty"`
	AssignAllDevices            *bool    `json:"assign_all_devices,omitempty"`
	AssignAllSelectedDeviceApps *bool    `json:"assign_all_selected_device_apps,omitempty"`
	AutoscaleGroupsId           []string `json:"autoscale_groups_id,omitempty"`
	BoundEntitySelected         *int64   `json:"bound_entity_selected,omitempty"`
	Description                 *string  `json:"description,omitempty"`
	EnableSessionTimeout        *bool    `json:"enable_session_timeout,omitempty"`
	Name                        string   `json:"name,omitempty"`
	Permission                  string   `json:"permission,omitempty"`
	Role                        string   `json:"role,omitempty"`
	Roles                       []string `json:"roles,omitempty"`
	SelectIndividualEntity      *bool    `json:"select_individual_entity,omitempty"`
	SessionTimeout              *int64   `json:"session_timeout,omitempty"`
	SessionTimeoutUnit          string   `json:"session_timeout_unit,omitempty"`
	StandaloneInstancesId       []string `json:"standalone_instances_id,omitempty"`
	TenantId                    string   `json:"tenant_id,omitempty"`
	Users                       []string `json:"users,omitempty"`
}
