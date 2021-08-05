package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiphertext() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"plaintext": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"padding": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PKCS1.5",
				ValidateFunc: validation.StringInSlice([]string{"PKCS1.5", "OAEP"}, true),
			},

			"hash": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "SHA256",
				ValidateFunc: validation.StringInSlice([]string{"SHA256", "SHA512"}, true),
			},

			"ciphertext": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
