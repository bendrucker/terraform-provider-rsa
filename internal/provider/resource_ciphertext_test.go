package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCiphertext(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCiphertext,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("rsa_ciphertext.test", "ciphertext"),
				),
			},
		},
	})
}

const testAccResourceCiphertext = `
	resource "rsa_ciphertext" "test" {
		plaintext = "Hello World"
		
		public_key = <<-PEM
			-----BEGIN PUBLIC KEY-----
			MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0u6+kZ7yoy1IMUjfoDY+
			WwaTtfaQyLmYM/bs/CtekvDFbSQnOHIunnLKFo8OQW/PeLBur+BIcfzS5spVpXB5
			07P3yzf/mUYwX3sdy1Zu3JdcWKKho793niIAdKNQg48xipCniVg6J5l3WK5816KB
			Dc2+Bjwer2z5cE9G1pUPRnK3m0uHrVsFxmMnk38RZcZnGmokoBzMjUa/2w1kCHuD
			Eq3kdSHvLBmmo5bP9OHHV9F4KVlB8cDp3TSc74U0BVEUDe3BBf9VgXfvqhjDTRJh
			lpC+QxgdBj958K/h8BnRB6vkW3l5OXirowyXg4ZAWQn0XJ+lby5w7yCg4HetyYH5
			0QIDAQAB
			-----END PUBLIC KEY-----
		PEM
	}
`
