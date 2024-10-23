// Copyright 2024 FoxyUtils ehf. All rights reserved.
// Example removes all tables from a document.

package main

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
)

func init() {
	err := license.SetMeteredKey(os.Getenv("UNIDOC_LICENSE_API_KEY"))
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage %s go run main.go input.docx output.docx", os.Args[0])
		return
	}

	inputPath := args[1]
	outputPath := args[2]

	doc, err := document.Open(inputPath)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	nodes := doc.Nodes()
	for _, node := range nodes.X() {
		switch node.X().(type) {
		case *document.Table:
			node.Remove()
			// Uncomment block below to remove paragraphs
			// case *document.Paragraph:
			//	node.Remove()
		}
	}

	err = doc.SaveToFile(outputPath)
	if err != nil {
		panic(err)
	}
}
