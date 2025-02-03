// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"math/rand"
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
	sheet := ss.AddSheet()

	row := sheet.AddRow()
	row.AddCell()
	for i := 0; i < 99; i++ {
		row.AddCell().SetString("Header")
	}
	for i := 0; i < 100; i++ {
		row = sheet.AddRow()
		row.AddCell().SetString("Header")
		for j := 0; j < 99; j++ {
			row.AddCell().SetNumber(rand.Float64() * 100)
		}
	}

	// freeze the first row and column
	sheet.SetFrozen(true, true)

	/* this is equivalent to
	v := sheet.InitialView()
	v.SetState(sml.ST_PaneStateFrozen)
	v.SetYSplit(1)
	v.SetXSplit(1)
	v.SetTopLeft("B2")
	*/

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("freeze-rows-cols.xlsx")
}
