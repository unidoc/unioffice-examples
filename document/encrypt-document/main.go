package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/crypto"
	"github.com/unidoc/unioffice/v2/document"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	doc := document.New()
	defer doc.Close()
	doc.AddParagraph().AddRun().AddText("This document is encrypted with a password.")

	// SaveOptions carries the password and, optionally, the cipher strength.
	// The zero value of Encryption (EncryptionDefault) is agile AES-256.
	opts := &crypto.SaveOptions{
		Password:   "unioffice",
		Encryption: crypto.EncryptionAES256,
	}
	if err := doc.SaveToFileWithOptions("encrypted.docx", opts); err != nil {
		fmt.Println(err)
		return
	}

	// Weaker keys for compatibility. All are agile encryption and open in
	// Office 2010 and later.
	for name, alg := range map[string]crypto.EncryptionAlgorithm{
		"aes128": crypto.EncryptionAES128,
		"aes192": crypto.EncryptionAES192,
	} {
		d := document.New()
		d.AddParagraph().AddRun().AddText("Encrypted with " + name + ".")
		if err := d.SaveToFileWithOptions("encrypted-"+name+".docx", &crypto.SaveOptions{Password: "unioffice", Encryption: alg}); err != nil {
			fmt.Println(err)
		}
		d.Close()
	}

	// Read an encrypted document back with the password.
	got, err := document.OpenWithOptions("encrypted.docx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer got.Close()
	fmt.Println("opened encrypted document")
}
