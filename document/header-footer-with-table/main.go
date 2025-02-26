/*
 * Copyright 2025 FoxyUtils ehf. All rights reserved.
 *
 * This example demonstrates how to add a table to the header and footer of a document.
 */

package main

import (
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/color"
	"github.com/unidoc/unioffice/v2/common"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"
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
	doc := document.New()
	defer doc.Close()

	// Add header.
	hdr := doc.AddHeader()

	// Add table into header.
	{
		// Create a table.
		table := hdr.AddTable()
		table.Properties().SetWidthPercent(50)
		table.Properties().SetAlignment(wml.ST_JcTableCenter)

		// Set borders for the table.
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		// Add a row to the table.
		row := table.AddRow()
		cell := row.AddCell()
		cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)

		para := cell.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.AddText("hello")

		cell = row.AddCell()
		cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)

		para = cell.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run = para.AddRun()
		run.AddText("world")
	}

	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	// Add footer
	ftr := doc.AddFooter()

	// Add table into footer
	{
		// Create a table.
		table := ftr.AddTable()
		table.Properties().SetWidthPercent(100)

		// Set borders for the table.
		row := table.AddRow()

		// Add a logo to the first cell.
		logo, err := common.ImageFromFile("logo.png")
		if err != nil {
			log.Fatalf("unable to create image: %s", err)
		}

		logoRef, err := ftr.AddImage(logo)
		if err != nil {
			log.Fatalf("unable to add image to document: %s", err)
		}

		// Add a logo to the first cell.
		cell := row.AddCell()
		para := cell.AddParagraph()
		run := para.AddRun()

		drw, err := run.AddDrawingInline(logoRef)
		if err != nil {
			log.Fatalf("unable to add inline image: %s", err)
		}
		drw.SetSize(67*measurement.Point, 20*measurement.Point)

		// Add text to the second cell.
		cell = row.AddCell()
		para = cell.AddParagraph()
		para.SetAlignment(wml.ST_JcRight)
		para.Properties().AddTabStop(6*measurement.Inch, wml.ST_TabJcRight, wml.ST_TabTlcNone)
		run = para.AddRun()
		run.AddText("Pg ")
		run.AddField(document.FieldCurrentPage)
		run.AddText(" of ")
		run.AddField(document.FieldNumberOfPages)
	}

	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)

	// Add text body
	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("Header and footer with table")

	err := doc.SaveToFile("header-footer-with-table.docx")
	if err != nil {
		panic(err)
	}
}
