package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func resourceAzureConnector() *schema.Resource {
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
			"directory_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authentication_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
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
		},
		Create: resourceAzureConnectorCreate,
		Read:   resourceAzureConnectorRead,
		Update: resourceAzureConnectorUpdate,
		Delete: resourceAzureConnectorDelete,
		Exists: resourceAzureConnectorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceAzureConnectorCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connector := client.AzureConnector{
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		DirectoryId:       d.Get("directory_id").(string),
		SubscriptionId:    d.Get("subscription_id").(string),
		ApplicationId:     d.Get("application_id").(string),
		AuthenticationKey: d.Get("authentication_key").(string),
		IsGovCloud:        d.Get("is_gov_cloud").(bool),
	}

	newConnector, err := apiClient.NewAzureConnector(&connector)

	if err != nil {
		return err
	}
	d.SetId(newConnector.ConnectorId)
	d.Set("last_synced_on", newConnector.LastSyncedOn)
	d.Set("total_assets", newConnector.TotalAssets)
	d.Set("state", newConnector.State)
	return resourceAzureConnectorRead(d, m)
}

func resourceAzureConnectorRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connectorId := d.Id()
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
	//The API doesn't return the auth key (security)
	//d.Set("authentication_key", connector.AuthenticationKey)
	d.Set("is_gov_cloud", connector.IsGovCloud)
	d.Set("last_synced_on", connector.LastSyncedOn)
	d.Set("total_assets", connector.TotalAssets)
	d.Set("state", connector.State)
	return nil
}

func resourceAzureConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connector := client.AzureConnector{
		ConnectorId:    d.Id(),
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		DirectoryId:    d.Get("directory_id").(string),
		SubscriptionId: d.Get("subscription_id").(string),
		ApplicationId:  d.Get("application_id").(string),
		//Since Read returns an empty auth key, TF thinks it needs to update it
		AuthenticationKey: d.Get("authentication_key").(string),
		IsGovCloud:        d.Get("is_gov_cloud").(bool),
	}

	err := apiClient.UpdateAzureConnector(&connector)
	if err != nil {
		return err
	}
	return nil
}

func resourceAzureConnectorDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connectorId := d.Id()

	err := apiClient.DeleteAzureConnector(connectorId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceAzureConnectorExists(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	connectorId := d.Id()
	_, err := apiClient.GetAzureConnector(connectorId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
