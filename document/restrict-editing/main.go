package main

import (
	"fmt"
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

	// Restrict Editing: allow no edits, enforced with a password.
	prot := doc.Settings.DocumentProtection()
	prot.SetEdit(wml.ST_DocProtectReadOnly)
	prot.Enforce(true)
	prot.SetPassword("unioffice")

	if err := doc.SaveToFile("restrict-editing.docx"); err != nil {
		fmt.Println(err)
	}
}
