package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/ofc/sharedTypes"
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

var text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."

func main() {
	doc := document.New()
	defer doc.Close()

	// Adding content
	for i := 0; i < 25; i++ {
		doc.AddParagraph().AddRun().AddText(text)
	}

	// Getting the image
	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	// Checking if there is a header
	hdr, ok := doc.BodySection().GetHeader(wml.ST_HdrFtrDefault)
	if !ok {
		// If not, creating it
		hdr = doc.AddHeader()
		doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	}

	// Main header with an image
	hdrPara := hdr.AddParagraph()
	hdrRun := hdrPara.AddRun()
	hdrRun.AddText("Main Document Title")
	hdrRun.AddBreak()

	iref, err := hdr.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to header: %s", err)
	}
	imgInl, _ := hdrPara.AddRun().AddDrawingInline(iref)
	imgInl.SetSize(0.5*measurement.Inch, 0.5*measurement.Inch)

	// Even header
	evenHdr := doc.AddHeader()
	evenHdr.AddParagraph().AddRun().AddText("Even Header")
	doc.BodySection().SetHeader(evenHdr, wml.ST_HdrFtrEven)

	// Odd header
	oddHdr := doc.AddHeader()
	oddHdr.AddParagraph().AddRun().AddText("Odd Header")
	doc.BodySection().SetHeader(oddHdr, wml.ST_HdrFtrDefault)

	// Set EvenAndOddHeaders flag
	boolTrue := true
	doc.Settings.X().EvenAndOddHeaders = &wml.CT_OnOff{
		ValAttr: &sharedTypes.ST_OnOff{Bool: &boolTrue},
	}

	// Add Footer
	ftr := doc.AddFooter()
	ftrPara := ftr.AddParagraph()
	ftrPara.Properties().AddTabStop(6*measurement.Inch, wml.ST_TabJcRight, wml.ST_TabTlcNone)
	ftrRun := ftrPara.AddRun()
	ftrRun.AddText("Some subtitle goes here")
	ftrRun.AddTab()
	ftrRun.AddText("Pg ")
	ftrRun.AddField(document.FieldCurrentPage)
	ftrRun.AddText(" of ")
	ftrRun.AddField(document.FieldNumberOfPages)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrEven)

	// Save the file
	if err := doc.SaveToFile("combined-header-footer.docx"); err != nil {
		log.Fatalf("Failed to write file %s", err)
	}
}
