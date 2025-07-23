---
subcategory: "CDN"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cdn_show_tags"
description: |-
  Provides a resource to query tags of a CDN domain in HuaweiCloud.
---

# huaweicloud_cdn_show_tags

Provides a resource to query tags of a CDN domain in HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_cdn_show_tags" "example" {
  resource_id = "your-cdn-domain-id"
}

output "cdn_tags" {
  value = huaweicloud_cdn_show_tags.example.tags
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required, String) Specifies the ID of the CDN domain to query tags for.

## Attributes Reference

The following attributes are exported:

* `tags` - A list of tags associated with the resource. Each tag contains the following fields:

  * `key` - The tag key.
  * `value` - The tag value.

## Notes
- This resource is read-only and will not modify any CDN resource.
- The `resource_id` must be a valid CDN domain ID. 