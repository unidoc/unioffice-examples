/*
 * Copyright 2025 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates how to add HTML tags to format a text inside the document.
 */
package main

import (
	"fmt"
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
	doc := document.New()
	defer doc.Close()

	// <b> tag is used to make the text bold.
	createParaRunHTML(doc, "<b>HTML bold text</b>")

	// <strong> tag is also used to make the text bold.
	createParaRunHTML(doc, "<strong>HTML bold text</stong>")

	// <i> tag is used to make the text italic.
	createParaRunHTML(doc, "<i>HTML italic text</i>")

	// <em> tag is also used to make the text italic.
	createParaRunHTML(doc, "<em>HTML italic text</em>")

	// <u> tag is used to underline the text.
	createParaRunHTML(doc, "<u>HTML underline text</u>")

	// <strike> tag is used to strike through the text.
	createParaRunHTML(doc, "<strike>HTML strike text</strike>")

	// <sub> tag is used to make the text subscript.
	createParaRunHTML(doc, "<sub>HTML subscript text</sub>")

	// <sup> tag is used to make the text superscript.
	createParaRunHTML(doc, "<sup>HTML superscript text</sup>")

	// <mark> tag is used to highlight the text using default color of yellow.
	createParaRunHTML(doc, `<mark>HTML highlighted text with default color</mark>`)

	// We could provide a color attribute to the <mark> tag to change the color of the highlighted text.
	createParaRunHTML(doc, `<mark color="green">HTML green highlighted text</mark>`)

	// Combining multiple tags to format the text.
	createParaRunHTML(doc, "<b><i>HTML bold and italic text</i></b>")

	doc.AddParagraph()
	doc.AddParagraph()

	// Tags can be nested to format the text.
	htmlPara := doc.AddParagraph()
	htmlPara.AddHTML("<b><i>HTML text</i></b> and this one is <u><strike>underlined and stroked</strike></u>")

	htmlPara = doc.AddParagraph()
	htmlPara.AddHTML("<b>This is a <em>sample</em> of <i>HTML text</i></b> and this one is <u><strike>stroked</strike> and underlined text</u>")

	htmlPara = doc.AddParagraph()
	htmlPara.AddHTML("<u>Another <b>sample</b> of <i>HTML text</i> here</u>")

	htmlPara = doc.AddParagraph()
	htmlPara.AddHTML("<u>Another <b>sample of <i>HTML text here</i></b></u>")

	htmlPara = doc.AddParagraph()
	htmlPara.AddHTML("<u><b><i>Another sample</i> of HTML text</b> here</u>")

	htmlPara = doc.AddParagraph()
	htmlPara.AddHTML("<u><b>Malformed HTML text might not be correctly <i>formatted</i></u>")

	htmlPara = doc.AddParagraph()
	r := htmlPara.AddRun()
	r.AddText("I want to show html tag <i>italic</i> and the result is ")
	r.AddHTML("<i>italic</i>")

	err := doc.SaveToFile("text_formatting_using_tag.docx")
	if err != nil {
		fmt.Printf("error saving document: %s", err)
	}
}

func createParaRunHTML(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddHTML(s)
	return run
}
