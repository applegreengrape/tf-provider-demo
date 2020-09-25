package metadata

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("env_host", nil),
			},
			"auth_tok": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("auth_tok", nil),
		},
		ResourcesMap: map[string]*schema.Resource{
			//"demo_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			//"demo_user":     dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}


//https://github.com/hashicorp/terraform-provider-hashicups/blob/167ac4714f9275713d60f29f07a3db8fb8b08eaa/vendor/github.com/hashicorp-demoapp/hashicups-client-go/client.go

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	auth_tok := d.Get("auth_tok").(string)
	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if auth_tok != "" {
		c, err := NewClient(host, &tok)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to authenticate user for authenticated HashiCups client",
			})

			return nil, diags
		}

		return c, diags
	}


	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return c, diags
}