package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func resourceAssetViewTag() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
		Create: resourceAssetViewTagCreate,
		Read:   resourceAssetViewTagRead,
		Update: resourceAssetViewTagUpdate,
		Delete: resourceAssetViewTagDelete,
		Exists: resourceAssetViewTagExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceAssetViewTagCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	ruleType := d.Get("rule_type").(string)
	ruleText := d.Get("rule_text").(string)
	color := d.Get("color").(string)
	name := d.Get("name").(string)

	tag := client.AssetViewDataTag{}

	if name != "" {
		tag.Name = &name
	}
	if ruleType != "" {
		tag.RuleType = &ruleType
	}
	if ruleText != "" {
		tag.RuleText = &ruleText
	}
	if color != "" {
		tag.Color = &color
	}

	children, err := setChildrenTags(d)
	if err != nil {
		return err
	}
	tag.Children = children

	newTag, err := apiClient.NewAssetViewDataTag(&tag)

	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(*newTag.TagId))
	d.Set("created", newTag.Created)
	d.Set("modified", newTag.Modified)

	return resourceAssetViewTagRead(d, m)
}

func setChildrenTags(d *schema.ResourceData) (*client.AssetViewDataTagsChildren, error) {
	var avChildren *client.AssetViewDataTagsChildren
	if d.Get("children") != nil {
		childTags := d.Get("children").(map[string]interface{})

		var err error
		avChildren, err = readChildrenTags(childTags)
		if err != nil {
			return nil, err
		}
	}
	return avChildren, nil
}

func resourceAssetViewTagRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	var tagIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	tag, err := apiClient.GetAssetViewDataTag(tagIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Tag with ID %s", d.Id())
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

func resourceAssetViewTagUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	// Due to the way the API works, if the name stays the same, we can't pass it
	// As the API says a tag with the same name already exists ;(
	var newName *string
	if d.HasChange("name") {
		name := d.Get("name").(string)
		newName = &name
	}

	var tagIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	ruleType := d.Get("rule_type").(string)
	ruleText := d.Get("rule_text").(string)
	color := d.Get("color").(string)

	tag := client.AssetViewDataTag{
		TagId: &tagIdInt,
		Name:  newName,
	}

	if ruleType != "" {
		tag.RuleType = &ruleType
	}
	if ruleText != "" {
		tag.RuleText = &ruleText
	}
	if color != "" {
		tag.Color = &color
	}

	childrenTags, err := setChildrenTags(d)
	if err != nil {
		return err
	}

	if d.HasChange("children") {
		tag.Children = childrenTags
	}

	tag.Children = childrenTags

	_, err = apiClient.UpdateAssetViewDataTag(&tag)
	if err != nil {
		return err
	}

	// Returning Read as for example when updating tags, the API doesn't return the tag names
	return resourceAssetViewTagRead(d, m)
}

func resourceAssetViewTagDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	var tagIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	err = apiClient.DeleteAssetViewDataTag(tagIdInt)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceAssetViewTagExists(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	var tagIdInt, err = strconv.Atoi(d.Id())

	_, err = apiClient.GetAssetViewDataTag(tagIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
