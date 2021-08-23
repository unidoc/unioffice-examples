// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/reference"
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
	var width = measurement.Distance(100)

	// Set style of cell to be not protected and not hidden
	cellStyle1 := ss.StyleSheet.AddCellStyle()
	cellStyle1.SetProtection(false, false)
	// Apply cellStyle1 to range of cells A1:D10 with looping.
	from, to, err := reference.ParseRangeReference("A1:D10")
	if err != nil {
		panic(err)
	}
	for rowIdx := from.RowIdx; rowIdx <= to.RowIdx; rowIdx++ {
		for colIdx := from.ColumnIdx; colIdx <= to.ColumnIdx; colIdx++ {
			currentCell := reference.IndexToColumn(colIdx)
			sheet.Row(rowIdx).Cell(currentCell).SetStyle(cellStyle1)
		}
	}

	// Apply cellStyle1 to column F.
	sheet.Column(6).SetStyle(cellStyle1)
	sheet.Column(6).SetWidth(width)

	// Add protection to sheet and lock it.
	sp := sheet.Protection()
	sp.LockSheet(true)
	// Set password for unlocking sheet.
	sp.SetPassword("unioffice")

	if err := ss.SaveToFile("cell-protection.xlsx"); err != nil {
		fmt.Println(err)
	}
}
