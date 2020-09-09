package demo

import (
	"fmt"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext: resourceUserCreate,
		UpdateContext: resourceUserUpate,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userName := d.Get("user").(string)

	var diags diag.Diagnostics
	err := ClientCreate(userName)
	if err != nil {
		//
	}
	
	d.SetId(userName)

	return diags
}

func resourceUserUpate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("user") {
		o, n := d.GetChange("user")
		old := fmt.Sprintf("%v", o)
		new := fmt.Sprintf("%v", n)

		err:= ClientUpdate(old, new)
		if err != nil {

		}
	
		d.SetId(new)
	}
	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userName := d.Get("user").(string)

	var diags diag.Diagnostics
	
	err := ClientDelete(userName)
	if err != nil {

	}

	d.SetId("")

	return diags
}