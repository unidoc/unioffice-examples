// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()
	fmt.Println("Total number of pages in the opened document:", doc.AppProperties.Pages())
	newDoc := document.New()
	fmt.Println("Total number of pages in the new document:", newDoc.AppProperties.Pages())
}
