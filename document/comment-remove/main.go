/*
 * This example demonstrates how to remove a comment from a DOCX file.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
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
	if len(os.Args) < 3 {
		fmt.Printf("Syntax: go run main.go input.docx comment_id\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	commentId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("error parsing comment id: %s", err)
	}

	doc, err := document.Open(inputPath)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	listComments(doc)

	if sucess := doc.RemoveComment(int64(commentId)); !sucess {
		fmt.Println("Failed removing comment")

		return
	}

	fmt.Println("\nComment removed successfully.")
	fmt.Println("")

	listComments(doc)
}

func listComments(doc *document.Document) {
	comments := doc.Comments()
	fmt.Printf("Document has %d comments.\n", len(comments))

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
