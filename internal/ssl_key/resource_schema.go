package ssl_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func sslKeyResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description:         "Configuration for SSL key File resource. Uploads an SSL key file from a local path to the SDX appliance via /nitro/v2/upload/ssl_key.",
		MarkdownDescription: "Configuration for SSL key File resource. Uploads an SSL key file from a local path to the SDX appliance via `/nitro/v2/upload/ssl_key`.",
		Attributes: map[string]schema.Attribute{
			"file_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "File Name. The name of the key file inside file_location_path, and the name it is stored under on the SDX appliance (equals the resource id). Minimum length =  1 Maximum length =  256",
				MarkdownDescription: "File Name. The name of the key file inside `file_location_path`, and the name it is stored under on the SDX appliance (equals the resource `id`). Minimum length =  1 Maximum length =  256",
			},
			"file_location_path": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "File Location on Client for upload. This is the local filesystem directory that contains the key file named by file_name, on the machine running terraform apply. The uploaded file is file_location_path/file_name. Minimum length =  1",
				MarkdownDescription: "File Location on Client for upload. This is the local filesystem **directory** that contains the key file named by `file_name`, on the machine running `terraform apply`. The uploaded file is `<file_location_path>/<file_name>`. Minimum length =  1",
			},
			"file_size": schema.Int64Attribute{
				Computed:            true,
				Description:         "File size in bytes as reported by the SDX appliance.",
				MarkdownDescription: "File size in bytes as reported by the SDX appliance.",
			},
			"file_last_modified": schema.StringAttribute{
				Computed:            true,
				Description:         "Last Modified timestamp of the file on the SDX appliance.",
				MarkdownDescription: "Last Modified timestamp of the file on the SDX appliance.",
			},
			"file_last_modified_epoch": schema.StringAttribute{
				Computed:            true,
				Description:         "Last Modified (Epoch) timestamp of the file on the SDX appliance, in epoch seconds.",
				MarkdownDescription: "Last Modified (Epoch) timestamp of the file on the SDX appliance, in epoch seconds.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "The ID of this resource. It is the same as the file_name value.",
				MarkdownDescription: "The ID of this resource. It is the same as the `file_name` value.",
			},
		},
	}
}

type sslKeyModel struct {
	FileName              types.String `tfsdk:"file_name"`
	FileLocationPath      types.String `tfsdk:"file_location_path"`
	FileSize              types.Int64  `tfsdk:"file_size"`
	FileLastModified      types.String `tfsdk:"file_last_modified"`
	FileLastModifiedEpoch types.String `tfsdk:"file_last_modified_epoch"`
	Id                    types.String `tfsdk:"id"`
}

// sslKeySetAttrFromGet maps the NITRO GET response back into the model.
//
// IMPORTANT: file_location_path is intentionally NOT copied from the server
// response. The operator's input is the LOCAL path of the file to upload (e.g.
// "/tmp/my.key"), but the server returns its own server-side storage path
// (e.g. "/var/mps/tenants/root/ns_ssl_keys/"). Overwriting the operator's
// value would produce perpetual drift on every plan. Treat it like a
// write-only field (mirrors snmp_user's handling of passwords).
func sslKeySetAttrFromGet(ctx context.Context, data *sslKeyModel, getResponseData map[string]interface{}) *sslKeyModel {
	tflog.Debug(ctx, "In sslKeySetAttrFromGet Function")

	if v, ok := getResponseData["file_name"]; ok && v != nil {
		data.FileName = types.StringValue(v.(string))
	}
	if v, ok := getResponseData["file_size"]; ok && v != nil {
		data.FileSize = types.Int64Value(toInt64(v))
	}
	if v, ok := getResponseData["file_last_modified"]; ok && v != nil {
		data.FileLastModified = types.StringValue(toString(v))
	}
	if v, ok := getResponseData["file_last_modified_epoch"]; ok && v != nil {
		data.FileLastModifiedEpoch = types.StringValue(toString(v))
	}

	return data
}
