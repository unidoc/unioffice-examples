/*
 * Copyright 2026 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates saving a PPTX encrypted with a password and
 * opening it again with the password.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/crypto"
	"github.com/unidoc/unioffice/v2/presentation"
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
	ppt := presentation.New()
	defer ppt.Close()
	ppt.AddSlide().AddTextBox().AddParagraph().AddRun().SetText("This presentation is encrypted with a password.")

	// The zero value of Encryption (EncryptionDefault) is agile AES-256;
	// EncryptionAES192 and EncryptionAES128 select a smaller key if required.
	opts := &crypto.SaveOptions{
		Password:   "unioffice",
		Encryption: crypto.EncryptionAES256,
	}
	if err := ppt.SaveToFileWithOptions("encrypted.pptx", opts); err != nil {
		log.Fatal(err)
	}

	got, err := presentation.OpenWithOptions("encrypted.pptx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		log.Fatal(err)
	}
	defer got.Close()
	fmt.Println("slides:", len(got.Slides()))
	fmt.Println("text:", got.ExtractText().Text())
}
