// Copyright 2022 UniDoc ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/presentation"
	"github.com/unidoc/unioffice/schema/soo/dml"
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
	ppt := presentation.New()
	defer ppt.Close()

	for i := 0; i < 5; i++ {
		slide := ppt.AddSlide()

		tb := slide.AddTextBox()
		tb.Properties().SetGeometry(dml.ST_ShapeTypeStar10)

		tb.Properties().SetWidth(3 * measurement.Inch)
		pos := measurement.Distance(i) * measurement.Inch
		tb.Properties().SetPosition(pos, pos)

		tb.Properties().SetSolidFill(color.AliceBlue)
		tb.Properties().LineProperties().SetSolidFill(color.Blue)

		p := tb.AddParagraph()
		p.Properties().SetAlign(dml.ST_TextAlignTypeCtr)

		r := p.AddRun()
		r.SetText("gooxml")
		r.Properties().SetSize(24 * measurement.Point)
	}

	// SlideSize returns the presentation slide size,
	// using SlideSize, we can sets the slide size.
	// Example below sets the slide size width: 13.33" and height: 7.50" (wide 16x9).
	slideSize := ppt.SlideSize()
	slideSize.SetSize(presentation.SlideScreenSize16x9)

	if err := ppt.Validate(); err != nil {
		log.Fatal(err)
	}
	ppt.SaveToFile("slide-size.pptx")
}
