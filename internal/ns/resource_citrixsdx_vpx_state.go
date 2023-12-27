package ns

import (
	"context"
	"fmt"
	"time"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &vpxStateResource{}
	_ resource.ResourceWithConfigure = &vpxStateResource{}
)

// vpxStateResource defines the resource implementation
type vpxStateResource struct {
	client *service.NitroClient
}

func VpxStateResource() resource.Resource {
	return &vpxStateResource{}
}

type vpxStateResourceModel struct {
	VpxID types.String `tfsdk:"vpx_id"`
	State types.String `tfsdk:"state"`
	Id    types.String `tfsdk:"id"`
}

type vpxStateResourceReq struct {
	VpxId string `json:"vpx_id"`
	State string `json:"state"`
	Id    string `json:"id"`
}

// Metadata returns the resourcetype
func (r *vpxStateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpx_state"
}

// Configure configures the client resource
func (r *vpxStateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *vpxStateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Change the state of a VPX",
		Attributes: map[string]schema.Attribute{
			"vpx_id": schema.StringAttribute{
				Required:    true,
				Description: "VPX ID",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"state": schema.StringAttribute{
				Required:    true,
				Description: "Desired state of the VPX",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the resource",
			},
		},
	}
}

func (r *vpxStateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "[DEBUG] In Create Method of vpxStateResource")

	var data vpxStateResourceModel

	// Read Terraform plan data into data model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create request data
	vpxStateReq := vpxStateResourceReq{
		VpxId: data.VpxID.ValueString(),
		State: data.State.ValueString(),
	}

	vpxState := vpxStateReq.State
	vpxID := vpxStateReq.VpxId

	endpoint := "ns"
	payload := make(map[string]interface{})

	_, err := r.client.AddResourceWithActionParams(endpoint, payload, vpxState, vpxID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error in Create Method of VPX state Resource",
			fmt.Sprintf("unable to set VPX state to %s: %s", vpxState, err.Error()))
		return
	}

	if vpxState == "start" || vpxState == "reboot" || vpxState == "force_reboot" {

		// wait for VPX instance_state to be Up
		tflog.Debug(ctx, "Wait for VPX instance_state to be Up")

		for {
			time.Sleep(5 * time.Second)

			returnData, err := r.client.GetResource(endpoint, vpxID)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Getting Resource in Create Method of VPX state Resource",
					fmt.Sprintf("unable to get VPX: %s", err.Error()),
				)
				return
			}
			instanceState := returnData[endpoint].([]interface{})[0].(map[string]interface{})["instance_state"].(string)
			if instanceState == "Up" {
				break
			}
			tflog.Debug(ctx, fmt.Sprintf("VPX instance_state is %s", instanceState))
		}
	}

	data.Id = types.StringValue(vpxID)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (*vpxStateResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
}

func (*vpxStateResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
}

func (*vpxStateResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
}
