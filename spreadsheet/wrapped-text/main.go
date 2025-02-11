// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
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

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	ss := spreadsheet.New()
	defer ss.Close()
	// add a single sheet
	sheet := ss.AddSheet()

	row := sheet.AddRow()
	cell := row.AddCell()

	wrapped := ss.StyleSheet.AddCellStyle()
	wrapped.SetWrapped(true)
	cell.SetString(lorem)
	cell.SetStyle(wrapped)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("wrapped.xlsx")
}
