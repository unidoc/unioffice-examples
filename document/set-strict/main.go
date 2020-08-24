// Copyright 2020 FoxyUtils ehf. All rights reserved.

package main

import (
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
		panic(err)
	}
	defer doc.Close()
	doc.SetStrict(false) // document will be saved as Word document (this is a default option for new files)
	doc.SaveToFile("conformance_transitional.docx")
	doc.SetStrict(true) // document will be saved in the Strict mode
	doc.SaveToFile("conformance_strict.docx")
}
