// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unioffice/v2/measurement"
	"github.com/unidoc/unioffice/v2/schema/soo/wml"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

// NOTE: This example can only run on Windows and requires that Word be installed.

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	defer doc.Close()

	// Force the TOC to update upon opening the document
	doc.Settings.SetUpdateFieldsOnOpen(true)

	// Add a TOC
	doc.AddParagraph().AddRun().AddField(document.FieldTOC)
	// followed by a page break
	doc.AddParagraph().Properties().AddSection(wml.ST_SectionMarkNextPage)

	nd := doc.Numbering.AddDefinition()
	for i := 0; i < 9; i++ {
		lvl := nd.AddLevel()
		lvl.SetFormat(wml.ST_NumberFormatNone)
		lvl.SetAlignment(wml.ST_JcLeft)
		if i%2 == 0 {
			lvl.SetFormat(wml.ST_NumberFormatBullet)
			lvl.RunProperties().SetFontFamily("Symbol")
			lvl.SetText("")
		}
		lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(i) * measurement.Inch)
	}

	// and finally paragraphs at different heading levels
	for i := 0; i < 4; i++ {
		para := doc.AddParagraph()
		para.SetNumberingDefinition(nd)
		para.Properties().SetHeadingLevel(1)
		para.AddRun().AddText("First Level")

		doc.AddParagraph().AddRun().AddText(lorem)
		for i := 0; i < 3; i++ {
			para := doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.Properties().SetHeadingLevel(2)
			para.AddRun().AddText("Second Level")
			doc.AddParagraph().AddRun().AddText(lorem)

			para = doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.Properties().SetHeadingLevel(3)
			para.AddRun().AddText("Third Level")
			doc.AddParagraph().AddRun().AddText(lorem)
		}
	}
	doc.SaveToFile("toc.docx")

	cwd, _ := os.Getwd()
	err := UpdateFields(filepath.Join(cwd, "toc.docx"))
	if err != nil {
		log.Printf("error creating Word object: %s", err)
	}
}

// UpdateFields uses go-ole to convert a docx to a PDF using the Word application
func UpdateFields(source string) error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	iunk, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		return err
	}
	defer iunk.Release()

	word := iunk.MustQueryInterface(ole.IID_IDispatch)
	defer word.Release()

	docs := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer docs.Release()
	wordDoc := oleutil.MustCallMethod(docs, "Open", source).ToIDispatch()
	defer wordDoc.Release()

	const wdFormatXMLDocument = 12
	oleutil.MustCallMethod(wordDoc, "SaveAs2", source, wdFormatXMLDocument)
	oleutil.MustCallMethod(wordDoc, "Close")
	oleutil.MustCallMethod(word, "Quit")
	return nil
}
