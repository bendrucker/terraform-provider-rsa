package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				"rsa_ciphertext": resourceCiphertext(),
			},
		}
	}
}
