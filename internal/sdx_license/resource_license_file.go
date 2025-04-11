package sdx_license

import (
	"context"
	"fmt"
	"net/url"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*licenseFileResource)(nil)
var _ resource.ResourceWithConfigure = (*licenseFileResource)(nil)

func LicenseFileResource() resource.Resource {
	return &licenseFileResource{}
}

type licenseFileResource struct {
	client *service.NitroClient
}

func (r *licenseFileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sdx_license"
}

// Configure configures the client resource.
func (r *licenseFileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *licenseFileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = licenseFileResourceSchema()
}

func (r *licenseFileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of license_file Resource")

	var data licenseFileModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create an instance of LicenseFileData
	d := licenseFileData{
		Username: r.client.Username(),
		Password: r.client.Password(),
		Host:     r.client.Host(),
		FileName: data.FileName.ValueString(),
	}

	fmt.Println("hostIP", d.Host)

	sessionId, err := getSessionID(ctx, d)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting session ID",
			fmt.Sprintf("Error getting session ID: %s", err.Error()),
		)
		return
	}
	err = licenseFileMultipart(ctx, d, sessionId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error uploading license file",
			fmt.Sprintf("Error uploading license file: %s", err.Error()),
		)
		return
	}

	act_id, err := sdx_license(ctx, d, sessionId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error applying license",
			fmt.Sprintf("Error applying license: %s", err.Error()),
		)
		return
	}

	err = r.client.WaitForActivityCompletion(act_id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error waiting for activity completion",
			fmt.Sprintf("Error waiting for activity completion: %s", err.Error()),
		)
		return
	}

	data.Id = types.StringValue(d.FileName)
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

func (r *licenseFileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of license_file Resource with Id: %s", resId))

	var data licenseFileModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "license_file"

	responseData, err := r.client.GetResource(endpoint, url.PathEscape(data.Id.ValueString()))
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource license_file: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	licenseFileSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *licenseFileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of license_file Resource")

	var data licenseFileModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "license_file"
	_, err := r.client.DeleteResource(endpoint, url.PathEscape(data.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	d := licenseFileData{
		Username: r.client.Username(),
		Password: r.client.Password(),
		Host:     r.client.Host(),
		FileName: data.FileName.ValueString(),
	}
	sessionId, err := getSessionID(ctx, d)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting session ID",
			fmt.Sprintf("Error getting session ID: %s", err.Error()),
		)
		return
	}
	act_id, err := sdx_license(ctx, d, sessionId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error applying license",
			fmt.Sprintf("Error applying license: %s", err.Error()),
		)
		return
	}

	err = r.client.WaitForActivityCompletion(act_id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error waiting for activity completion",
			fmt.Sprintf("Error waiting for activity completion: %s", err.Error()),
		)
		return
	}

}

func (r *licenseFileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}
