// Copyright 2020 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/presentation"
	"github.com/unidoc/unioffice/schema/soo/dml"
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
	ppt := presentation.New()
	defer ppt.Close()
	slide := ppt.AddSlide()

	tbl := slide.AddTable()
	for ci := 0; ci < 4; ci++ {
		col := tbl.AddCol()
		col.SetWidth(measurement.Millimeter * 52)
	}
	for ri := 0; ri < 3; ri++ {
		row := tbl.AddRow()
		row.SetHeight(measurement.Inch)
		for ci, cell := range row.Cells() {
			cell.TxBody = dml.NewCT_TextBody()

			para := dml.NewCT_TextParagraph()
			cell.TxBody.P = append(cell.TxBody.P, para)

			egtr := dml.NewEG_TextRun()
			para.EG_TextRun = append(para.EG_TextRun, egtr)
			egtr.R = dml.NewCT_RegularTextRun()
			egtr.R.T = fmt.Sprintf("Cell %d:%d", ri, ci)
		}
	}

	style := dml.NewCT_TableStyle()
	style.WholeTbl = dml.NewCT_TablePartStyle()
	tcStyle := dml.NewCT_TableStyleCellStyle()
	tcStyle.Fill = dml.NewCT_FillProperties()
	tcStyle.Fill.SolidFill = dml.NewCT_SolidColorFillProperties()
	tcStyle.Fill.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	tcStyle.Fill.SolidFill.SrgbClr.ValAttr = "FF9900"
	style.WholeTbl.TcStyle = tcStyle
	tbl.SetStyle(style)

	tbl.SetOffsetX(measurement.Inch)
	tbl.SetOffsetY(measurement.Millimeter * 20)
	ppt.SaveToFile("out.pptx")
}
