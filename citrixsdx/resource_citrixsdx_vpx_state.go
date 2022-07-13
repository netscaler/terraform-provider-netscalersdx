package citrixsdx

import (
	"context"
	"fmt"
	"log"
	"time"

	"terraform-provider-citrixsdx/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpxState() *schema.Resource {
	return &schema.Resource{
		Description:   "Change the state of a VPX",
		CreateContext: resourceVpxStateCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema: map[string]*schema.Schema{
			"vpx_id": {
				Description: "VPX ID",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"state": {
				Description: "Desired state of the VPX",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					if !service.Contains([]string{"start", "stop", "force_stop", "reboot", "force_reboot"}, v.(string)) {
						errors = append(errors, fmt.Errorf("state must be one of: 'start' or 'stop' or 'force_stop' or 'reboot' or 'force_reboot'"))
					}
					return
				},
			},
		},
	}
}

func resourceVpxStateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In resourceVpxStateCreate")

	c := m.(*service.NitroClient)

	endpoint := "ns"
	data := make(map[string]interface{})

	vpxState := d.Get("state").(string)
	vpxID := d.Get("vpx_id").(string)

	_, err := c.AddResourceWithActionParams(endpoint, data, vpxState, vpxID)

	if err != nil {
		return diag.Errorf("unable to set VPX state to %s: %s", vpxState, err.Error())
	}

	if vpxState == "start" || vpxState == "reboot" || vpxState == "force_reboot" {
		// wait for VPX instance_state to be Up
		log.Printf("Wait for VPX instance_state to be Up")

		for {
			time.Sleep(5 * time.Second)

			returnData, err := c.GetResource(endpoint, vpxID)
			if err != nil {
				return diag.Errorf("unable to get VPX: %s", err.Error())
			}
			instanceState := returnData[endpoint].([]interface{})[0].(map[string]interface{})["instance_state"].(string)
			if instanceState == "Up" {
				break
			}
			log.Printf("VPX instance_state is %s", instanceState)
		}
	}

	d.SetId(vpxID)
	return diag.Diagnostics{}
}
