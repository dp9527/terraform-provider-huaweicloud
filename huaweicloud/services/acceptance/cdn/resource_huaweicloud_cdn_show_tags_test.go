package cdn

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cdn"
)

func getCdnShowTagsResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	var (
		region     = acceptance.HW_REGION_NAME
		product    = "cdn"
		resourceID = state.Primary.Attributes["resource_id"]
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return nil, fmt.Errorf("error creating CDN client: %s", err)
	}

	url := fmt.Sprintf("/v1.0/cdn/configuration/tags?resource_id=%s", resourceID)
	resp, err := client.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TestAccCdnShowTags_basic(t *testing.T) {
	var (
		obj   interface{}
		rName = "huaweicloud_cdn_show_tags.test"
	)

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getCdnShowTagsResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCdnShowTags_basic,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(rName, "resource_id"),
					resource.TestCheckResourceAttrSet(rName, "tags"),
				),
			},
		},
	})
}

const testCdnShowTags_basic = `
resource "huaweicloud_cdn_show_tags" "test" {
  resource_id = "your-cdn-domain-id"
}
` 