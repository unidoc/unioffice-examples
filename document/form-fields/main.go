// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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
	doc := document.New()
	defer doc.Close()

	p0 := doc.AddParagraph()

	checkbox := p0.AddCheckBox("checkbox1")
	checkbox.SetSize(20)
	checkbox.SetChecked(true)
	checkbox.SetEnabled(true)
	checkbox.SetCalcOnExit(false)

	p1 := doc.AddParagraph()

	textInput := p1.AddTextInput("textInput1")
	textInput.SetValue("Hello World")
	textInput.SetEnabled(false)

	p2 := doc.AddParagraph()

	ddList := p2.AddDropdownList("ddList1")
	ddList.SetPossibleValues([]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"})
	ddList.SetValue("Earth")
	ddList.SetName("Solar system")

	// FindAllFields is a helper function that traverses the document
	// identifying fields
	fields := doc.FormFields()
	fmt.Println("found", len(fields), "fields")

	for _, fld := range fields {
		fmt.Println("- Name:", fld.Name(), "Type:", fld.Type(), "Value:", fld.Value())
	}

	doc.SaveToFile("filled-form.docx")
}
