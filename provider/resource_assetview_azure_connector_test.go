package provider

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	guuid "github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

const testAccAvAzConnResType = "qualys_assetview_azure_connector"
const testAccAvAzConnResName = "testacc_av_az"
const testAccAvAzConnResName2 = "testacc_av_az2"

var testAccAvAzConnName = fmt.Sprintf("TF_AccTest_AV_AzConn_%s", guuid.New().String())
var testAccAvAzConnDirId = guuid.New().String()
var testAccAvAzConnSubId = guuid.New().String()
var testAccAvAzConnAppId = guuid.New().String()
var testAccAvAzConnDefTags = "{\"16568284\" = \"Cloud Agent\", \"17367145\" = \"AZURE-CONNECTOR\"}"
var testAccAvAzConnNamePostUpd = fmt.Sprintf("TF_AccTest_AV_AzConn_%s", guuid.New().String())
var testAccAvAzConnDirIdPostUpd = guuid.New().String()
var testAccAvAzConnSubIdPostUpd = guuid.New().String()
var testAccAvAzConnAppIdPostUpd = guuid.New().String()
var testAccAvAzConnDefTagsPostUpd = "{\"16568284\" = \"Cloud Agent\", \"17367145\" = \"AZURE-CONNECTOR\", \"16974467\" = \"Azure-POC\"}"

func testAccAvAzConnector_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAvAzConnectorBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAvAzConnectorExists(fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "name", testAccAvAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "directory_id", testAccAvAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "subscription_id", testAccAvAzConnSubId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "application_id", testAccAvAzConnAppId),
				),
			},
		},
	})
}

func TestAccAvAzConnector_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAvAzConnectorUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAvAzConnectorExists(fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "name", testAccAvAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "directory_id", testAccAvAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "application_id", testAccAvAzConnAppId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "subscription_id", testAccAvAzConnSubId),
				),
			},
			{
				Config: testAccCheckAvAzConnectorUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAvAzConnectorExists(fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "name", testAccAvAzConnNamePostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "directory_id", testAccAvAzConnDirIdPostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "application_id", testAccAvAzConnAppIdPostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName), "subscription_id", testAccAvAzConnSubIdPostUpd),
				),
			},
		},
	})
}

func TestAccAvAzConnector_Multiple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAvAzConnectorMultiple(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAvAzConnectorExists(fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName)),
					testAccCheckAvAzConnectorExists(fmt.Sprintf("%s.%s", testAccAvAzConnResType, testAccAvAzConnResName2)),
				),
			},
		},
	})
}

func testAccCheckAvAzConnectorBasic() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Asset View Azure Connector with random subscription and authentication details. Basic Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
	default_tags = {}
}
	`,
		testAccAvAzConnResType,
		testAccAvAzConnResName,
		testAccAvAzConnName,
		testAccAvAzConnDirId,
		testAccAvAzConnSubId,
		testAccAvAzConnAppId,
		authKey))
}

func testAccCheckAvAzConnectorExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AssetView Azure Connector ID is set for %s", resource)
		}
		var connIdInt, err = strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err = apiClient.GetAssetViewAzureConnector(connIdInt)
		if err != nil {
			return fmt.Errorf("error fetching AssetView Azure Connector %s with ID %d: %s", resource, connIdInt, err)
		}
		return nil
	}
}

func testAccCheckAvAzConnectorDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccAvAzConnResType {
			continue
		}

		var connIdInt, err = strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = apiClient.GetAssetViewAzureConnector(connIdInt)
		if err == nil {
			return fmt.Errorf("AssetView Azure Connector %d still exists", connIdInt)
		}
		notFoundErr := "INVALID_REQUEST - Unable to find"
		notFoundErr2 := "NOT_FOUND - AssetDataConnector"
		if (!strings.Contains(err.Error(), notFoundErr)) && (!strings.Contains(err.Error(), notFoundErr2)) {
			return fmt.Errorf("expected %s or %s, got %s", notFoundErr, notFoundErr2, err)
		}
	}

	return nil
}

func testAccCheckAvAzConnectorUpdatePre() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Asset View Azure Connector with random subscription and authentication details. Pre Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
	default_tags = %s
}
	`,
		testAccAvAzConnResType,
		testAccAvAzConnResName,
		testAccAvAzConnName,
		testAccAvAzConnDirId,
		testAccAvAzConnSubId,
		testAccAvAzConnAppId,
		authKey,
		testAccAvAzConnDefTags))
}

func testAccCheckAvAzConnectorUpdatePost() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Asset View Azure Connector with random subscription and authentication details. Post Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
	default_tags = %s
}
	`,
		testAccAvAzConnResType,
		testAccAvAzConnResName,
		testAccAvAzConnNamePostUpd,
		testAccAvAzConnDirIdPostUpd,
		testAccAvAzConnSubIdPostUpd,
		testAccAvAzConnAppIdPostUpd,
		authKey,
		testAccAvAzConnDefTagsPostUpd))
}

func testAccCheckAvAzConnectorMultiple() string {
	authKey := guuid.New().String()
	conn2name := fmt.Sprintf("TF_AccTest_AzConn2_%s", guuid.New().String())
	conn2subId := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Asset View Azure Connector with random subscription and authentication details. Multiple 1 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
	default_tags = {}
}

resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Asset View Azure Connector with random subscription and authentication details. Multiple 2 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
	default_tags = {}
}
	`,
		testAccAvAzConnResType,
		testAccAvAzConnResName,
		testAccAvAzConnName,
		testAccAvAzConnDirId,
		testAccAvAzConnSubId,
		testAccAvAzConnAppId,
		authKey,
		testAccAvAzConnResType,
		testAccAvAzConnResName2,
		conn2name,
		testAccAvAzConnDirId,
		conn2subId,
		testAccAvAzConnAppId,
		authKey))
}
