// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

// This example demonstrates outputing all cells in a row of an excel spreadsheet, including empty cells.

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
