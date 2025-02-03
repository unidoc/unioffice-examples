// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"

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
	hdrRow := sheet.AddRow()
	hdrRow.AddCell().SetString("Product Name")
	hdrRow.AddCell().SetString("Quantity")
	hdrRow.AddCell().SetString("Price")
	sheet.SetAutoFilter("A1:C6")

	// rows
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		row.AddCell().SetString(fmt.Sprintf("Product %d", r+1))
		row.AddCell().SetNumber(float64(r + 2))
		row.AddCell().SetNumber(float64(3*r + 1))

	}

	// sort column C, starting a row 2 to skip the header row
	sheet.Sort("C", 2, spreadsheet.SortOrderDescending)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("sort-filter.xlsx")
}
