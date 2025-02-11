// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
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
	doc, err := document.Open("form.docx")
	if err != nil {
		log.Fatalf("error opening form: %s", err)
	}
	defer doc.Close()

	// FindAllFields is a helper function that traverses the document
	// identifying fields
	fields := doc.FormFields()
	fmt.Println("found", len(fields), "fields")

	for _, fld := range fields {
		fmt.Println("- Name:", fld.Name(), "Type:", fld.Type(), "Value:", fld.Value())

		switch fld.Type() {
		case document.FormFieldTypeText:
			// you can directly set values on text fields
			fld.SetValue("testing 123")
		case document.FormFieldTypeCheckBox:
			// you can check check boxes
			fld.SetChecked(true)
		case document.FormFieldTypeDropDown:
			// and select items in a dropdown, here the value must be one of the
			// fields possible values
			lpv := len(fld.PossibleValues())
			if lpv > 0 {
				fld.SetValue(fld.PossibleValues()[lpv-1])
			}
		}
	}

	doc.SaveToFile("filled-form.docx")
}
