// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
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

	sheet.Cell("A1").SetString("Hello World!")
	sheet.Comments().AddCommentWithStyle("A1", "Gopher", "This looks interesting.")
	sheet.Comments().AddCommentWithStyle("C10", "Gopher", "This is a different comment.")

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("comments.xlsx")
}
