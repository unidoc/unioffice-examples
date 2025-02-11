// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// This example shows how Nodes can be used to work
// generically with document contents to find and copy
// contents across files.  In this example we
// 1. Load an input sample document sample.docx
// 2. Identify paragraphs with style "heading 1" and use that as a section divider
// 3. Create a new document for each section, and output each section to separate file with names: node-document-i.docx where i is the section index.
// In addition we illustrate how to perform some simple replacements that are included in the output files.

package main

import (
	"fmt"
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
	doc, err := document.Open("sample.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	// Get document element as nodes.
	nodes := doc.Nodes()

	// Replace text inside nodes.
	nodes.ReplaceText("Where can I get some?", "The Title is Replaced")
	nodes.ReplaceText("Why do we use it?", "I Am The New Title of Document")

	// Find nodes by style name.
	nodesByStyle := nodes.FindNodeByStyleName("heading 1")

	// Create new document.
	for i, nodeParent := range nodesByStyle {
		// Create new document.
		newDoc := document.New()
		defer newDoc.Close()
		fmt.Println("New document will be created")
		fmt.Println("Heading:", nodeParent.Text())

		nextNodeIndex := i + 1
		minIndex := -1
		for ni, node := range nodes.X() {
			if nodeParent.X() == node.X() {
				minIndex = ni
			}

			// If there's next node, break the loop and go to next parentNode.
			if len(nodesByStyle) > nextNodeIndex {
				if nodesByStyle[nextNodeIndex].X() == node.X() {
					minIndex = ni
					break
				}
			}

			// Insert node to new document.
			if ni >= minIndex && minIndex > -1 {
				newDoc.AppendNode(node)
			}
		}

		// Save new doucment.
		err := newDoc.SaveToFile(fmt.Sprintf("output/node-document-%d.docx", i))
		if err != nil {
			log.Fatalf("error while saving file: %v\n", err.Error())
		}
	}
}
