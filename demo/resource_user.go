package demo

import (
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read: resourceUserCreate,
		//Update: resourceRestApiUpdate,
		Delete: resourceUserCreate,

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

type jsonData struct {
	ID string `json:"id"`
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	userName := d.Get("user").(string)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//
	}

	var record jsonData
	if error := json.Unmarshal(body, &record); error != nil {
		//
	}

	d.SetId(record.ID)

	return nil
}