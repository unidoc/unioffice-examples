// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// This example shows how Nodes can be used to work
// generically with document contents to find and copy
// contents across files.  In this example we
// 1. Load an input sample document sample1.docx and sample2.docx
// 2. Identify table element at sample1.docx
// 3. Identify paragraph element at sample2.docx
// 4. Create a new document that contains table from sample1.docx and some of paragraph from sample2.docx
// 5.	Save the output to file with names: node-combine.docx.

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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
	doc1, err := document.Open("sample1.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc1.Close()

	doc2, err := document.Open("sample2.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc2.Close()

	// Get document element as nodes.
	nodes1 := doc1.Nodes()

	nodes2 := doc2.Nodes()

	// Find nodes by condition,
	// FindNodeByCondition take 2 argument, function and wholeElements: true or false.
	nodeList1 := nodes1.FindNodeByCondition(func(node *document.Node) bool {
		switch node.X().(type) {
		case *document.Table:
			return true
		}
		return false
	}, false)

	// Create new document of node tables.
	newDoc := document.New()
	defer newDoc.Close()

	for _, node := range nodeList1 {
		newDoc.AppendNode(node)
	}

	// Find nodes by condition.
	// FindNodeByCondition take 2 argument, function and wholeElements: true or false.
	nodeList2 := nodes2.FindNodeByCondition(func(node *document.Node) bool {
		switch node.X().(type) {
		case *document.Paragraph:
			return true
		}
		return false
	}, false)

	for _, node := range nodeList2 {
		newDoc.AppendNode(node)
	}

	// Save new doucment.
	err = newDoc.SaveToFile("output/node-combine.docx")
	if err != nil {
		log.Fatalf("error while saving file: %v\n", err.Error())
	}
}
