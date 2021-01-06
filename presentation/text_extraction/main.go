package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/presentation"
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
			if runProps.SolidFill != nil && runProps.SolidFill.SchemeClr != nil {
				fmt.Println("SolidFill:", runProps.SolidFill.SchemeClr.ValAttr)
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
