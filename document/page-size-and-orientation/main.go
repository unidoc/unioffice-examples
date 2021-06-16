// Copyright 2017 FoxyUtils ehf. All rights reserved.
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
