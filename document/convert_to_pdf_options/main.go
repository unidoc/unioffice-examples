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

func main() {
	doc, err := document.Open("merge_fields.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	// Set convert options for ProcessFields to true.
	co := &convert.Options{
		ProcessFields: true,
	}

	// Convert to PDF and process the fields in document.
	c := convert.ConvertToPdfWithOptions(doc, co)

	err = c.WriteToFile("output/merge_fields.pdf")
	if err != nil {
		log.Fatalf("error converting document: %s", err)
	}
}
