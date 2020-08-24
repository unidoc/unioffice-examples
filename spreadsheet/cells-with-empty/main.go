// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

// This example demonstrates outputing all cells in a row of an excel spreadsheet, including empty cells.

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
	ss, err := spreadsheet.Open("test.xlsx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer ss.Close()

	s := ss.Sheets()[0]

	maxColumnIdx := s.MaxColumnIdx()
	for _, row := range s.Rows() {
		for _, cell := range row.CellsWithEmpty(maxColumnIdx) {
			fmt.Println(cell.Reference(), ":", cell.GetFormattedValue())
		}
	}
	fmt.Print("\n\n\n")

	s.Cell("F4").SetString("Hello world")
	maxColumnIdx = s.MaxColumnIdx()
	for _, row := range s.Rows() {
		for _, cell := range row.CellsWithEmpty(maxColumnIdx) {
			fmt.Println(cell.Reference(), ":", cell.GetFormattedValue())
		}
	}
}
