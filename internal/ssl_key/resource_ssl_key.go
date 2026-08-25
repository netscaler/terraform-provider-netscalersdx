package ssl_key

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*sslKeyResource)(nil)
var _ resource.ResourceWithConfigure = (*sslKeyResource)(nil)
var _ resource.ResourceWithImportState = (*sslKeyResource)(nil)

func SslKeyResource() resource.Resource {
	return &sslKeyResource{}
}

type sslKeyResource struct {
	client *service.NitroClient
}

func (r *sslKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *sslKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_key"
}

// Configure configures the client resource.
func (r *sslKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *sslKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = sslKeyResourceSchema(ctx)
}

func (r *sslKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of ssl_key Resource")

	var data sslKeyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := "ssl_key"

	// file_location_path is the local DIRECTORY containing the file; file_name is
	// the file within it and the name the appliance stores it under (= id). The
	// upload is keyed by the multipart filename, so joining the two makes
	// file_name authoritative. The SDX form field is the plural "ssl_keys".
	localPath := filepath.Join(data.FileLocationPath.ValueString(), data.FileName.ValueString())
	_, err := r.client.UploadFile(endpoint, "ssl_keys", localPath)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error uploading file for resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	data.Id = types.StringValue(data.FileName.ValueString())

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

func (r *sslKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of ssl_key Resource with Id: %s", resId))

	var data sslKeyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ssl_key"

	responseData, err := r.client.GetResourceV2(endpoint, url.PathEscape(data.Id.ValueString()))
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ssl_key: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	sslKeySetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. ssl_key has no NITRO upload-update endpoint, so any real
// change to file_name or file_location_path is force-destroyed-and-recreated
// by the RequiresReplace plan modifiers on those attributes. Terraform still
// calls Update for internal state shuffles that the operator hasn't seen
// (e.g. when only Computed attributes change), so we accept the plan, persist
// it to state, and return.
func (r *sslKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of ssl_key Resource (no-op; RequiresReplace handles real changes)")

	var data sslKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *sslKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of ssl_key Resource")

	var data sslKeyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "ssl_key"
	_, err := r.client.DeleteResourceV2(endpoint, url.PathEscape(data.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}

// toInt64 coerces a NITRO scalar (which may be a string, float64, or already
// an int) into an int64. NITRO normally returns scalars as JSON strings, but
// we accept the float64 case too in case the upload endpoint deviates.
func toInt64(v interface{}) int64 {
	switch val := v.(type) {
	case string:
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	case float64:
		return int64(val)
	case int64:
		return val
	case int:
		return int64(val)
	default:
		return 0
	}
}

// toString coerces a NITRO scalar (which may be a string or a JSON number)
// to its string form. Used for file_last_modified_epoch, which the JSON
// metadata declares as Double — modeling it as string preserves any decimal
// precision the server returns without forcing a lossy float conversion.
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int64:
		return strconv.FormatInt(val, 10)
	case int:
		return strconv.Itoa(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
