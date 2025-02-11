// Copyright 2017 FoxyUtils ehf. All rights reserved.
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

	// Set page size to paper A4 and orientation to landscape.
	// Paper A4 size is 8.3" × 11.7".
	// You can set the orientation with wml.ST_PageOrientationLandscape or wml.ST_PageOrientationPortrait,
	section := doc.BodySection()
	section.SetPageSizeAndOrientation(measurement.Inch*8.3, measurement.Inch*11.7, wml.ST_PageOrientationLandscape)

	doc.SaveToFile("page-size-and-orientation.docx")
}
