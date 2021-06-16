/*
 * This example showcases PDF generation from docx document with UniOffice package.
 */

package main

import (
	"fmt"
	"log"
	"os"

	unipdflicense "github.com/unidoc/unipdf/v3/common/license"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/document/convert"
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
	"headers_footers",
	"image_square",
	"table",
	"text_only_portrait",
	"text_only_landscape",
	"textbox_anchor",
	"textbox_inline",
}

func main() {
	for _, filename := range filenames {
		outputPath := filename + ".pdf"
		doc, err := document.Open(filename + ".docx")
		if err != nil {
			log.Fatalf("error opening document: %s", err)
		}
		defer doc.Close()
		c := convert.ConvertToPdf(doc)

		err = c.WriteToFile(outputPath)
		if err != nil {
			log.Fatalf("error converting document: %s", err)
		}
	}
}
