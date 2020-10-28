package metadata

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	apiURL = "https://30cqdcgkq7.execute-api.eu-west-1.amazonaws.com"
)

// Config -
type Config struct {
	HostURL string
	APITok  string
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_tok": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENV_API_TOK", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			//"demo_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"metadata_tags": dataSourceTag(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiTok := d.Get("api_tok").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if apiTok != "" {
		c, err := newConfig(host, &apiTok)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create metadata http client",
			})

			return nil, diags
		}

		return c, diags
	}

	log.Fatal("api tok required for this terraform provider")

	return nil, diags
}

func newConfig(host, apiTok *string) (*Config, error) {
	c := Config{
		HostURL: apiURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	c.APITok = *apiTok

	return &c, nil
}
