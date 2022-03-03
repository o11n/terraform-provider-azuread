package synchronizationjobs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs"
)

var validStatusValues = [...]string{"ACTIVE", "PAUSED"} // TODO: What about "restart"?

func synchronizationJobResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: synchronizationJobResourceCreate,
		ReadContext:   synchronizationJobResourceRead,
		UpdateContext: synchronizationJobResourceUpdate,
		DeleteContext: synchronizationJobResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this job should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"template_id": {
				Description: "The id of the template that this synchronization job should be based on",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"status": {
				Description: "The desired status of the synchronization job",
				Type:        schema.TypeString,
				Required:    true,
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					value := v.(string)
					for _, validValue := range validStatusValues {
						if value == validValue {
							return diag.Diagnostics{}
						}
					}
					return diag.Diagnostics{
						diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Invalid value for status",
							Detail:   "The value of status must be one of ACTIVE or PAUSED",
						},
					}
				},
			},
		},
	}
}

func synchronizationJobResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).MSGraph
	svcpId := d.Get("service_principal_id").(string)
	templateId := d.Get("template_id").(string)
	synchronization := client.ServicePrincipalsById(svcpId).Synchronization()

	object := graph.NewSynchronizationJob()
	object.SetTemplateId(&templateId)
	options := &jobs.JobsRequestBuilderPostOptions{
		Body: object,
	}
	job, err := synchronization.Jobs().Post(options)
	if err != nil {
		return tf.ErrorDiagF(err, "creating synchronization job for service principal ID %q", svcpId)
	}

	jobId := *job.GetId()
	d.SetId(jobId)
	tf.Set(d, "template_id", job.GetTemplateId())

	if d.Get("status") == "Active" && job.GetStatus().GetCode().String() != "Active" {
		log.Printf("[DEBUG] Starting synchronization job with ID %q and service principal ID %q", jobId, svcpId)
		err = synchronization.JobsById(jobId).Start().Post(nil)
		if err != nil {
			return tf.ErrorDiagF(err, "starting synchronization job ID %q for service principal ID %q", jobId, svcpId)
		}
	}

	job, err = synchronization.JobsById(jobId).Get(nil)
	if err != nil {
		return tf.ErrorDiagF(err, "retrieving synchronization job with ID %q and service principal ID %q", jobId, svcpId)
	}

	tf.Set(d, "status", job.GetStatus().GetCode().String())

	return nil
}

func synchronizationJobResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).MSGraph
	jobId := d.Id()
	svcpId := d.Get("service_principal_id").(string)

	job, err := client.ServicePrincipalsById(svcpId).Synchronization().JobsById(jobId).Get(nil)
	if err != nil {
		return tf.ErrorDiagF(err, "retrieving synchronization job with ID %q and service principal ID %q", jobId, svcpId)
	}

	tf.Set(d, "template_id", job.GetTemplateId())
	tf.Set(d, "status", job.GetStatus().GetCode().String())

	return nil
}

func synchronizationJobResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).MSGraph
	jobId := d.Id()
	svcpId := d.Get("service_principal_id").(string)
	job := client.ServicePrincipalsById(svcpId).Synchronization().JobsById(jobId)

	if d.HasChange("status") {
		switch d.Get("status").(string) {
		case "Active":
			log.Printf("[DEBUG] Starting synchronization job with ID %q and service principal ID %q", jobId, svcpId)
			if err := job.Start().Post(nil); err != nil {
				return tf.ErrorDiagF(err, "starting synchronization job with ID %q and service principal ID %q", jobId, svcpId)
			}
		case "Paused":
			log.Printf("[DEBUG] Pausing synchronization job with ID %q and service principal ID %q", jobId, svcpId)
			if err := job.Pause().Post(nil); err != nil {
				return tf.ErrorDiagF(err, "pausing synchronization job with ID %q and service principal ID %q", jobId, svcpId)
			}
		}
	}

	return nil
}

func synchronizationJobResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).MSGraph
	svcpId := d.Get("service_principal_id").(string)
	jobId := d.Id()

	job := client.ServicePrincipalsById(svcpId).Synchronization().JobsById(jobId)
	if err := job.Delete(nil); err != nil {
		return tf.ErrorDiagF(err, "deleting synchronization job with ID %q and service principal ID %q", jobId, svcpId)
	}

	return nil
}
