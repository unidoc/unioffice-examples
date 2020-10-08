package main

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/schema/soo/pml"

	"github.com/unidoc/unioffice/presentation"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
eyJsaWNlbnNlX2lkIjoiNTRlNTE4MmItNGFlMS00YTQ2LTQxZTUtZDRmNGM4ZDEwM2U1IiwiY3VzdG9tZXJfaWQiOiJhYTkyZjI3Mi1kYTM4LTRiY2MtNmM1ZC1kZjAwZmI5ODQ0ODEiLCJjdXN0b21lcl9uYW1lIjoiZGFuaXNoIiwiY3VzdG9tZXJfZW1haWwiOiJkYW5pc2hAdW5pZG9jLmlvIiwidGllciI6ImJ1c2luZXNzIiwiY3JlYXRlZF9hdCI6MTYwMTY0MzI2MCwiZXhwaXJlc19hdCI6MTYzMzIxOTE5OSwiY3JlYXRvcl9uYW1lIjoiVW5pRG9jIFN1cHBvcnQiLCJjcmVhdG9yX2VtYWlsIjoic3VwcG9ydEB1bmlkb2MuaW8iLCJ1bmlwZGYiOnRydWUsInVuaW9mZmljZSI6dHJ1ZSwidHJpYWwiOmZhbHNlfQ==
+
Cyipjn8hCDzPZBp2DOyWpez1QGk4SnS5PL+Voy58NJze1ywj+LxgSrKi6u2M2tTh0mbTvM4hAPq+/6M0mxDWS54bDwOWY0kbmJiSv7SoILaJqrhJod+fuVfCyskTF+C5BV+y4z/X7EvadTwEFpR8yRLibsAmFWjGTXvs3Tk+mgHFonU/u+PIaUy7Lmon5dPuREHXJmZiiOUXeEan/iQv+WWckk1VNVAY0U0y/gO/eeDAptyagpsCRNDQlfwlp8dMZGUCECwSofg7I01bfwGKxOcM0HOdu288Tzxs6cOgB7+jqoUGaMRfuq19n5ytvBgpfx4dTI+f70PEw7uVGFwqmw==
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `danish`)
	if err != nil {
		panic(err)
	}
}

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
		print("error")
	}
	saleData := res.SaleData

	for _, data := range saleData { // Iterate through the sale data
		// remove any existing slides
		for _, s := range ppt.Slides() {
			ppt.RemoveSlide(s)
		}

		l, err := ppt.GetLayoutByName("Title and Caption")
		if err != nil {
			log.Fatalf("error retrieving layout: %s", err)
		}

		sld, err := ppt.AddDefaultSlideWithLayout(l)
		if err != nil {
			log.Fatalf("error adding slide: %s", err)
		}

		ph, _ := sld.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
		ph.SetText(fmt.Sprintf("Sale Data For: %s", data.Area))
		ph, _ = sld.GetPlaceholder(pml.ST_PlaceholderTypeBody)
		ph.SetText("Created with github.com/unidoc/unioffice/")

		tac, _ := ppt.GetLayoutByName("Title and Content")

		sld, err = ppt.AddDefaultSlideWithLayout(tac)
		if err != nil {
			log.Fatalf("error adding slide: %s", err)
		}

		ph, _ = sld.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
		ph.SetText(fmt.Sprintf("Data for %s, Managed by %s", data.Area, data.Manager))
		ph, _ = sld.GetPlaceholderByIndex(1)
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

		if err != nil {
			log.Fatalf("error opening template: %s", err)
		}
		ppt.SaveToFile(fmt.Sprintf("%s.pptx", data.Area))
	}
}
