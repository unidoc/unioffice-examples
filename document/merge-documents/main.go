// Copyright 2017 FoxyUtils ehf. All rights reserved.
// This example demonstrates merging two documents into one.

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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
	doc0, err := document.Open("document0.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc0.Close()
	doc1, err := document.Open("document1.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc1.Close()
	doc0.AddParagraph().AddRun().AddPageBreak()

	err = doc0.Append(doc1)
	if err != nil {
		log.Fatalf("error appending document: %s", err)
	}
	doc0.SaveToFile("merged.docx")
}
