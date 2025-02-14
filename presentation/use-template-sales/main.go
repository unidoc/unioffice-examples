// Copyright 2025 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"

	"encoding/json"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/schema/soo/pml"

	"github.com/unidoc/unioffice/v2/presentation"
)

type SaleInfo struct {
	Area      string
	Sale      string
	Customers int
	Manager   string
}

type Data struct {
	Year     int
	ID       string
	SaleData []SaleInfo
}

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	jsonData := []byte(`
		{
			"year":2020,
			"id":"JI23SA",
			"SaleData":[
				{
					"area": "Michigan",
					"sale": "10 million",
					"customers": 10000,
					"manager": "Henry"
				},
				{
					"area": "Cincinnati",
					"sale": "2 million",
					"customers": 490,
					"manager": "John Green"
				},
				{
					"area": "Washington",
					"sale": "150 million",
					"customers": 100000,
					"manager": "Smith Johnson"
				}
			]
		}
	`)

	ppt, err := presentation.OpenTemplate("template.pptx")
	if err != nil {
		log.Fatalf("unable to open template: %s", err)
	}
	defer ppt.Close()

	for i, layout := range ppt.SlideLayouts() {
		fmt.Println(i, " LL ", layout.Name(), "/", layout.Type())
	}

	var res Data
	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}
	saleData := res.SaleData

	for _, data := range saleData { // Iterate through the sale data
		// Remove any existing slides
		for _, s := range ppt.Slides() {
			err := ppt.RemoveSlide(s)
			if err != nil {
				log.Fatalf("error removing slide: %s", err)
			}
		}

		l, err := ppt.GetLayoutByName("Title and Caption")
		if err != nil {
			log.Fatalf("error retrieving layout: %s", err)
		}

		sld, err := ppt.AddDefaultSlideWithLayout(l)
		if err != nil {
			log.Fatalf("error adding slide: %s", err)
		}

		ph, err := sld.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
		if err != nil {
			log.Fatalf("error getting placeholder type title: %s", err)
		}

		ph.SetText(fmt.Sprintf("Sale Data For: %s", data.Area))

		ph, err = sld.GetPlaceholder(pml.ST_PlaceholderTypeBody)
		if err != nil {
			log.Fatalf("error getting placeholder type body: %s", err)
		}

		ph.SetText("Created with github.com/unidoc/unioffice/")

		tac, err := ppt.GetLayoutByName("Title and Content")
		if err != nil {
			log.Fatalf("error retrieving layout: %s", err)
		}

		sld, err = ppt.AddDefaultSlideWithLayout(tac)
		if err != nil {
			log.Fatalf("error adding slide: %s", err)
		}

		ph, err = sld.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
		if err != nil {
			log.Fatalf("error getting placeholder type title: %s", err)
		}

		ph.SetText(fmt.Sprintf("Data for %s, Managed by %s", data.Area, data.Manager))

		ph, err = sld.GetPlaceholderByIndex(1)
		if err != nil {
			log.Fatalf("error getting placeholder by index: %s", err)
		}

		ph.ClearAll()

		para := ph.AddParagraph()
		run := para.AddRun()
		run.SetText(fmt.Sprintf("Here is the number of sales in %s: $%s", data.Area, data.Sale))

		para = ph.AddParagraph()
		run = para.AddRun()
		run.SetText(fmt.Sprintf("Number of Customers: %d", data.Customers))

		para = ph.AddParagraph()
		run = para.AddRun()
		run.SetText(fmt.Sprintf("Manager: %s", data.Manager))

		if err := ppt.SaveToFile(fmt.Sprintf("%s.pptx", data.Area)); err != nil {
			log.Fatalf("error saving presentation: %s", err)
		}
	}
}
