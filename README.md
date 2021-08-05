# terraform-provider-rsa [![terraform provider](https://img.shields.io/badge/terraform-registry-623CE4)](https://registry.terraform.io/providers/bendrucker/rsa/latest)

Terraform provider for RSA public key encryption. 

```tf
terraform {
  required_providers {
    rsa = {
      source = "bendrucker/rsa"
    }
  }
}

resource "rsa_ciphertext" "example" {
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

output "ciphertext" {
  value = rsa_ciphertext.example.ciphertext
}
```

## Terraform Core

Terraform includes an [`rsadecrypt`](https://www.terraform.io/docs/language/functions/rsadecrypt.html) function, so why not `rsaencrypt`? RSA implementations [use randomized padding](https://en.wikipedia.org/wiki/RSA_(cryptosystem)#Padding_schemes) which results in non-deterministic output. For a given plaintext, encrypting will result in different ciphertext each time. Terraform functions are stateless and should generally be idempotent to avoid perpetual diffs. Implementing a resource provides the necessary state tracking to ensure that the ciphertext is only re-computed when the plaintext or other encryption attributes have changed.
