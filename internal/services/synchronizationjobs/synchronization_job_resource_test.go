package synchronizationjobs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type synchronizationJobResource struct{}

func TestAccSynchonizationJob_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_synchronization_job", "test")
	r := synchronizationJobResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("template_id").HasValue("scim"),
				check.That(data.ResourceName).Key("status").HasValue("ACTIVE"),
			),
		},
	})
}

func (r synchronizationJobResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.MSGraph
	svcpId := state.Attributes["service_principal_id"]

	job, err := client.ServicePrincipalsById(svcpId).Synchronization().JobsById(state.ID).Get(nil)
	if err != nil {
		return nil, fmt.Errorf("retrieving synchronization job with ID %q and service principal ID %q: %w", state.ID, svcpId, err)
	}

	return utils.Bool(job.GetId() == &state.ID), nil
}

func (r synchronizationJobResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application_template" "test" {
  display_name = "Custom"
}

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
	template_id  = data.azuread_application_template.test.id
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_synchronization_job" "this" {
  service_principal_id = azuread_service_principal.test.id
  template_id          = "scim"
  status               = "ACTIVE"
}
`, data.RandomInteger)
}
