// Copyright 2018 FoxyUtils ehf. All rights reserved.
package main

import (
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
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
