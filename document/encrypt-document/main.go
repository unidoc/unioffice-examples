/*
 * Copyright 2026 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates saving a DOCX encrypted with a password, choosing
 * the AES key size, and opening the encrypted file again with the password.
 */

package main

import (
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	// AES-128 and AES-192 are available for policies or readers that require
	// a specific key size.
	for _, c := range []struct {
		name string
		alg  crypto.EncryptionAlgorithm
	}{
		{"aes128", crypto.EncryptionAES128},
		{"aes192", crypto.EncryptionAES192},
	} {
		d := document.New()
		d.AddParagraph().AddRun().AddText("Encrypted with " + c.name + ".")
		err := d.SaveToFileWithOptions("encrypted-"+c.name+".docx", &crypto.SaveOptions{Password: "unioffice", Encryption: c.alg})
		d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Read an encrypted document back with the password.
	got, err := document.OpenWithOptions("encrypted.docx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		log.Fatal(err)
	}
	defer got.Close()
	for _, p := range got.Paragraphs() {
		for _, r := range p.Runs() {
			fmt.Println("text:", r.Text())
		}
	}
}
