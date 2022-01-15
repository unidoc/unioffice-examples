/*
 * Paragraph borders
 * Create paragraph that contains border.
 * For border line break paragraph, as MS Word function automatic border line -
 * its similar with the paragraph that having border bottom value.
 *
 * The reference regarding this can be seen here https://www.avantixlearning.ca/microsoft-word/how-to-insert-and-remove-lines-in-microsoft-word
 */

// Copyright 2021 FoxyUtils ehf. All rights reserved.

package main

import (
	"os"

	"github.com/unidoc/unioffice/color"
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

func main() {
	doc := document.New()
	defer doc.Close()

	// Create paragraph and apply Title style to paragraph.
	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("What is Lorem Ipsum?")

	// Create paragraph and apply Heading1 style to paragraph.
	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Lorem Ipsum is simply dummy text of the printing and typesetting industry.")

	// Add border bottom to paragraph.
	border := para.Borders()
	border.SetBottom(wml.ST_BorderDotted, color.Auto, 1*measurement.Point)

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	// Add border bottom to paragraph.
	border = para.Borders()
	border.SetBottom(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Where can I get some?")

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.")

	// Add border All (left, right, top, bottom) to paragraph.
	border = para.Borders()
	border.SetAll(wml.ST_BorderDotDotDash, color.Auto, 1*measurement.Point)

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Why do we use it?")

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).")

	// Add border left, right, top, bottom to paragraph.
	border = para.Borders()
	border.SetLeft(wml.ST_BorderDotDotDash, color.Auto, 1*measurement.Point)
	border.SetRight(wml.ST_BorderDotted, color.Auto, 1*measurement.Point)
	border.SetTop(wml.ST_BorderThick, color.Auto, 1*measurement.Point)
	border.SetBottom(wml.ST_BorderTriple, color.Auto, 1*measurement.Point)

	err := doc.SaveToFile("paragraph-borders.docx")
	if err != nil {
		panic(err)
	}
}
