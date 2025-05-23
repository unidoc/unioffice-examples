// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

// This example demonstrates converting a workbook to a PDF file.

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/spreadsheet"
	"github.com/unidoc/unioffice/v2/spreadsheet/convert"
	pdflicense "github.com/unidoc/unipdf/v4/common/license"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
	err = pdflicense.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

var filenames = []string{
	"simple",
}

func main() {
	for _, filename := range filenames {
		wb, err := spreadsheet.Open(filename + ".xlsx")
		if err != nil {
			log.Fatalf("error opening spreadsheet: %s", err)
		}
		defer wb.Close()
		if _, err = os.Stat(filename); os.IsNotExist(err) {
			err = os.Mkdir(filename, 0755)
			if err != nil {
				log.Fatalf("error creating folder: %s", err)
			}
		}
		for si, s := range wb.Sheets() {
			c := convert.ConvertToPdf(&s)
			err = c.WriteToFile(fmt.Sprintf("%s/sheet_%d.pdf", filename, si))
			if err != nil {
				log.Fatalf("error saving PDF: %s", err)
			}
		}
	}
}
