/*
 * Copyright 2026 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates Word's Restrict Editing: making a DOCX read-only
 * with a password. This is not encryption. The content stays readable by
 * anyone, and the password only gates Word's "Stop Protection". Use
 * document/encrypt-document to keep the content secret.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"
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

	para := doc.AddParagraph()
	para.AddRun().AddText("This document is read-only until protection is stopped with the password.")

	// Allow no edits, enforced with a password. SetPassword stores a salted
	// SHA-512 hash, not the password itself.
	prot := doc.Settings.DocumentProtection()
	prot.SetEdit(wml.ST_DocProtectReadOnly)
	prot.Enforce(true)
	prot.SetPassword("unioffice")

	fmt.Println("password matches:", prot.VerifyPassword("unioffice"))
	fmt.Println("wrong password matches:", prot.VerifyPassword("wrong"))

	if err := doc.SaveToFile("restrict-editing.docx"); err != nil {
		log.Fatal(err)
	}
}
