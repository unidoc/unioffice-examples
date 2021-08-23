// Copyright 2020 FoxyUtils ehf. All rights reserved.

// This example shows how to create a document and add headers
// by separating them into odd pages and even pages

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
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
	// Create a document
	doc := document.New()
	defer doc.Close()

	// First add some content
	for i := 0; i < 100; i++ {
		doc.AddParagraph().AddRun().AddText(lorem)
	}

	// Construct even header
	evenHdr := doc.AddHeader()
	evenHdr.AddParagraph().AddRun().AddText("Even Header")
	doc.BodySection().SetHeader(evenHdr, wml.ST_HdrFtrEven)

	// Construct odd header
	oddHdr := doc.AddHeader()
	oddHdr.AddParagraph().AddRun().AddText("Odd Header")
	doc.BodySection().SetHeader(oddHdr, wml.ST_HdrFtrDefault)

	// Set the EvenAndOddHeaders flag
	boolTrue := true
	doc.Settings.X().EvenAndOddHeaders = &wml.CT_OnOff{
		ValAttr: &sharedTypes.ST_OnOff{Bool: &boolTrue},
	}

	// Save the file
	if err := doc.SaveToFile("even-odd-header.docx"); err != nil {
		fmt.Println(err)
	}
}
