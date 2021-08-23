// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

// This example demonstrates flattening all formulas from an input Excel file and outputs the flattened values to a new xlsx.

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
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
