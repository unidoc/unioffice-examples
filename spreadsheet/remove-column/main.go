// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main
// This example demonstrates flattening all formulas from an input Excel file and outputs the flattened values to a new xlsx.

import (
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
	ss, err := spreadsheet.Open("original.xlsx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer ss.Close()

	sheet0, err := ss.GetSheet("Cells")
	if err != nil {
		log.Fatalf("error opening sheet: %s", err)
	}

	err = sheet0.RemoveColumn("C")
	if err != nil {
		log.Fatalf("error removing column: %s", err)
	}

	sheet1, err := ss.GetSheet("MergedCells")
	if err != nil {
		log.Fatalf("error opening sheet: %s", err)
	}

	err = sheet1.RemoveColumn("C")
	if err != nil {
		log.Fatalf("error removing column: %s", err)
	}

	ss.SaveToFile("removed.xlsx")
}
