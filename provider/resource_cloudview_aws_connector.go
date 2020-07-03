package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func resourceCloudViewAWSConnector() *schema.Resource {
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
			"aws_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_gov_cloud": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_china_region": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_portal_connector": &schema.Schema{
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
		Create: resourceCloudViewAWSConnectorCreate,
		Read:   resourceCloudViewAWSConnectorRead,
		Update: resourceCloudViewAWSConnectorUpdate,
		Delete: resourceCloudViewAWSConnectorDelete,
		Exists: resourceCloudViewAWSConnectorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudViewAWSConnectorCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connector := client.CloudViewAWSConnector{
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		ARN:               d.Get("arn").(string),
		ExternalId:        d.Get("external_id").(string),
		IsGovCloud:        d.Get("is_gov_cloud").(bool),
		IsChinaRegion:     d.Get("is_china_region").(bool),
		IsPortalConnector: d.Get("is_portal_connector").(bool),
	}

	//TODO Add some validation to check if account_id is not already in use

	newConnector, err := apiClient.NewCloudViewAWSConnector(&connector)

	if err != nil {
		return err
	}
	d.SetId(newConnector.ConnectorId)
	d.Set("aws_account_id", newConnector.AWSAccountId)
	d.Set("last_synced_on", newConnector.LastSyncedOn)
	d.Set("total_assets", newConnector.TotalAssets)
	d.Set("state", newConnector.State)
	return resourceCloudViewAWSConnectorRead(d, m)
}

func resourceCloudViewAWSConnectorRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connectorId := d.Id()
	connector, err := apiClient.GetCloudViewAWSConnector(connectorId)
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
	d.Set("aws_account_id", connector.AWSAccountId)
	d.Set("arn", connector.ARN)
	d.Set("external_id", connector.ExternalId)
	d.Set("is_gov_cloud", connector.IsGovCloud)
	d.Set("is_china_region", connector.IsChinaRegion)
	d.Set("is_portal_connector", connector.IsPortalConnector)
	d.Set("last_synced_on", connector.LastSyncedOn)
	d.Set("total_assets", connector.TotalAssets)
	d.Set("state", connector.State)
	return nil
}

func resourceCloudViewAWSConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connector := client.CloudViewAWSConnector{
		ConnectorId:       d.Id(),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		AWSAccountId:      d.Get("aws_account_id").(string),
		ARN:               d.Get("arn").(string),
		ExternalId:        d.Get("external_id").(string),
		IsGovCloud:        d.Get("is_gov_cloud").(bool),
		IsChinaRegion:     d.Get("is_china_region").(bool),
		IsPortalConnector: d.Get("is_portal_connector").(bool),
	}

	err := apiClient.UpdateCloudViewAWSConnector(&connector)
	if err != nil {
		return err
	}
	return nil
}

func resourceCloudViewAWSConnectorDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	connectorId := d.Id()

	err := apiClient.DeleteCloudViewAWSConnector(connectorId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceCloudViewAWSConnectorExists(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	connectorId := d.Id()
	_, err := apiClient.GetCloudViewAWSConnector(connectorId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}