package snmp_manager

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*snmpManagerDataSource)(nil)

func SnmpManagerDataSource() datasource.DataSource {
	return &snmpManagerDataSource{}
}

type snmpManagerDataSource struct {
	client *service.NitroClient
}

func (d *snmpManagerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_manager"
}

func (d *snmpManagerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *snmpManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = snmpManagerDataSourceSchema()
}

func (d *snmpManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data snmpManagerModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_manager"

	var resid string
	if !data.IpAddress.IsNull() {
		resid = data.IpAddress.ValueString()
	} else if !data.Id.IsNull() {
		resid = data.Id.ValueString()
	}

	if resid == "" {
		resp.Diagnostics.AddError(
			"Missing SNMP manager ID",
			"Either 'ip_address' or 'id' must be provided to identify the SNMP manager.",
		)
		return
	}

	responseData, err := d.client.GetResource(endpoint, resid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.IpAddress.ValueString()),
		)
		return
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("No %s resource found in the remote for the given query.", endpoint),
		)
		return
	}

	data.Id = types.StringValue(resid)

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	snmpManagerSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
