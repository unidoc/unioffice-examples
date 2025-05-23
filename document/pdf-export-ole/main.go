// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"
	"os"
	"path/filepath"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	
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

// NOTE: This example can only run on Windows and requires that Word be installed.

func main() {
	doc := document.New()
	defer doc.Close()

	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("Simple Document Formatting")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	run = para.AddRun()
	run.AddText("Some Heading Text")
	doc.SaveToFile("simple.docx")

	cwd, _ := os.Getwd()
	err := ConvertToPDF(filepath.Join(cwd, "simple.docx"), filepath.Join(cwd, "simple.pdf"))
	if err != nil {
		log.Printf("error creating Word object: %s", err)
	}
}

// ConvertToPDF uses go-ole to convert a docx to a PDF using the Word application
func ConvertToPDF(source, destination string) error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	iunk, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		return err
	}

	word := iunk.MustQueryInterface(ole.IID_IDispatch)
	defer word.Release()

	// opening then saving works due to the call to doc.Settings.SetUpdateFieldsOnOpen(true) above

	docs := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	wordDoc := oleutil.MustCallMethod(docs, "Open", source).ToIDispatch()

	// file format constant comes from https://msdn.microsoft.com/en-us/vba/word-vba/articles/wdsaveformat-enumeration-word
	const wdFormatPDF = 17
	oleutil.MustCallMethod(wordDoc, "SaveAs2", destination, wdFormatPDF)
	oleutil.MustCallMethod(wordDoc, "Close")
	oleutil.MustCallMethod(word, "Quit")
	return nil
}
