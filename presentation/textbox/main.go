// Copyright 2020 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/measurement"
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
	ppt, err := presentation.Open("source.pptx")
	if err != nil {
		panic(err)
	}
	defer ppt.Close()
	slide := ppt.Slides()[0] // taking the first slide

	// Getting the list of text boxes
	tbs := slide.GetTextBoxes() // getting all textboxes

	for _, tb := range tbs {
		for _, p := range tb.X().TxBody.P {
			for _, tr := range p.EG_TextRun {
				fmt.Println(tr.TextRunChoice.R.T)
			}
		}
	}

	// Editing the existing text box
	tb := tbs[0]                                            // taking first of them
	run := tb.X().TxBody.P[0].EG_TextRun[0].TextRunChoice.R // taking the first run of the first paragraph
	run.T = "Edited TextBox text"                           // changing the text of the run

	// creating a new text box
	newTb := slide.AddTextBox()
	newTb.SetOffsetX(measurement.Inch * 5)
	newTb.SetOffsetY(measurement.Inch * 4)

	newPara := newTb.AddParagraph()
	newRun := newPara.AddRun()
	newRun.SetText("New TextBox text")

	ppt.SaveToFile("mod.pptx")
}
