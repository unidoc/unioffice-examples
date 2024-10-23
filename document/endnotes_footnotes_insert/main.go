// Copyright 2023 FoxyUtils ehf. All rights reserved.

package main

import (
	"os"

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

func main() {
	doc := document.New()
	defer doc.Close()

	p := doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr := p.AddRun()
	pr.AddText("This is some paragraph text.")

	p = doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr = p.AddRun()
	pr.AddText("This is more paragraph text. This paragraph has a footnote.")

	addFootnote(p, []string{
		"This is a footnote.",
		"It has multiple paragraphs",
		"Indeed this is the last.",
	})

	pr = p.AddRun()
	pr.AddText(" It also has text after the footnote.")

	p = doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr = p.AddRun()
	pr.AddText("This is yet more paragraph text. This paragraph has another footnote.")

	addFootnote(p, []string{"This footnote we will modify by changing the text."})

	p = doc.AddParagraph()
	p.SetAfterSpacing(measurement.Point * 12)

	pr = p.AddRun()
	pr.AddText("This is the final paragraph. It has an endnote.")

	addEndnote(p, []string{"This is an end note."})

	pr = p.AddRun()
	pr.AddText(" Because we want to make sure those aren't different.")

	addEndnote(p, []string{"Second end note"})

	doc.SaveToFile("add_footnote_endnote.docx")
}

func addFootnote(p document.Paragraph, notes []string) {
	if len(notes) == 0 {
		return
	}

	fn := p.AddFootnote(notes[0])
	for i, n := range notes {
		if i == 0 {
			continue
		}

		fnp := fn.AddParagraph()
		fnpr := fnp.AddRun()
		fnpr.AddText(n)
	}
}

func addEndnote(p document.Paragraph, notes []string) {
	if len(notes) == 0 {
		return
	}

	en := p.AddEndnote(notes[0])
	for i, n := range notes {
		if i == 0 {
			continue
		}

		enp := en.AddParagraph()
		enpr := enp.AddRun()
		enpr.AddText(n)
	}
}
