package ldap_server

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*ldapServerResource)(nil)
var _ resource.ResourceWithConfigure = (*ldapServerResource)(nil)

func LdapServerResource() resource.Resource {
	return &ldapServerResource{}
}

type ldapServerResource struct {
	client *service.NitroClient
}

func (r *ldapServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ldap_server"
}

// Configure configures the client resource.
func (r *ldapServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ldapServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ldapServerResourceSchema(ctx)
}

func (r *ldapServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of ldap_server Resource")

	var data ldapServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ldapServerReq := ldapServerGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "ldap_server"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, ldapServerReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Example data value setting
	data.Id = types.StringValue(resID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

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

func (r *ldapServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of ldap_server Resource with Id: %s", resId))

	var data ldapServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ldap_server"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ldap_server: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.AuthTimeout.IsNull() {
		val, _ := strconv.Atoi(getResponseData["auth_timeout"].(string))
		data.AuthTimeout = types.Int64Value(int64(val))
	}
	if !data.Authentication.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["authentication"].(string))
		data.Authentication = types.BoolValue(val)
	}
	if !data.BaseDn.IsNull() {
		data.BaseDn = types.StringValue(getResponseData["base_dn"].(string))
	}
	if !data.BindDn.IsNull() {
		data.BindDn = types.StringValue(getResponseData["bind_dn"].(string))
	}
	if !data.BindPasswd.IsNull() {
		data.BindPasswd = types.StringValue(getResponseData["bind_passwd"].(string))
	}
	if !data.ChangePassword.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["change_password"].(string))
		data.ChangePassword = types.BoolValue(val)
	}
	if !data.DefaultAuthenticationGroup.IsNull() {
		data.DefaultAuthenticationGroup = types.StringValue(getResponseData["default_authentication_group"].(string))
	}
	if !data.FollowReferrals.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["follow_referrals"].(string))
		data.FollowReferrals = types.BoolValue(val)
	}
	if !data.GroupAttrName.IsNull() {
		data.GroupAttrName = types.StringValue(getResponseData["group_attr_name"].(string))
	}
	if !data.GroupNameIdentifier.IsNull() {
		data.GroupNameIdentifier = types.StringValue(getResponseData["group_name_identifier"].(string))
	}
	if !data.GroupSearchAttribute.IsNull() {
		data.GroupSearchAttribute = types.StringValue(getResponseData["group_search_attribute"].(string))
	}
	if !data.GroupSearchFilter.IsNull() {
		data.GroupSearchFilter = types.StringValue(getResponseData["group_search_filter"].(string))
	}
	if !data.GroupSearchSubattribute.IsNull() {
		data.GroupSearchSubattribute = types.StringValue(getResponseData["group_search_subattribute"].(string))
	}
	if !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !data.LdapHostName.IsNull() {
		data.LdapHostName = types.StringValue(getResponseData["ldap_host_name"].(string))
	}
	if !data.LoginName.IsNull() {
		data.LoginName = types.StringValue(getResponseData["login_name"].(string))
	}
	if !data.MaxLdapReferrals.IsNull() {
		val, _ := strconv.Atoi(getResponseData["max_ldap_referrals"].(string))
		data.MaxLdapReferrals = types.Int64Value(int64(val))
	}
	if !data.MaxNestingLevel.IsNull() {
		val, _ := strconv.Atoi(getResponseData["max_nesting_level"].(string))
		data.MaxNestingLevel = types.Int64Value(int64(val))
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.NestedGroupExtraction.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["nested_group_extraction"].(string))
		data.NestedGroupExtraction = types.BoolValue(val)
	}
	if !data.Port.IsNull() {
		val, _ := strconv.Atoi(getResponseData["port"].(string))
		data.Port = types.Int64Value(int64(val))
	}
	if !data.SearchFilter.IsNull() {
		data.SearchFilter = types.StringValue(getResponseData["search_filter"].(string))
	}
	if !data.SecType.IsNull() {
		data.SecType = types.StringValue(getResponseData["sec_type"].(string))
	}
	if !data.SshPublicKey.IsNull() {
		data.SshPublicKey = types.StringValue(getResponseData["ssh_public_key"].(string))
	}
	if !data.SubattributeName.IsNull() {
		data.SubattributeName = types.StringValue(getResponseData["subattribute_name"].(string))
	}
	if !data.Type.IsNull() {
		data.Type = types.StringValue(getResponseData["type"].(string))
	}
	if !data.ValidateLdapServerCerts.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["validate_ldap_server_certs"].(string))
		data.ValidateLdapServerCerts = types.BoolValue(val)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ldapServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of ldap_server Resource")

	var data ldapServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ldapServerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "ldap_server"
	requestPayload := ldapServerGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, resourceId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Resource",
			fmt.Sprintf("Error updating resource: %s", err.Error()),
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ldapServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of ldap_server Resource")

	var data ldapServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "ldap_server"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
