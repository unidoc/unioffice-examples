// Copyright 2025 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
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

	for i := 0; i < 7; i++ {
		img, err := common.ImageFromFile("gopher.png")
		if err != nil {
			log.Fatalf("unable to create image: %s", err)
		}

		imgref, err := doc.AddImage(img)
		if err != nil {
			log.Fatalf("unable to create image reference: %s", err)
		}
		anchored, err := doc.AddParagraph().AddRun().AddDrawingAnchored(imgref)
		if err != nil {
			log.Fatalf("unable to create anchored drawing: %s", err)
		}

		// Showcase for different supported image effects.

		switch i {
		case 0:
			// Add soft edge with radius equal to 10 points.
			anchored.SetSoftEdgeImageEffect(measurement.Point * 10)
		case 1:
			// Add glow effect with 8 points radius and blue color.
			anchored.SetGlowImageEffect(measurement.Point*8, color.Blue)
		case 2:
			// Add inner shadow effect with 15 points radius, 10 points offset, red color and 150 degrees rotation
			anchored.SetInnerShadowImageEffect(measurement.Point*15, measurement.Point*10, color.Red, 150)
		case 3:
			// Add outer shadow effect with 27 points radius, 10 points offset, and light blue color
			anchored.SetOuterShadowImageEffect(measurement.Point*27, measurement.Point*10, color.LightBlue, 0)
		case 4:
			// Add reflection effect with 1 point radius, 50% initial transparency and 90% size
			anchored.SetReflectionImageEffect(measurement.Point, 50, 90)
		case 5:
			// Add bevel effect
			anchored.SetBevelImageEffect()
		case 6:
			// Add 3D rotation effect
			anchored.Set3DRotationImageEffect()
		}
	}

	doc.SaveToFile("image_effects.docx")
}
