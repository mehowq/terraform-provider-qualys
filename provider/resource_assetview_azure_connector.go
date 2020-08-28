package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func resourceAssetViewAzureConnector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"directory_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authentication_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"is_gov_cloud": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"last_synced_on": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_assets": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceAssetViewAzureConnectorCreate,
		Read:   resourceAssetViewAzureConnectorRead,
		Update: resourceAssetViewAzureConnectorUpdate,
		Delete: resourceAssetViewAzureConnectorDelete,
		Exists: resourceAssetViewAzureConnectorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceAssetViewAzureConnectorCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	dirId := d.Get("directory_id").(string)
	subId := d.Get("subscription_id").(string)
	appId := d.Get("application_id").(string)
	authKey := d.Get("authentication_key").(string)
	authRec := client.AssetViewDataAuthRecord{
		DirectoryId:       &dirId,
		SubscriptionId:    &subId,
		ApplicationId:     &appId,
		AuthenticationKey: &authKey,
	}

	defTags, err := setDefaultTags(d)
	if err != nil {
		return err
	}

	var newName *string
	if d.HasChange("name") {
		name := d.Get("name").(string)
		newName = &name
	}

	connector := client.AssetViewAzureConnector{
		Name:        newName,
		Description: d.Get("description").(string),
		AuthRecord:  authRec,
		IsGovCloud:  d.Get("is_gov_cloud").(bool),
		DefaultTags: defTags,
	}

	newConnector, err := apiClient.NewAssetViewAzureConnector(&connector)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(*newConnector.ConnectorId))
	d.Set("last_synced_on", newConnector.LastSyncedOn)
	d.Set("total_assets", newConnector.TotalAssets)
	d.Set("state", newConnector.State)

	return resourceAssetViewAzureConnectorRead(d, m)
}

func resourceAssetViewAzureConnectorRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	var connIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	connector, err := apiClient.GetAssetViewAzureConnector(connIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Connector with ID %s", d.Id())
		}
	}

	d.SetId(strconv.Itoa(*connector.ConnectorId))
	d.Set("name", connector.Name)
	d.Set("description", connector.Description)
	d.Set("is_gov_cloud", connector.IsGovCloud)
	d.Set("last_synced_on", connector.LastSyncedOn)
	d.Set("total_assets", connector.TotalAssets)
	d.Set("state", connector.State)
	d.Set("type", connector.Type)

	// The API is very inconsistent - when you make a Get request the AuthRecord is SOMETIMES not returned!
	// make a few more exactly the same calls and you will get the AuthRecord no problemo!
	// Therefore we're setting up the authRecord only if it comes back in the response
	// The API is very inconsistent - when you make a Get request the AuthRecord is SOMETIMES not returned!
	// make a few more exactly the same calls and you will get the AuthRecord no problemo!
	// Therefore we're setting up the authRecord only if it comes back in the response
	if connector.AuthRecord.DirectoryId != nil {
		d.Set("directory_id", *connector.AuthRecord.DirectoryId)
	}
	if connector.AuthRecord.SubscriptionId != nil {
		d.Set("subscription_id", *connector.AuthRecord.SubscriptionId)
	}
	if connector.AuthRecord.ApplicationId != nil {
		d.Set("application_id", *connector.AuthRecord.ApplicationId)
	}
	//The API doesn't return the auth key (security)

	tagsMap := make(map[string]string)
	if (connector.DefaultTags != nil) && (connector.DefaultTags.TagsList != nil) {
		tagsMap = convertToTagsMap(connector.DefaultTags.TagsList.TagSimple)
	}
	d.Set("default_tags", tagsMap)

	return nil
}

func resourceAssetViewAzureConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	// Due to the way the API works, if the name stays the same, we can't pass it
	var newName *string
	if d.HasChange("name") {
		name := d.Get("name").(string)
		newName = &name
	}

	var ctrIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	defTags, err := setDefaultTags(d)
	if err != nil {
		return err
	}
	connector := client.AssetViewAzureConnector{
		ConnectorId: &ctrIdInt,
		Name:        newName,
		Description: d.Get("description").(string),
		IsGovCloud:  d.Get("is_gov_cloud").(bool),
		DefaultTags: defTags,
	}

	_, err = apiClient.UpdateAssetViewAzureConnector(&connector)
	if err != nil {
		return err
	}

	// Returning Read as for example when updating tags, the API doesn't return the tag names
	return resourceAssetViewAzureConnectorRead(d, m)
}

func setDefaultTags(d *schema.ResourceData) (*client.AssetViewDataTagsChildren, error) {
	var avDefTags *client.AssetViewDataTagsChildren
	if d.Get("default_tags") != nil {
		defTags := d.Get("default_tags").(map[string]interface{})

		var err error
		avDefTags, err = readChildrenTags(defTags)
		if err != nil {
			return nil, err
		}
	}
	return avDefTags, nil
}

func readChildrenTags(childTags map[string]interface{}) (*client.AssetViewDataTagsChildren, error) {
	avChildren := client.AssetViewDataTagsChildren{}
	tagsSimple := make([]client.AssetViewDataTagSimple, len(childTags))
	var i = 0
	for k := range childTags {
		var tagIdInt, err = strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		tagsSimple[i].Id = &tagIdInt
		//The API only allows to set tag Id :(
		i++
	}
	tagsSet := client.AssetViewDataTagsSet{}
	tagsSet.TagSimple = tagsSimple
	avChildren.TagsSet = &tagsSet
	return &avChildren, nil
}

func convertToTagsMap(tags []client.AssetViewDataTagSimple) map[string]string {
	tagsMap := make(map[string]string)
	for i := 0; i < len(tags); i++ {
		tagName := ""
		if &tags[i] != nil {
			tagName = *tags[i].Name
		}
		tagsMap[strconv.Itoa(*tags[i].Id)] = tagName
	}
	return tagsMap
}

func resourceAssetViewAzureConnectorDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	var ctrIdInt, err = strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	err = apiClient.DeleteAssetViewAzureConnector(ctrIdInt)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceAssetViewAzureConnectorExists(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	var ctrIdInt, err = strconv.Atoi(d.Id())

	_, err = apiClient.GetAssetViewAzureConnector(ctrIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
