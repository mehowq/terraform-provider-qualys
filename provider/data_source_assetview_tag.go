package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func dataSourceAssetViewTag() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modified": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"children": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
		Read: dataSourceAssetViewTagRead,
	}
}

func dataSourceAssetViewTagRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	name := d.Get("name").(string)
	tagId := d.Get("tag_id").(int)

	var tag *client.AssetViewDataTag
	if name != "" {
		crTagId := client.AssetViewFiltersCriteria{
			Field:    "name",
			Operator: "EQUALS",
			Value:    name,
		}
		criteria := []client.AssetViewFiltersCriteria{crTagId}
		tags, err := apiClient.SearchAssetViewTag(criteria)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				d.SetId("")
			} else {
				return fmt.Errorf("error finding AssetView Tag with name %s", name)
			}
		} else {
			if (len(*tags)) == 1 {
				tag = &(*tags)[0]
			} else {
				return fmt.Errorf("AssetView Tag with name %s doesn't exist", name)
			}
		}
	} else {
		var err error
		tag, err = apiClient.GetAssetViewDataTag(tagId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				d.SetId("")
			} else {
				return fmt.Errorf("error finding AssetView Tag with ID %d", tagId)
			}
		}
	}

	d.SetId(strconv.Itoa(*tag.TagId))
	d.Set("name", tag.Name)
	d.Set("created", *tag.Created)
	d.Set("modified", *tag.Modified)

	if tag.RuleType != nil {
		d.Set("rule_type", *tag.RuleType)
	}
	if tag.RuleText != nil {
		d.Set("rule_text", *tag.RuleText)
	}
	if tag.Color != nil {
		d.Set("color", *tag.Color)
	}

	childrenMap := make(map[string]string)
	if (tag.Children != nil) && (tag.Children.TagsList != nil) {
		childrenMap = convertToTagsMap(tag.Children.TagsList.TagSimple)
	}
	d.Set("children", childrenMap)

	return nil
}
