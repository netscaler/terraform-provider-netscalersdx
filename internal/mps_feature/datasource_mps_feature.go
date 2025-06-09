package mps_feature

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*mpsFeatureDataSource)(nil)

func MpsFeatureDataSource() datasource.DataSource {
	return &mpsFeatureDataSource{}
}

type mpsFeatureDataSource struct {
	client *service.NitroClient
}

func (d *mpsFeatureDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mps_feature"
}

func (d *mpsFeatureDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *mpsFeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = mpsFeatureDataSourceSchema()
}

func (d *mpsFeatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data mpsFeatureModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "mps_feature"

	args := make(map[string]string)
	if !data.FeatureName.IsNull() {
		args["feature_name"] = data.FeatureName.ValueString()
	} else if !data.Id.IsNull() {
		args["feature_name"] = data.Id.ValueString()
	}

	responseData, err := d.client.GetResourceWithArgs(endpoint, args)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource mps_feature: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"LDAP Server Not Found",
			fmt.Sprintf("No ldap_server resource found in the remote for the given query."),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Id = types.StringValue(getResponseData["feature_name"].(string))

	mpsFeatureSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
