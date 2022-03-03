package provider

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/domains"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/invitations"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/synchronizationjobs"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/users"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{
		administrativeunits.Registration{},
		applications.Registration{},
		approleassignments.Registration{},
		conditionalaccess.Registration{},
		directoryroles.Registration{},
		domains.Registration{},
		groups.Registration{},
		invitations.Registration{},
		serviceprincipals.Registration{},
		synchronizationjobs.Registration{},
		users.Registration{},
	}
}
