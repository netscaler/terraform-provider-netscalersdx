package snmp_mib

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*snmpMibDataSource)(nil)

func SnmpMibDataSource() datasource.DataSource {
	return &snmpMibDataSource{}
}

type snmpMibDataSource struct {
	client *service.NitroClient
}

func (d *snmpMibDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_mib"
}

func (d *snmpMibDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *snmpMibDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = snmpMibDataSourceSchema()
}

func (d *snmpMibDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data snmpMibModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_mib"

	var resid string
	if !data.Name.IsNull() {
		resid = data.Name.ValueString()
	} else if !data.Id.IsNull() {
		resid = data.Id.ValueString()
	}

	if resid == "" {
		resp.Diagnostics.AddError(
			"Missing snmp mib ID",
			"Either 'name' or 'id' must be provided to identify the snmp mib.",
		)
		return
	}

	responseData, err := d.client.GetAllResource(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s is not present in the remote", endpoint),
		)
		return
	}

	data.Id = types.StringValue(resid)

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	snmpMibSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
