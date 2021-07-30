// Copyright 2017 FoxyUtils ehf. All rights reserved.
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

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	defer doc.Close()

	para := doc.AddParagraph()
	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)
	}

	// Set watermark to doc.
	watermark := doc.AddWatermarkText("TEST")

	// Change style of watermark text.
	// Currently having 2 function,
	// SetTextStyleBold and SetTextStyleItalic.
	watermark.SetTextStyleBold(true)
	watermark.SetTextStyleItalic(false)

	doc.SaveToFile("watermark-text.docx")
}
