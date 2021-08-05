package provider

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"hash"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiphertext() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCiphertextCreate,
		Schema: map[string]*schema.Schema{
			"plaintext": {
				Description: "The plaintext to encrypt",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"public_key": {
				Description:      "The public key used for encryption, in PEM format",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validatePublicKey,
			},

			"padding": {
				Description:      "The padding mode to use",
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "PKCS1.5",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"PKCS1.5", "OAEP"}, true)),
			},

			"hash": {
				Description:      "The hash algorithm to use, for OAEP only",
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "SHA256",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"SHA256", "SHA512"}, true)),
			},

			"ciphertext": {
				Description: "The encrypted ciphertext, base64 encoded",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceCiphertextCreate(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	plaintext := d.Get("plaintext").(string)
	publicKey := d.Get("public_key").(string)
	padding := d.Get("padding").(string)
	hashName := d.Get("hash").(string)

	block, _ := pem.Decode([]byte(publicKey))

	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return diag.Errorf("failed to parse public key: %v", err)
	}

	var b []byte

	switch padding {
	case "PKCS1.5":
		b, err = rsa.EncryptPKCS1v15(rand.Reader, key, []byte(plaintext))
		if err != nil {
			return diag.Errorf("failed to encrypt with PKCS1.5: %v", err)
		}

	case "OAEP":
		b, err = rsa.EncryptOAEP(getHash(hashName), rand.Reader, key, []byte(plaintext), []byte(nil))
		if err != nil {
			return diag.Errorf("failed to encrypt with OAEP: %v", err)
		}
	}

	d.Set("ciphertext", base64.StdEncoding.EncodeToString(b))
	return nil
}

func getHash(name string) hash.Hash {
	switch name {
	case "SHA256":
		return sha256.New()

	case "SHA512":
		return sha512.New()
	}

	return nil
}

func validatePublicKey(v interface{}, path cty.Path) (diags diag.Diagnostics) {
	block, _ := pem.Decode([]byte(v.(string)))

	if block == nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Invalid public key",
			Detail:        "The public key must be in PEM format. No PEM-encoded blocks were found.",
			AttributePath: path,
		})

		return diags
	}

	_, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Invalid public key",
			Detail:        fmt.Errof("The public key must be in PEM format and be a valid RSA public key. Error: %v", err),
			AttributePath: path,
		})
	}

	return diags
}
