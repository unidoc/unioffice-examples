// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/spreadsheet"
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
	ss := spreadsheet.New()
	defer ss.Close()
	// add a single sheet
	sheet := ss.AddSheet()

	// rows
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		// and cells
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			// cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
			rt := cell.SetRichTextString()
			run := rt.AddRun()
			run.SetText(fmt.Sprintf("row %d ", r))
			run.SetBold(true)
			run.SetColor(color.Red)

			run = rt.AddRun()
			run.SetSize(16)
			run.SetItalic(true)
			run.SetFont("Courier")
			run.SetText(fmt.Sprintf("cell %d", c))

		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("rich-text.xlsx")
}
