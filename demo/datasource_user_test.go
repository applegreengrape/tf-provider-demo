package demo

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceDemoUser_basic(t *testing.T) {
	datasourceName := "data.demo_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDemoUserConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "user", "abc06"),
				),
			},
		},
	})
}

const testDataSourceDemoUserConfig_basic = `
data "demo_user" "test" {
    user = "abc06"
}
`