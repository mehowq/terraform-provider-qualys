package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func dataSourceAzureConnector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
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
		},
		Read:   dataSourceAzureConnectorRead,
	}
}

func dataSourceAzureConnectorRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connectorId := d.Get("connector_id").(string)
	connector, err := apiClient.GetAzureConnector(connectorId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Connector with ID %s", connectorId)
		}
	}

	d.SetId(connector.ConnectorId)
	d.Set("name", connector.Name)
	d.Set("description", connector.Description)
	d.Set("directory_id", connector.DirectoryId)
	d.Set("subscription_id", connector.SubscriptionId)
	d.Set("application_id", connector.ApplicationId)
	d.Set("is_gov_cloud", connector.IsGovCloud)
	d.Set("last_synced_on", connector.LastSyncedOn)
	d.Set("total_assets", connector.TotalAssets)
	d.Set("state", connector.State)
	return nil
}