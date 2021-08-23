// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/unidoc/unioffice/common"
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

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	defer doc.Close()

	img1, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}
	img2data, err := ioutil.ReadFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	img2, err := common.ImageFromBytes(img2data)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	img1ref, err := doc.AddImage(img1)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}
	img2ref, err := doc.AddImage(img2)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	para := doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(img1ref)
	if err != nil {
		log.Fatalf("unable to add anchored image: %s", err)
	}

	// Drop square text wrap image.
	anchored.SetName("Gopher")
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	run := para.AddRun()
	for i := 0; i < 40; i++ {
		run.AddText(lorem)

		// drop an inline image in
		if i == 13 {
			inl, err := run.AddDrawingInline(img1ref)
			if err != nil {
				log.Fatalf("unable to add inline image: %s", err)
			}
			inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}

		// Drop image behind text.
		if i == 18 {
			anchorDraw, err := run.AddDrawingAnchored(img2ref)
			if err != nil {
				log.Fatalf("unable to add behind text image: %s", err)
			}
			anchorDraw.SetTextWrapBehindText()
			anchorDraw.SetOrigin(wml.WdST_RelFromHColumn, wml.WdST_RelFromVParagraph)
			anchorDraw.SetHAlignment(wml.WdST_AlignHUnset)
			anchorDraw.SetOffset(2*measurement.Inch, 3*measurement.Inch)
			anchorDraw.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}

		// Drop image in front of text.
		if i == 21 {
			anchorDraw, err := run.AddDrawingAnchored(img2ref)
			if err != nil {
				log.Fatalf("unable to add in front of text image: %s", err)
			}
			anchorDraw.SetTextWrapInFrontOfText()
			anchorDraw.SetOrigin(wml.WdST_RelFromHColumn, wml.WdST_RelFromVParagraph)
			anchorDraw.SetHAlignment(wml.WdST_AlignHUnset)
			anchorDraw.SetOffset(5*measurement.Inch, 4*measurement.Inch)
			anchorDraw.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}

		// Drop image with bottom and top wrap text.
		if i == 23 {
			anchorDraw, err := run.AddDrawingAnchored(img2ref)
			if err != nil {
				log.Fatalf("unable to add top and bottom wrap text image: %s", err)
			}
			anchorDraw.SetTextWrapTopAndBottom()
			anchorDraw.SetOrigin(wml.WdST_RelFromHColumn, wml.WdST_RelFromVParagraph)
			anchorDraw.SetHAlignment(wml.WdST_AlignHCenter)
			anchorDraw.SetYOffset(7 * measurement.Inch)
			anchorDraw.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}

		// Drop image with through wrap text.
		if i == 27 {
			anchorDraw, err := run.AddDrawingAnchored(img2ref)
			if err != nil {
				log.Fatalf("unable to add through wrap text image: %s", err)
			}
			// Create anchor draw option, parameter for options is followImageShape true or false.
			anchorDrawOption := document.NewAnchorDrawWrapOptions()
			anchorDraw.SetTextWrapThrough(anchorDrawOption)
			anchorDraw.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVParagraph)
			anchorDraw.SetOffset(2*measurement.Inch, 3*measurement.Inch)
			anchorDraw.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}

		// Drop image with tight wrap text.
		if i == 30 {
			anchorDraw, err := run.AddDrawingAnchored(img2ref)
			if err != nil {
				log.Fatalf("unable to add tight wrap text image: %s", err)
			}
			// Create anchor draw option, parameter for options is followImageShape true or false.
			anchorDrawOption := document.NewAnchorDrawWrapOptions()
			anchorDraw.SetTextWrapTight(anchorDrawOption)
			anchorDraw.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVParagraph)
			anchorDraw.SetOffset(5*measurement.Inch, 2*measurement.Inch)
			anchorDraw.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}
	}
	doc.SaveToFile("image-text-wrap.docx")
}
