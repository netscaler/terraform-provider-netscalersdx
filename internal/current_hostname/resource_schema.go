package current_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func currentHostnameResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for System Hostname resource.",
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Hostname. Minimum length =  1 Maximum length =  63",
				MarkdownDescription: "Hostname. Minimum length =  1 Maximum length =  63",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description:         "Hypervisor Hostname. Minimum length =  1 Maximum length =  63",
				MarkdownDescription: "Hypervisor Hostname. Minimum length =  1 Maximum length =  63",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type currentHostnameModel struct {
	Hostname           types.String `tfsdk:"hostname"`
	HypervisorHostname types.String `tfsdk:"hypervisor_hostname"`
	Id                 types.String `tfsdk:"id"`
}

func currentHostnameGetThePayloadFromtheConfig(ctx context.Context, data *currentHostnameModel) currentHostnameReq {
	tflog.Debug(ctx, "In currentHostnameGetThePayloadFromtheConfig Function")
	currentHostnameReqPayload := currentHostnameReq{
		Hostname:           data.Hostname.ValueString(),
		HypervisorHostname: data.HypervisorHostname.ValueString(),
	}
	return currentHostnameReqPayload
}
func currentHostnameSetAttrFromGet(ctx context.Context, data *currentHostnameModel, getResponseData map[string]interface{}) *currentHostnameModel {
	tflog.Debug(ctx, "In currentHostnameSetAttrFromGet Function")

	data.Hostname = types.StringValue(getResponseData["hostname"].(string))
	data.HypervisorHostname = types.StringValue(getResponseData["hypervisor_hostname"].(string))

	return data
}

type currentHostnameReq struct {
	Hostname           string `json:"hostname,omitempty"`
	HypervisorHostname string `json:"hypervisor_hostname,omitempty"`
}
