// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/spreadsheet"

	"github.com/unidoc/unioffice/v2/schema/soo/sml"
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
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("Hello World!")
	sheet.Cell("B1").SetString("will not be visible") // as it's not the first cell within a merged range Excel warns you when you do this through the UI
	sheet.AddMergedCells("A1", "C2")

	centered := ss.StyleSheet.AddCellStyle()
	centered.SetHorizontalAlignment(sml.ST_HorizontalAlignmentCenter)
	centered.SetVerticalAlignment(sml.ST_VerticalAlignmentCenter)
	sheet.Cell("A1").SetStyle(centered)

	for _, m := range sheet.MergedCells() {
		fmt.Println("merged region", m.Reference(), "has contents", m.Cell().GetString())
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("merged.xlsx")
}
