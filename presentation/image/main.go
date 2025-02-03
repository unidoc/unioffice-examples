// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

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

func main() {
	ppt := presentation.New()
	defer ppt.Close()
	imgColor, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}
	imgBW, err := common.ImageFromFile("gopher.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	irefColor, err := ppt.AddImage(imgColor)
	if err != nil {
		log.Fatal(err)
	}

	irefBW, err := ppt.AddImage(imgBW)
	if err != nil {
		log.Fatal(err)
	}

	slide := ppt.AddSlide()

	ibColor := slide.AddImage(irefColor)
	ibColor.Properties().SetWidth(2 * measurement.Inch)
	ibColor.Properties().SetHeight(irefColor.RelativeHeight(2 * measurement.Inch))

	ibBW := slide.AddImage(irefBW)
	ibBW.Properties().SetWidth(2 * measurement.Inch)
	ibBW.Properties().SetHeight(irefBW.RelativeHeight(2 * measurement.Inch))
	ibBW.Properties().SetPosition(4*measurement.Inch, 4*measurement.Inch)

	if err := ppt.Validate(); err != nil {
		log.Fatal(err)
	}
	ppt.SaveToFile("image.pptx")
}
