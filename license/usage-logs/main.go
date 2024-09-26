/*
 * This example showcases PDF generation from docx document with UniOffice package.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/unidoc/unipdf/v3/common"
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

	// Set the log level to info or higher
	common.SetLogger(common.NewConsoleLogger(common.LogLevelInfo))
	// Enable the verbose mode logging
	license.SetMeteredKeyUsageLogVerboseMode(true)
	license.SetMeteredKeyPersistentCache(false)

}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run unipdf_license_usage_log.go input_dir output_dir\n")
		os.Exit(1)
	}

	inputDir := os.Args[1]

	files, err := filepath.Glob(inputDir + "*.docx")
	if err != nil {
		panic(err)
	}

	for _, filePath := range files {
		base := path.Base(filePath)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		outputPath := fmt.Sprintf("output/%s.pdf", name)
		doc, err := document.Open(filePath + ".docx")
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
