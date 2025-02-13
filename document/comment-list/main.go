/*
 * Copyright 2025 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates how to retrieve and display comments stored within a DOCX file.
 */

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
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
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

	// Check if the document has comments.
	if !doc.HasComments() {
		fmt.Println("Document has no comment")
		return
	}

	comments := doc.Comments()

	fmt.Printf("Document has %d comments:\n", len(comments))

	// Iterate through the comments structure and display the comment text.
	for _, c := range comments {
		cmt := c.X()
		fmt.Printf("%d. Comment by %s: ", cmt.IdAttr, cmt.AuthorAttr)

		for _, ble := range cmt.EG_BlockLevelElts {
			for _, cbc := range ble.BlockLevelEltsChoice.EG_ContentBlockContent {
				for _, p := range cbc.ContentBlockContentChoice.P {
					for _, pc := range p.EG_PContent {
						for _, crc := range pc.PContentChoice.EG_ContentRunContent {
							for _, ric := range crc.ContentRunContentChoice.R.EG_RunInnerContent {
								if ric.RunInnerContentChoice.T != nil {
									fmt.Println(ric.RunInnerContentChoice.T.Content)
								}
							}
						}
					}
				}
			}
		}
	}
}
