// Copyright 2025 FoxyUtils ehf. All rights reserved.

package main

import (
	"log"
	"os"
	"time"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"
)

func init() {
	if err := license.SetMeteredKey(os.Getenv("UNIDOC_LICENSE_API_KEY")); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func main() {
	doc := document.New()
	defer doc.Close()

	doc.Settings.X().TrackRevisions = wml.NewCT_OnOff()

	p := doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr := p.AddRun()
	pr.AddText("This is some paragraph text.")

	currentTime := time.Now()
	author := "Author 1"
	revisionNumber := "002773B1"

	// Add insertion to current paragraph.
	p.AddInsertedText("Some inserted text is here.", revisionNumber, author, currentTime, 0)

	p = doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr = p.AddRun()
	pr.AddText("This is more paragraph text.")

	currentTime = time.Now()
	author = "Author 2"
	revisionNumber = ""

	// Add deletion to current paragraph.
	p.AddDeletedText("Some deleted text is there.", revisionNumber, author, currentTime, 1)

	doc.SaveToFile("insertions_deletions.docx")
}
