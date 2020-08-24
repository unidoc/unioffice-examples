// Copyright 2018 FoxyUtils ehf. All rights reserved.
package main

import (
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
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

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	defer doc.Close()

	ftr := doc.AddFooter()
	para := ftr.AddParagraph()
	para.Properties().AddTabStop(3*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)

	run := para.AddRun()
	run.AddTab()
	run.AddFieldWithFormatting(document.FieldCurrentPage, "", false)
	run.AddText(" of ")
	run.AddFieldWithFormatting(document.FieldNumberOfPages, "", false)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)

	for i := 0; i < 20; i++ {
		para := doc.AddParagraph()
		for j := 0; j < 5; j++ {
			run := para.AddRun()
			run.AddText(lorem)
		}
	}
	doc.SaveToFile("page-numbers.docx")
}
