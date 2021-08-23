// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	defer doc.Close()

	para := doc.AddParagraph()
	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)
	}

	// Put image watermark to document.
	img1, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	// Add image to document.
	img1ref, err := doc.AddImage(img1)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	// Set watermark to doc.
	watermark := doc.AddWatermarkPicture(img1ref)
	watermark.SetPictureWashout(true)

	// Get image size and set watermark size.
	// SetPictureSize accept parameter image width and image height.
	imageSize := img1ref.Size()
	watermark.SetPictureSize(int64(imageSize.X*measurement.Point), int64(imageSize.Y*measurement.Point))

	doc.SaveToFile("watermark-picture.docx")
}
