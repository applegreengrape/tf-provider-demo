package demo

import (
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDemoUser_basic(t *testing.T) {
	//resourceName := "demo_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDemoUserConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

const testDemoUserConfig_basic = `
provider "demo" {
	host = "http://127.0.0.1:5000"
}

resource "demo_user" "test" {
  user = "abc123"
}
`
