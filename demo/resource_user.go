package demo

import (
	"fmt"
	"context"
	"net/http"
	"bytes"
	"encoding/json"
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

type payload struct {
	User string `json:"user"`
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userName := d.Get("user").(string)

	var diags diag.Diagnostics
	data := payload{
		User: userName, 
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/create", bytes.NewReader(payloadBytes))
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	d.SetId(userName)

	return diags
}


type update struct {
	Old string `json:"old"`
	New string `json:"new"`
}
func resourceUserUpate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("user") {
		o, n := d.GetChange("user")
		old := fmt.Sprintf("%v", o)
		new := fmt.Sprintf("%v", n)
		data := update{
			Old: old, 
			New: new,
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			// handle err
		}
	
		req, err := http.NewRequest("POST", "http://127.0.0.1:5000/update", bytes.NewReader(payloadBytes))
		if err != nil {
			// handle err
		}
		req.Header.Set("Content-Type", "application/json")
	
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
		defer resp.Body.Close()
	
		d.SetId(new)
	}
	return diags
}


func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	userName := d.Get("user").(string)

	var diags diag.Diagnostics
	data := payload{
		User: userName, 
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/delete", bytes.NewReader(payloadBytes))
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	d.SetId("")

	return diags
}