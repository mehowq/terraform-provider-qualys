package provider

import (
	"fmt"
	"regexp"
	"testing"

	guuid "github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

const testAccCvAzConnResType = "qualys_cloudview_azure_connector"
const testAccCvAzConnResName = "testacc_cv_az"
const testAccCvAzConnResName2 = "testacc_cv_az2"

var testAccCvAzConnName = fmt.Sprintf("TF_AccTest_CV_AzConn_%s", guuid.New().String())
var testAccCvAzConnDirId = guuid.New().String()
var testAccCvAzConnSubId = guuid.New().String()
var testAccCvAzConnAppId = guuid.New().String()
var testAccCvAzConnNamePostUpd = fmt.Sprintf("TF_AccTest_CV_AzConn_%s", guuid.New().String())
var testAccCvAzConnDirIdPostUpd = guuid.New().String()
var testAccCvAzConnSubIdPostUpd = guuid.New().String()
var testAccCvAzConnAppIdPostUpd = guuid.New().String()

func TestAccCvAzConnector_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCvAzConnectorBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCvAzConnectorExists(fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "name", testAccCvAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "directory_id", testAccCvAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "subscription_id", testAccCvAzConnSubId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "application_id", testAccCvAzConnAppId),
				),
			},
		},
	})
}

func TestAccCvAzConnector_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCvAzConnectorUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCvAzConnectorExists(fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "name", testAccCvAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "directory_id", testAccCvAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "application_id", testAccCvAzConnAppId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "subscription_id", testAccCvAzConnSubId),
				),
			},
			{
				Config: testAccCheckCvAzConnectorUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCvAzConnectorExists(fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "name", testAccCvAzConnNamePostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "directory_id", testAccCvAzConnDirIdPostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "application_id", testAccCvAzConnAppIdPostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName), "subscription_id", testAccCvAzConnSubIdPostUpd),
				),
			},
		},
	})
}

func TestAccCvAzConnector_Multiple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCvAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCvAzConnectorMultiple(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCvAzConnectorExists(fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName)),
					testAccCheckCvAzConnectorExists(fmt.Sprintf("%s.%s", testAccCvAzConnResType, testAccCvAzConnResName2)),
				),
			},
		},
	})
}

func testAccCheckCvAzConnectorBasic() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Cloud View Azure Connector with random subscription and authentication details. Basic Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
		testAccCvAzConnResType,
		testAccCvAzConnResName,
		testAccCvAzConnName,
		testAccCvAzConnDirId,
		testAccCvAzConnSubId,
		testAccCvAzConnAppId,
		authKey))
}

func testAccCheckCvAzConnectorExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No CloudView Azure Connector ID is set for %s", resource)
		}
		connectorId := rs.Primary.ID
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetCloudViewAzureConnector(connectorId)
		if err != nil {
			return fmt.Errorf("error fetching CloudView Azure Connector %s with ID %s: %s", resource, connectorId, err)
		}
		return nil
	}
}

func testAccCheckCvAzConnectorDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccCvAzConnResType {
			continue
		}

		connectorId := rs.Primary.ID
		_, err := apiClient.GetCloudViewAzureConnector(connectorId)
		if err == nil {
			return fmt.Errorf("CloudView Azure Connector %s still exists", connectorId)
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccCheckCvAzConnectorUpdatePre() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Cloud View Azure Connector with random subscription and authentication details. Pre Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
		testAccCvAzConnResType,
		testAccCvAzConnResName,
		testAccCvAzConnName,
		testAccCvAzConnDirId,
		testAccCvAzConnSubId,
		testAccCvAzConnAppId,
		authKey))
}

func testAccCheckCvAzConnectorUpdatePost() string {
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Cloud View Azure Connector with random subscription and authentication details. Post Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
		testAccCvAzConnResType,
		testAccCvAzConnResName,
		testAccCvAzConnNamePostUpd,
		testAccCvAzConnDirIdPostUpd,
		testAccCvAzConnSubIdPostUpd,
		testAccCvAzConnAppIdPostUpd,
		authKey))
}

func testAccCheckCvAzConnectorMultiple() string {
	authKey := guuid.New().String()
	conn2name := fmt.Sprintf("TF_AccTest_AzConn2_%s", guuid.New().String())
	conn2subId := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Cloud View Azure Connector with random subscription and authentication details. Multiple 1 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}

resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Cloud View Azure Connector with random subscription and authentication details. Multiple 2 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
		testAccCvAzConnResType,
		testAccCvAzConnResName,
		testAccCvAzConnName,
		testAccCvAzConnDirId,
		testAccCvAzConnSubId,
		testAccCvAzConnAppId,
		authKey,
		testAccCvAzConnResType,
		testAccCvAzConnResName2,
		conn2name,
		testAccCvAzConnDirId,
		conn2subId,
		testAccCvAzConnAppId,
		authKey))
}
