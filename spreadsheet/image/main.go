// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"math"
	"os"

	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/spreadsheet"
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
	ss := spreadsheet.New()
	defer ss.Close()
	// add a single sheet
	sheet := ss.AddSheet()

	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	iref, err := ss.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to workbook: %s", err)
	}

	dwng := ss.AddDrawing()
	sheet.SetDrawing(dwng)
	for i := float64(0); i < 360; i += 30 {
		anc := dwng.AddImage(iref, spreadsheet.AnchorTypeAbsolute)

		ang := i * math.Pi / 180
		x := 2 + 2*math.Cos(ang)
		y := 2 + +2*math.Sin(ang)
		anc.SetColOffset(measurement.Distance(x) * measurement.Inch)
		anc.SetRowOffset(measurement.Distance(y) * measurement.Inch)

		// set the image to 1x1 inches
		var w measurement.Distance = 1 * measurement.Inch
		anc.SetWidth(w)
		anc.SetHeight(iref.RelativeHeight(w))
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("image.xlsx")
}
