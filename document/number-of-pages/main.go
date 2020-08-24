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
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()
	fmt.Println("Total number of pages in the opened document:", doc.AppProperties.Pages())
	newDoc := document.New()
	fmt.Println("Total number of pages in the new document:", newDoc.AppProperties.Pages())
}
