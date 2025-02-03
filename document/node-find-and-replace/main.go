// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// This example shows how Nodes can be used to work
// generically with document contents to find and copy
// contents across files.  In this example we
// 1. Load an input sample document sample.docx
// 2. Identify paragraphs with text "Cell 1" and regexp "What is.*"
// 3. Create a new document output named: node-find-and-replace.docx.

package main

import (
	"log"
	"os"
	"regexp"

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
	doc, err := document.Open("sample.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	// Get document element as nodes.
	nodes := doc.Nodes()

	// Find node element by text.
	nodeList := nodes.FindNodeByText("Cell 1")
	for _, node := range nodeList {
		log.Println("Find by text found: ", node.Text())
	}

	// Replace text.
	nodes.ReplaceText("Cell 1", "Cell Replacement")

	// Find node element by regexp.
	nodeList = nodes.FindNodeByRegexp(regexp.MustCompile(`What is.*`))
	for _, node := range nodeList {
		log.Println("Find by regex found: ", node.Text())
	}

	// Replace text using regexp.
	nodes.ReplaceTextByRegexp(regexp.MustCompile(`What is.*`), "Why I am replaced?")

	// Save new doucment.
	err = doc.SaveToFile("output/node-find-and-replace.docx")
	if err != nil {
		log.Fatalf("error while saving file: %v\n", err.Error())
	}
}
