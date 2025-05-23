// Copyright 2025 FoxyUtils ehf. All rights reserved.
// This example shows how to generate chart and insert it to document as a picture.
package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	unipdflicense "github.com/unidoc/unipdf/v4/common/license"

	"github.com/disintegration/imaging"
	"github.com/unidoc/unichart"
	"github.com/unidoc/unichart/dataset"

	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"

	"github.com/unidoc/unipdf/v4/creator"
	"github.com/unidoc/unipdf/v4/model"
	"github.com/unidoc/unipdf/v4/render"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := unipdflicense.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
		panic(err)
	}

	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err = license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	doc := document.New()
	defer doc.Close()

	chart := &unichart.PieChart{
		Values: []dataset.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}
	chart.SetHeight(500)

	// Create unipdf chart component.
	c := creator.New()
	chartComponent := creator.NewChart(chart)

	// Draw chart component.
	if err := c.Draw(chartComponent); err != nil {
		log.Fatalf("failed to draw chart: %v", err)
	}

	// Save output file.
	if err := c.WriteToFile("output.pdf"); err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	// Render PDF file to image.
	renderPDFToImage("output.pdf")

	// If you have a ready chart image from another source, you can start here.
	img, err := common.ImageFromFile("preview.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	img1ref, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	para := doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(img1ref)
	if err != nil {
		log.Fatalf("unable to add anchored image: %s", err)
	}

	anchored.SetName("Chart")
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	doc.SaveToFile("output.docx")
}

func renderPDFToImage(filename string) {
	// Create reader.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open pdf file: %v", err)
	}
	defer file.Close()

	reader, err := model.NewPdfReader(file)
	if err != nil {
		log.Fatalf("Could not create reader: %v", err)
	}

	// Render pages.
	device := render.NewImageDevice()
	device.OutputWidth = 2000

	// Get page.
	page, err := reader.GetPage(1)
	if err != nil {
		log.Fatalf("Could not retrieve page: %v", err)
	}

	// Render page to PNG file.
	err = device.RenderToPath(page, "preview.png")
	if err != nil {
		log.Fatalf("Image rendering error: %v", err)
	}

	cropImage("preview.png")
}

func cropImage(imagePath string) {
	// Open the input image
	img, err := imaging.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Find the bounding box of the non-empty (foreground) region
	boundingBox := findBoundingBox(img)

	// Crop the image using the bounding box
	croppedImg := imaging.Crop(img, boundingBox)

	// Save the cropped image
	outFile, err := os.Create(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, croppedImg)
	if err != nil {
		log.Fatal(err)
	}
}

func findBoundingBox(img image.Image) image.Rectangle {
	bounds := img.Bounds()
	minX, maxX := bounds.Dx(), 0
	minY, maxY := bounds.Dy(), 0
	boundPadding := 100

	// Iterate over each pixel to find the bounding box
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			r, g, b, a := pixelColor.RGBA()

			if r != 65535 || g != 65535 || b != 65535 || a != 65535 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Define the bounding box based on the min and max values
	boundingBox := image.Rect(minX-boundPadding, minY-boundPadding, maxX+boundPadding+1, maxY+boundPadding+1)
	return boundingBox
}
