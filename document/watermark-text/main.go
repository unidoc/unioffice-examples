// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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

	para := doc.AddParagraph()
	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)
	}

	// Set watermark to doc.
	watermark := doc.AddWatermarkText("TEST")

	// Change style of watermark text.
	watermark.SetTextStyleBold(true)
	watermark.SetTextStyleItalic(false)

	// Change font of watermark text.
	watermark.SetFontFamily("Arial")
	watermark.SetFontSize(20)

	// Change color of watermark text.
	watermark.SetColor("red")

	// Change opacity of watermark text.
	watermark.SetOpacity(0.3)

	// Change layout to diagonal, by default watermark is horizontal.
	watermark.EnableDiagonalLayout(true)

	doc.SaveToFile("watermark-text.docx")
}
