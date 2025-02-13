// Copyright 2024 FoxyUtils ehf. All rights reserved.
// Example removes all tables from a document.

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
)

func init() {
	err := license.SetMeteredKey(os.Getenv("UNIDOC_LICENSE_API_KEY"))
	if err != nil {
		panic(err)
	}
}

func main() {
	doc, err := document.Open("sample.docx")
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

	err = doc.SaveToFile("output.docx")
	if err != nil {
		panic(err)
	}
}
