// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	// err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	// if err != nil {
	// 	panic(err)
	// }
}

// Letter represents a given message with letter content and receiver name.
type Letter struct {
	Receiver   string   `json:"receiver"`
	Paragraphs []string `json:"paragraphs"`
}

func main() {

	// We can now print out all styles in the document, verifying that they
	// exist.
	// for _, s := range doc.Styles.Styles() {
	// 	fmt.Println("style", s.Name(), "has ID of", s.StyleID(), "type is", s.Type())
	// }

	letters, err := loadMessage("./data/letters.json")
	if err != nil {
		panic(err)
	}

	for _, letter := range letters {
		generateDoc("Sample_Company_Letter.docx", letter, fmt.Sprintf("letter_to_%s.docx", letter.Receiver))
	}

}

func loadMessage(dataPath string) ([]Letter, error) {
	// Open the JSON file
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// Read the contents of the file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// Create a slice of Message objects to hold the data
	var letters []Letter

	// Deserialize (unmarshal) the JSON data into the slice
	err = json.Unmarshal(byteValue, &letters)
	if err != nil {
		return nil, err
	}

	return letters, nil

}

func generateDoc(templatePath string, letter Letter, outputName string) error {
	// When Word saves a document, it removes all unused styles.  This means to
	// copy the styles from an existing document, you must first create a
	// document that contains text in each style of interest.  As an example,
	// see the template.docx in this directory.  It contains a paragraph set in
	// each style that Word supports by default.
	templateDoc, err := document.OpenTemplate(templatePath)
	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}
	defer templateDoc.Close()

	// set header of the document
	if len(templateDoc.Headers()) > 0 {
		h := templateDoc.Headers()[0]
		templateDoc.BodySection().SetHeader(h, wml.ST_HdrFtrDefault)
		para := h.AddParagraph()
		run := para.AddRun()
		run.AddBreak()
	}

	// And create documents setting their style to the style ID (not style name).

	t := time.Now()
	dateTime := t.Format("January 2, 2006")

	// take the formatting from the template for date time text
	para := templateDoc.AddParagraph()
	para.SetStyle("Normal") // style name taken from the doc.Styles.Styles()
	para.AddRun().AddText(dateTime)
	run := para.AddRun()
	run.AddBreak()

	intro := fmt.Sprintf("Dear %s,", letter.Receiver)
	para = templateDoc.AddParagraph()
	para.SetStyle("Normal") // style name taken from the doc.Styles.Styles()
	para.AddRun().AddText(intro)
	run = para.AddRun()
	run.AddBreak()

	for _, par := range letter.Paragraphs {
		para = templateDoc.AddParagraph()
		para.SetStyle("Normal") // style name taken from the doc.Styles.Styles()
		para.AddRun().AddText(par)
		run = para.AddRun()
		run.AddBreak()
	}

	if len(templateDoc.Footers()) > 0 {
		f := templateDoc.Footers()[0]
		templateDoc.BodySection().SetFooter(f, wml.ST_HdrFtrDefault)
		para = f.AddParagraph()
		run := para.AddRun()
		run.AddBreak()
	}

	err = templateDoc.SaveToFile(outputName)
	return err
}
