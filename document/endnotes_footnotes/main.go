// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
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
	doc, err := document.Open("footnotes_endnotes.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	if doc.HasFootnotes() {
		fmt.Printf("Document has %02d footnotes.\n", len(doc.Footnotes()))
	} else {
		fmt.Println("Document has no footnotes")
	}

	if doc.HasEndnotes() {
		fmt.Printf("Document has %02d endnotes.\n", len(doc.Endnotes()))
	} else {
		fmt.Println("Document has no endnotes")
	}
}
