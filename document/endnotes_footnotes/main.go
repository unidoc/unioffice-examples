// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
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
	doc, err := document.Open("footnotes_endnotes.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	if doc.HasFootnotes() {
		fmt.Printf("Document has %02d footnotes.\n", len(doc.Footnotes()))
	} else {
		fmt.Println("Document has no footnotes")
	}

	if doc.HasEndnotes() {
		fmt.Printf("Document has %02d endnotes.\n", len(doc.Endnotes()))
	} else {
		fmt.Println("Document has no endnotes")
	}
}
