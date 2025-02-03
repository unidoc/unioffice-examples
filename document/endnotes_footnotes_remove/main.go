// Copyright 2023 FoxyUtils ehf. All rights reserved.

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
	doc, err := document.Open("footnotes_endnotes.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if ok, fnID := r.IsFootnote(); ok {
				if fnID == 2 {
					p.RemoveFootnote(2)

					break
				}
			}

			if ok, enID := r.IsEndnote(); ok {
				if enID == 1 {
					p.RemoveEndnote(1)

					break
				}
			}
		}
	}

	doc.SaveToFile("removed_footnote_endnote.docx")
}
