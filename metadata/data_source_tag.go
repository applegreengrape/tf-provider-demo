package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Description: "api path",
				Required:    true,
			},
			"query_string": {
				Type:        schema.TypeString,
				Description: "api query string",
				Optional:    true,
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	path := d.Get("path").(string)
	q := d.Get("query_string").(string)

	cfg := m.(*Config)
	url := fmt.Sprintf("%s%s?team=%s", cfg.HostURL, path, q)

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.APITok))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	tagMap := make(map[string]string)
	for _, t := range result["tags"].([]interface{}) {
		for k, v := range t.(map[string]interface{}) {
			tagMap[fmt.Sprintf(k)]=fmt.Sprintf("%v", v)
		}
	}

	d.Set("tags", tagMap)

	var diags diag.Diagnostics

	d.SetId(fmt.Sprintf("tag-%s", strconv.FormatInt(time.Now().Unix(), 10)))

	return diags
}
