// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package main

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

func main() {
	wb, err := spreadsheet.Open("extract_styles.xlsx")
	if err != nil {
		panic(err)
	}
	defer wb.Close()

	extracted := wb.ExtractText()
	flattened := extracted.Text()

	fmt.Println(flattened)

	sheetCells := extracted.Sheets[0].Cells

	styleSheet := wb.StyleSheet
	cellStyles := styleSheet.CellStyles()

	for ri := 0; ri < 4; ri++ {
		for ci := 0; ci < 4; ci++ {
			i := ri * 4 + ci
			fmt.Printf("\nRow: %d, Column: %d\n", ri, ci)
			cell := sheetCells[i]
			cellX := cell.Cell.X()
			fmt.Printf("Text: %s\n", cell.Text)
			if cellX.SAttr == nil {
				panic("expected style to be non-nil")
			}
			style := cellStyles[*cellX.SAttr]
			font := style.GetFont()
			if font.B != nil {
				fmt.Println("Bold")
			}
			if font.I != nil {
				fmt.Println("Italic")
			}
			if len(font.Color) == 0 {
				panic("expected font to have a color")
			}
			fontColor := font.Color[0]
			if fontColor == nil {
				panic("expected font color to be non-nil")
			}
			fmt.Println("Font color theme:", *fontColor.ThemeAttr)
			fmt.Println("Font color tint:", *fontColor.TintAttr)
			fill := style.GetFill()
			patternFill := fill.PatternFill
			if patternFill == nil {
				panic("expected pattern fill to be non-nil")
			}
			cellColor := patternFill.FgColor
			if cellColor == nil {
				panic("expected foreground color to be non-nil")
			}
			fmt.Println("Cell color theme:", *cellColor.ThemeAttr)
			fmt.Println("Cell color tint:", *cellColor.TintAttr)
		}
	}
}
