// Copyright 2020 FoxyUtils ehf. All rights reserved.

package main

import (
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
	doc, err := document.Open("document.docx")
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	doc.SetStrict(false) // document will be saved as Word document (this is a default option for new files)
	doc.SaveToFile("conformance_transitional.docx")
	doc.SetStrict(true) // document will be saved in the Strict mode
	doc.SaveToFile("conformance_strict.docx")
}
