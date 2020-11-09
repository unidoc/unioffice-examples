// Copyright 2017 FoxyUtils ehf. All rights reserved.
// This example demonstrates merging two documents into one.

package main

import (
	"log"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	doc0, err := document.Open("document0.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc0.Close()
	doc1, err := document.Open("document1.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc1.Close()
	doc0.AddParagraph().AddRun().AddPageBreak()

	err = doc0.Append(doc1)
	if err != nil {
		log.Fatalf("error appending document: %s", err)
	}
	doc0.SaveToFile("merged.docx")
}
