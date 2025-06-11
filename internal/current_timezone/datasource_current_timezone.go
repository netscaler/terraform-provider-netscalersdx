package current_timezone

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*currentTimezoneDataSource)(nil)

func CurrentTimezoneDataSource() datasource.DataSource {
	return &currentTimezoneDataSource{}
}

type currentTimezoneDataSource struct {
	client *service.NitroClient
}

func (d *currentTimezoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_current_timezone"
}

func (d *currentTimezoneDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *currentTimezoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = currentTimezoneDataSourceSchema()
}

func (d *currentTimezoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data currentTimezoneModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "current_timezone"

	responseData, err := d.client.GetAllResource(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s is not present in the remote", endpoint),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	currentTimezoneSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
