// ---------------------------------------------------------------
// @Product CDN
// ---------------------------------------------------------------

package cdn

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CDN GET /v1.0/cdn/configuration/tags
func ResourceCdnShowTags() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCdnShowTagsCreate,
		ReadContext:   resourceCdnShowTagsRead,
		DeleteContext: resourceCdnShowTagsDelete,

		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ID of the CDN domain to query tags for.`,
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of tags associated with the resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The tag key.`,
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The tag value.`,
						},
					},
				},
			},
		},
	}
}

func resourceCdnShowTagsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// 只做Read操作，Create时直接调用Read
	return resourceCdnShowTagsRead(ctx, d, meta)
}

func resourceCdnShowTagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	client, err := cfg.NewServiceClient("cdn", region)
	if err != nil {
		return diag.Errorf("error creating CDN client: %s", err)
	}

	resourceID := d.Get("resource_id").(string)
	url := fmt.Sprintf("/v1.0/cdn/configuration/tags?resource_id=%s", resourceID)
	opt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
	}
	resp, err := client.Request("GET", url, &opt)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error querying CDN resource tags")
	}
	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.Errorf("error parsing CDN resource tags response: %s", err)
	}

	randomUUID, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}
	d.SetId(randomUUID)

	mErr := multierror.Append(nil,
		d.Set("tags", flattenCdnShowTagsTags(utils.PathSearch("tags", respBody, make([]interface{}, 0)).([]interface{}))),
	)
	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceCdnShowTagsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// 查询型资源，删除时直接移除状态
	d.SetId("")
	return nil
}

// flattenCdnShowTagsTags 解析tags列表
func flattenCdnShowTagsTags(tags []interface{}) []map[string]interface{} {
	if len(tags) < 1 {
		return nil
	}
	result := make([]map[string]interface{}, 0, len(tags))
	for _, t := range tags {
		tag := t.(map[string]interface{})
		result = append(result, map[string]interface{}{
			"key":   utils.PathSearch("key", tag, ""),
			"value": utils.PathSearch("value", tag, ""),
		})
	}
	return result
} 