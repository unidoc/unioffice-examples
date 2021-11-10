// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// This example shows how Nodes can be used to work
// generically with document contents to find and copy
// contents across files.  In this example we
// 1. Load an input sample document sample.docx
// 2. Identify element from sample.docx for table and paragraphs then use that as a section divider
// 3. Create a new document for each section
// 4. Save output for each section to separate file with names: node-selection-paragraph.docx and node-selection-table.docx

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
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

	// Get document element as nodes.
	nodes := doc.Nodes()

	// Find nodes by condition.
	nodeList := nodes.FindNodeByCondition(func(node *document.Node) bool {
		switch node.X().(type) {
		case *document.Table:
			return true
		}
		return false
	}, false)

	// Create new document of node tables.
	newDocTable := document.New()
	defer newDocTable.Close()

	for _, node := range nodeList {
		newDocTable.AppendNode(node)
	}

	// Save new doucment.
	err = newDocTable.SaveToFile("output/node-selection-table.docx")
	if err != nil {
		log.Fatalf("error while saving file: %v\n", err.Error())
	}

	// Find nodes by condition.
	nodeList = nodes.FindNodeByCondition(func(node *document.Node) bool {
		switch node.X().(type) {
		case *document.Paragraph:
			return true
		}
		return false
	}, false)

	// Create new document of node.
	newDocParagraph := document.New()
	defer newDocParagraph.Close()

	// Append paragraph to document.
	for i := 0; i < 5; i++ {
		newDocParagraph.AppendNode(nodeList[i])
	}

	// Save new doucment.
	err = newDocParagraph.SaveToFile("output/node-selection-paragraph.docx")
	if err != nil {
		log.Fatalf("error while saving file: %v\n", err.Error())
	}
}
