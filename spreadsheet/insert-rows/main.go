// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `Company Name`)
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
			cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
		}
	}

	// no insert some rows after row 2
	for i := 0; i < 4; i++ {
		sheet.InsertRow(2).AddCell().SetString(fmt.Sprintf("inserted at 2, iter %d", i))
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("insert-rows.xlsx")
}
