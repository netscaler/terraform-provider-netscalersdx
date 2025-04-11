package sdx_license

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func licenseFileResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Upload and apply license file to the Citrix SDX appliance.",
		Attributes: map[string]schema.Attribute{
			"file_name": schema.StringAttribute{
				Required: true,
				// We have below code insted of ForceNew
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "License file name to be uploaded and applied. Note: License file should be present in the current directory where the Terraform configuration file is located.",
				MarkdownDescription: "License file name to be uploaded and applied. Note: License file should be present in the current directory where the Terraform configuration file is located",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type licenseFileModel struct {
	FileName types.String `tfsdk:"file_name"`
	Id       types.String `tfsdk:"id"`
}

func licenseFileSetAttrFromGet(ctx context.Context, data *licenseFileModel, getResponseData map[string]interface{}) *licenseFileModel {
	tflog.Debug(ctx, "In licenseFileSetAttrFromGet Function")

	data.FileName = types.StringValue(getResponseData["file_name"].(string))

	return data
}

type licenseFileData struct {
	Username string
	Password string
	Host     string
	FileName string
}
