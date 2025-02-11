// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/schema/soo/dml"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/presentation"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

const lorem = "Lorem ipsum dolor sit amet."

func main() {
	ppt := presentation.New()
	defer ppt.Close()
	imgColor, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	irefColor, err := ppt.AddImage(imgColor)
	if err != nil {
		log.Fatal(err)
	}

	slide := ppt.AddSlide()

	img := slide.AddImage(irefColor)
	img.Properties().SetWidth(2 * measurement.Inch)
	img.Properties().SetHeight(irefColor.RelativeHeight(2 * measurement.Inch))

	tb := slide.AddTextBox()
	tb.SetTextAnchor(dml.ST_TextAnchoringTypeCtr) // vertical center
	para := tb.AddParagraph()
	para.Properties().SetAlign(dml.ST_TextAlignTypeCtr) // horizontal center
	run := para.AddRun()

	run.Properties().SetBold(true)
	run.Properties().SetSolidFill(color.Red)
	run.SetText("Look a Gopher!")

	tb.Properties().SetGeometry(dml.ST_ShapeTypeChevron)
	tb.Properties().SetFlipHorizontal(true)
	tb.Properties().SetSolidFill(color.LightBlue)
	tb.Properties().LineProperties().SetWidth(0.125 * measurement.Inch)
	tb.Properties().LineProperties().SetSolidFill(color.DarkBlue)
	tb.Properties().SetPosition(2.5*measurement.Inch, 0.5*measurement.Inch)

	tb = slide.AddTextBox()
	tb.Properties().SetPosition(3.5*measurement.Inch, 2.5*measurement.Inch)
	for i := 0; i < 4; i++ {
		para = tb.AddParagraph()
		para.Properties().SetBulletFont("Wingdings")
		para.Properties().SetBulletChar("Ø")
		para.Properties().SetLevel(int32(i))
		run = para.AddRun()
		if i%2 == 1 {
			run.Properties().SetSolidFill(color.DarkRed)
		}
		run.SetText("Foo")
	}

	if err := ppt.Validate(); err != nil {
		log.Fatal(err)
	}
	ppt.SaveToFile("complex.pptx")
}
