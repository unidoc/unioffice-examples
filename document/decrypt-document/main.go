/*
 * Copyright 2026 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates opening an encrypted DOCX: the errors returned
 * without a password and with a wrong one, opening it with the right password,
 * and saving an unencrypted copy.
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2"
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
	if err := createEncrypted("encrypted.docx", "unioffice"); err != nil {
		log.Fatal(err)
	}

	// document.Open cannot read an encrypted file and says so.
	_, err := document.Open("encrypted.docx")
	switch {
	case errors.Is(err, unioffice.ErrEncryptedDocument):
		fmt.Println("document.Open:", err)
	case err != nil:
		log.Fatal(err)
	default:
		log.Fatal("expected an encrypted-document error")
	}

	// A wrong password is reported separately from other read errors.
	_, err = document.OpenWithOptions("encrypted.docx", &crypto.OpenOptions{Password: "wrong"})
	switch {
	case errors.Is(err, unioffice.ErrIncorrectPassword):
		fmt.Println("wrong password:", err)
	case err != nil:
		log.Fatal(err)
	default:
		log.Fatal("expected an incorrect-password error")
	}

	doc, err := document.OpenWithOptions("encrypted.docx", &crypto.OpenOptions{Password: "unioffice"})
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			fmt.Println("text:", r.Text())
		}
	}

	// Saving without options writes an ordinary, unencrypted copy.
	if err := doc.SaveToFile("decrypted.docx"); err != nil {
		log.Fatal(err)
	}
}

func createEncrypted(path, password string) error {
	doc := document.New()
	defer doc.Close()
	doc.AddParagraph().AddRun().AddText("Only readers with the password can see this.")
	return doc.SaveToFileWithOptions(path, &crypto.SaveOptions{Password: password})
}
