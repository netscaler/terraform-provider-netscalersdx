package mpsgroup

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*mpsgroupDataSource)(nil)

func MpsgroupDataSource() datasource.DataSource {
	return &mpsgroupDataSource{}
}

type mpsgroupDataSource struct {
	client *service.NitroClient
}

func (d *mpsgroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpsgroup"
}

func (d *mpsgroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *mpsgroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = mpsgroupDataSourceSchema()
}

func (d *mpsgroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data mpsgroupModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "mpsgroup"

	var responseData map[string]interface{}
	var err error

	if !data.Name.IsNull() {
		args := map[string]string{
			"name": data.Name.ValueString(),
		}

		responseData, err = d.client.GetResourceWithArgs(endpoint, args)
		if err != nil {
			resp.State.RemoveResource(ctx)
			tflog.Warn(ctx, fmt.Sprintf("removing resource ldap_server: %v from state because it is not present in the remote", data.Id.ValueString()))
			return
		}
	} else {
		responseData, err = d.client.GetResource(endpoint, data.Id.ValueString())
		if err != nil {
			resp.State.RemoveResource(ctx)
			tflog.Warn(ctx, fmt.Sprintf("removing resource ldap_server: %v from state because it is not present in the remote", data.Id.ValueString()))
			return
		}
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"LDAP Server Not Found",
			fmt.Sprintf("No ldap_server resource found in the remote for the given query."),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Id = types.StringValue(getResponseData["id"].(string))

	mpsgroupSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
