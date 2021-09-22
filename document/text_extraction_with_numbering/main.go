// Copyright 2020 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
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
	doc, err := document.Open("numbered_list.docx")
	if err != nil {
		panic(err)
	}
	// To extract the text and work with the formatted info in a simple fashion, you can use:
	extracted := doc.ExtractText()
	for ei, e := range extracted.Items {
		fmt.Println(ei)
		fmt.Println("Text:", e.Text)
		if e.Run != nil && e.Run.RPr != nil {
			runProps := e.Run.RPr
			fmt.Println("Bold:", runProps.B != nil)
			fmt.Println("Italic:", runProps.I != nil)
			if color := runProps.Color; color != nil {
				fmt.Printf("Color: #%s\n", runProps.Color.ValAttr)
			}
			if highlight := runProps.Highlight; highlight != nil {
				fmt.Printf("Highlight: %s\n", runProps.Highlight.ValAttr.String())
			}
		}
		if tblInfo := e.TableInfo; tblInfo != nil {
			if tc := tblInfo.Cell; tc != nil {
				fmt.Println("Row:", tblInfo.RowIndex)
				fmt.Println("Column:", tblInfo.ColIndex)
				if pr := tc.TcPr; pr != nil {
					if pr.Shd != nil {
						fmt.Printf("Shade color: #%s\n", pr.Shd.FillAttr)
					}
				}
			}
		}
		if drawingInfo := e.DrawingInfo; drawingInfo != nil {
			fmt.Println("Height in mm:", measurement.FromEMU(drawingInfo.Height)/measurement.Millimeter)
			fmt.Println("Width in mm:", measurement.FromEMU(drawingInfo.Width)/measurement.Millimeter)
		}
		fmt.Println("--------")
	}
	// document.ExtractTextOptions contains options to include the numbering and it's spacing indent.
	// If you want to work with the flattened text, simply use:
	fmt.Println("\nFLATTENED:")
	fmt.Println(extracted.TextWithOptions(
		document.ExtractTextOptions{
			WithNumbering:   true,
			NumberingIndent: " ",
		}))
}
