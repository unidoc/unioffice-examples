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
	d, err := document.Open("mm.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer d.Close()
	for _, v := range d.MergeFields() {
		log.Println("replacing", v)
	}
	rep := map[string]string{}
	rep["Title"] = "mr."      // has a \* Upper attribute on the field
	rep["FirstName"] = "JOHN" // has a \* Lower attribute on the field
	d.MailMerge(rep)
	d.SaveToFile("merged.docx")
}
