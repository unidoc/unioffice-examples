// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/schema/soo/sml"
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

	hdrStyle := ss.StyleSheet.AddCellStyle()
	hdrStyle.SetHorizontalAlignment(sml.ST_HorizontalAlignmentCenter)

	lightGray := ss.StyleSheet.Fills().AddFill()
	lightGrayPattern := lightGray.SetPatternFill()
	lightGrayPattern.SetFgColor(color.LightGray)
	hdrStyle.SetFill(lightGray)

	hdr := sheet.AddRow()
	hdrCell := hdr.AddCell()
	hdrCell.SetString("Products")
	hdrCell.SetStyle(hdrStyle)

	hdrCell = hdr.AddCell()
	hdrCell.SetString("# Sold")
	hdrCell.SetStyle(hdrStyle)

	for i := 0; i < 10; i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetString(fmt.Sprintf("Product %d", i+1))
		cell = row.AddCell()
		cell.SetNumber(float64(i + 1))

	}

	totalRow := sheet.AddRow()
	totalCell := totalRow.AddCell()
	totalCell = totalRow.AddCell()
	totalCell.SetFormulaRaw("SUM(B2:B11)")

	ss.RecalculateFormulas()
	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating: %s", err)
	}
	ss.SaveToFile("formula.xlsx")
}
