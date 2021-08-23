// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
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
	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum.`

	// single spaced
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(lorem)
	run.AddText(lorem)
	run.AddBreak()

	// double spaced is twice the text height (24 points in this case as the text height is 12 points)
	para = doc.AddParagraph()
	para.Properties().Spacing().SetLineSpacing(24*measurement.Point, wml.ST_LineSpacingRuleAuto)
	run = para.AddRun()
	run.AddText(lorem)
	run.AddText(lorem)
	run.AddBreak()

	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	doc.SaveToFile("line-spacing.docx")
}
