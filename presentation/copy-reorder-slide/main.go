// Copyright 2024 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/presentation"
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
	// Open presentation from which we will copy slide
	pptFrom, err := presentation.Open("source.pptx")
	if err != nil {
		fmt.Println("presentation.Open error ", err)
		os.Exit(1)
	}
	defer pptFrom.Close()

	// Open presentation from which we will copy slide
	pptTo, err := presentation.Open("extract.pptx")
	if err != nil {
		fmt.Println("presentation.Open error ", err)
		os.Exit(1)
	}
	defer pptTo.Close()

	// copy slide from one presentation from another, it becomes the last slide
	_, err = pptTo.CopySlide(pptFrom.Slides()[0])
	if err != nil {
		fmt.Println("error copying slide ", err)
		os.Exit(1)
	}

	// move inserted slide to the very beginning of presentation, it becomes slide 0
	err = pptTo.MoveSlide(4, 0)
	if err != nil {
		fmt.Println("error moving slide ", err)
		os.Exit(1)
	}

	pptTo.SaveToFile("output.pptx")
}
