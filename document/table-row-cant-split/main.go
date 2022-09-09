// Copyright 2022 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
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

var lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

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

	doc.AddParagraph().AddRun().AddText(lorem)
	doc.AddParagraph().AddRun().AddText(lorem)

	// Create Table
	{
		table := doc.AddTable()

		// Width of the page.
		table.Properties().SetWidthPercent(100)

		// With thick borders.
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

		for i := 0; i < 25; i++ {
			row := table.AddRow()

			// Get row properties and set row can't split.
			rowProps := row.Properties()
			rowProps.SetCantSplit(true)

			para := row.AddCell().AddParagraph()

			// Get paragraph properties and set keep next.
			paraProps := para.Properties()
			paraProps.SetKeepWithNext(true)

			run := para.AddRun()
			run.AddText(fmt.Sprintf("Row No. %d", i+1))
		}
	}

	doc.AddParagraph().AddRun().AddText(lorem)

	if err := doc.SaveToFile("table-row-cant-split.docx"); err != nil {
		panic(err)
	}
}
