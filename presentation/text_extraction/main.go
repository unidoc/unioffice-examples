// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/presentation"
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
	// Start building pptx
	ppt, err := presentation.Open("extract.pptx")
	if err != nil {
		fmt.Println("presentation.OpenTemplate err ", err)
		os.Exit(1)
	}
	defer ppt.Close()

	pe := ppt.ExtractText()
	fmt.Println(pe.Text()) // Output as plain text

	for _, slide := range pe.Slides {
		for i, item := range slide.Items {
			fmt.Println(i)
			fmt.Println(item.Text)
			runProps := item.Run.RPr
			fmt.Println("Bold:", runProps.BAttr != nil)
			fmt.Println("Italic:", runProps.IAttr != nil)
			if runProps.SzAttr != nil {
				fmt.Println("Font size:", *runProps.SzAttr/100)
			}
			if runProps.FillPropertiesChoice.SolidFill != nil && runProps.FillPropertiesChoice.SolidFill.SchemeClr != nil {
				fmt.Println("SolidFill:", runProps.FillPropertiesChoice.SolidFill.SchemeClr.ValAttr)
			}
			if tblInfo := item.TableInfo; tblInfo != nil {
				fmt.Println("Row:", tblInfo.RowIndex)
				fmt.Println("Column:", tblInfo.ColIndex)
				if row := item.TableInfo.Row; row != nil {
					fmt.Println("height:", row.HAttr)
				}
				grid := tblInfo.Table.TblGrid
				fmt.Println("width:", grid.GridCol[tblInfo.ColIndex].WAttr)
			}
			fmt.Println("--------")
		}
	}
}
