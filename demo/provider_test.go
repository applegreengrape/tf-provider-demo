package demo

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccProviderFunc func() *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"demo": testAccProvider,
	}

	testAccProviderFunc = func() *schema.Provider { return testAccProvider }
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	//
}