// Copyright 2017 FoxyUtils ehf. All rights reserved.
/*
 * This example showcases PDF generation from docx document with UniOffice package.
 */

package main

import (
	"fmt"
	"log"
	"os"

	unipdflicense "github.com/unidoc/unipdf/v4/common/license"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/presentation"
	"github.com/unidoc/unioffice/v2/presentation/convert"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := unipdflicense.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
		panic(err)
	}

	// This example requires both for unioffice and unipdf.
	err = license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
		panic(err)
	}
}

var filenames = []string{
	"chart",
	"lists",
	"picture",
	"several_slides",
	"table",
}

func main() {
	for _, filename := range filenames {
		ppt, err := presentation.Open(filename + ".pptx")
		if err != nil {
			log.Fatalf("error opening document: %s", err)
		}
		defer ppt.Close()
		outputFilename := "output/" + filename
		c := convert.ConvertToPdf(ppt)
		err = c.WriteToFile(fmt.Sprintf("%s.pdf", outputFilename))
		if err != nil {
			log.Fatalf("error saving PDF: %s", err)
		}
	}
}
