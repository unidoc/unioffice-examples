// Copyright UniDoc ehf. All rights reserved.
// This example show how's to set the paragraph run properties for Right to Left text.
package main

import (
	"os"

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

	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.SetAfterSpacing(measurement.Point * 12)
	run := para.AddRun()
	run.AddText("Try to write RTL Paragraph Text")

	para = doc.AddParagraph()
	para.SetStyle("Heading 2")
	run = para.AddRun()
	run.AddText("Hebrew Text:")

	para = doc.AddParagraph()
	para.SetAlignment(wml.ST_JcEnd)
	para.SetAfterSpacing(measurement.Point * 12)

	run = para.AddRun()
	run.AddText("ב אחד כלליים חופשית, עוד אירועים ותשובות האנציקלופדיה על. אל אתה מושגי הבהרה ויקימדיה. ב עזה חשמל בלשנות, מה החלה.")
	run.Properties().SetRightToLeft(true)

	para = doc.AddParagraph()
	para.SetStyle("Heading 2")
	run = para.AddRun()
	run.AddText("Arabic Text:")

	para = doc.AddParagraph()
	para.SetAlignment(wml.ST_JcEnd)
	para.SetAfterSpacing(measurement.Point * 12)

	run = para.AddRun()
	run.AddText("أبسط معادلة هي \"1 + 1 = 2\". إنه أساس الرياضيات الحديثة.")
	run.Properties().SetRightToLeft(true)

	err := doc.SaveToFile("paragraph-rtl.docx")
	if err != nil {
		panic(err)
	}
}
