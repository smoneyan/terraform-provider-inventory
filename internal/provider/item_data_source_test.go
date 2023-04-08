package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccItemDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				/*
				   The Inventory service always gives the first `item` created
				   the `id` number 1000, so that is why we default to testing
				   for that specific `id`.
				*/
				Config: providerConfig + `
data "inventory_item" "test" {
 id = 1000
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("data.inventory_item.test", "id"),
				),
			},
		},
	})
}
