package demo

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
		},
		ResourcesMap: map[string]*schema.Resource{
			"demo_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"demo_user":     dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// HostURL - 
const HostURL string = "http://localhost:5000"

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return c, diags
}