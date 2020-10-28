package metadata

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)


func TestDataSourceTag(t *testing.T) {
	path := "/v1/tag"
	query_string := "abc"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConfigBasic(path, query_string),
				//Check: resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccCheckDestroy(s *terraform.State) error {
	return nil 
}

func testAccCheckConfigBasic(path, q string) string {
	return fmt.Sprintf(`
data "metadata_tags" "test" {
	path = "%s" 
	query_string = "%s"
}`, path, q)
}

