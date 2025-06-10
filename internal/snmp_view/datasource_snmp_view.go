package snmp_view

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*snmpViewDataSource)(nil)

func SnmpViewDataSource() datasource.DataSource {
	return &snmpViewDataSource{}
}

type snmpViewDataSource struct {
	client *service.NitroClient
}

func (d *snmpViewDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_view"
}

func (d *snmpViewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *snmpViewDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = snmpViewDataSourceSchema()
}

func (d *snmpViewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data snmpViewModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_view"

	var resid string
	if !data.Name.IsNull() {
		resid = data.Name.ValueString()
	} else if !data.Id.IsNull() {
		resid = data.Id.ValueString()
	}

	if resid == "" {
		resp.Diagnostics.AddError(
			"Missing snmp view ID",
			"Either 'name' or 'id' must be provided to identify the snmp view.",
		)
		return
	}

	responseData, err := d.client.GetResource(endpoint, resid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s: %v is not present in the remote", endpoint, resid),
		)
		return
	}

	data.Id = types.StringValue(resid)

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	snmpViewSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
