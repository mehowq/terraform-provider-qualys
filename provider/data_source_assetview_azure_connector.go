package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func dataSourceAssetViewAzureConnector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"directory_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_gov_cloud": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
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
		Read: dataSourceAssetViewAzureConnectorRead,
	}
}

func dataSourceAssetViewAzureConnectorRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	subscriptionId := d.Get("subscription_id").(string)
	connectorId := d.Get("connector_id").(int)

	var connector *client.AssetViewAzureConnector
	if subscriptionId != "" {
		crSubId := client.AssetViewFiltersCriteria{
			Field:    "authRecord.subscriptionId",
			Operator: "EQUALS",
			Value:    subscriptionId,
		}
		criteria := []client.AssetViewFiltersCriteria{crSubId}
		connectors, err := apiClient.SearchAssetViewAzureConnectors(criteria)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				d.SetId("")
			} else {
				return fmt.Errorf("error finding AssetView Azure Connector with subscription_id %s", subscriptionId)
			}
		} else {
			if (len(*connectors)) == 1 {
				connector = &(*connectors)[0]
			} else {
				return fmt.Errorf("AssetView Azure Connector with subscription_id %s doesn't exist", subscriptionId)
			}
		}
	} else {
		var err error
		connector, err = apiClient.GetAssetViewAzureConnector(connectorId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				d.SetId("")
			} else {
				return fmt.Errorf("error finding AssetView Azure Connector with ID %d", connectorId)
			}
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
