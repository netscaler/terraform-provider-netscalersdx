package aclrule

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*aclruleDataSource)(nil)

func AclruleDataSource() datasource.DataSource {
	return &aclruleDataSource{}
}

type aclruleDataSource struct {
	client *service.NitroClient
}

func (d *aclruleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aclrule"
}

func (d *aclruleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *aclruleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = aclruleDataSourceSchema()
}

func (d *aclruleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data aclruleModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "aclrule"

	var responseData map[string]interface{}
	var err error

	if !data.Name.IsNull() {
		args := map[string]string{
			"name": data.Name.ValueString(),
		}

		responseData, err = d.client.GetResourceWithArgs(endpoint, args)
		if err != nil {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.Name.ValueString()),
			)
			return
		}
	} else {
		responseData, err = d.client.GetResource(endpoint, data.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.Id.ValueString()),
			)
			return
		}
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("No %s resource found in the remote for the given query.", endpoint),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Id = types.StringValue(getResponseData["id"].(string))

	aclruleSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
