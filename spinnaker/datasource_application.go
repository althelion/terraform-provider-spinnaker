package spinnaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceApplication() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validateApplicationName,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"accounts": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"instance_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
		Read: resourceApplicationRead,
	}
}
