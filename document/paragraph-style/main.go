// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
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

	// Create paragraph and apply Title style to paragraph.
	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("What is Lorem Ipsum?")

	// Create paragraph and apply Heading1 style to paragraph.
	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Lorem Ipsum is simply dummy text of the printing and typesetting industry.")

	// Create custom style named CustomStyle1 for paragraph style.
	// Set last boolean parameter to TRUE if you want to make it default for paragraph.
	style := doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("Custom Style 1")
	// Set spacing of paragraph before and after.
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(measurement.Inch * 2)
	// Set line spacing to single, in this case, 12 points is text height.
	customStyle.ParagraphProperties().SetLineSpacing(12*measurement.Point, wml.ST_LineSpacingRuleAuto)
	// Apply style `CustomStyle1` to paragraph.
	para = doc.AddParagraph()
	para.SetStyle("CustomStyle1")
	run = para.AddRun()
	run.AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	doc.SaveToFile("paragraph-style.docx")
}
