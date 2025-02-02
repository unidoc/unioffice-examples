// Copyright 2020 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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

// setHeader sets the header by creating new or using existing header
func setOrCreateHeader(doc *document.Document, text string) {
	// Check if header with the given type exists already
	hdr, ok := doc.BodySection().GetHeader(wml.ST_HdrFtrDefault)
	if !ok {
		hdr = doc.AddHeader()
		doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	}

	// Add Text to header
	para := hdr.AddParagraph()
	run := para.AddRun()
	run.AddBreak()
	run.AddText(text)
}

func main() {
	doc := document.New()

	// This will create a new header with text "Header 1"
	setOrCreateHeader(doc, "Header 1")

	// This will add a new text to the existing header
	setOrCreateHeader(doc, "Header 2")

	if err := doc.SaveToFile("doc-existing-header.docx"); err != nil {
		fmt.Println(err)
		return
	}
}
