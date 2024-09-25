// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

// letter represents a given message with letter content and receiver name.
type letter struct {
	Receiver   string   `json:"receiver"`
	Paragraphs []string `json:"paragraphs"`
}

func main() {
	letters, err := loadLetters("./data/letters.json")
	if err != nil {
		panic(err)
	}

	for _, letter := range letters {
		generateDoc("letter_template.docx", letter, fmt.Sprintf("./output/letter_to_%s.docx", letter.Receiver))
	}
}

// loadLetters loads the list of letters from json file
func loadLetters(dataPath string) ([]letter, error) {
	// Open the JSON file
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// Read the contents of the file
	b, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// letters holds the objects to hold the data.
	var letters []letter

	// get letter data from json
	err = json.Unmarshal(b, &letters)
	if err != nil {
		return nil, err
	}

	return letters, nil
}

// generateDoc creates a docx file using the template provided by `templatePath` and writes it to `outputName`.
func generateDoc(templatePath string, l letter, outputName string) error {
	templateDoc, err := document.OpenTemplate(templatePath)
	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}
	defer templateDoc.Close()

	// We can confirm the existence of the styles by printing them one by one
	// for _, s := range templateDoc.Styles.Styles() {
	// 	fmt.Println("style", s.Name(), "has ID of", s.StyleID(), "type is", s.Type())
	// }

	// set header of the document if it exists.
	if len(templateDoc.Headers()) > 0 {
		h := templateDoc.Headers()[0]
		templateDoc.BodySection().SetHeader(h, wml.ST_HdrFtrDefault)
		para := h.AddParagraph()
		run := para.AddRun()
		run.AddBreak()
	}

	t := time.Now()
	dateTime := t.Format("January 2, 2006")

	// take the formatting from the template for date time text
	para := templateDoc.AddParagraph()
	para.SetStyle("Normal") // style name taken from the doc.Styles.Styles() list
	para.AddRun().AddText(dateTime)
	run := para.AddRun()
	run.AddBreak()

	intro := fmt.Sprintf("Dear %s,", l.Receiver)
	para = templateDoc.AddParagraph()
	para.SetStyle("Normal") // style name taken from the doc.Styles.Styles()
	para.AddRun().AddText(intro)

	for _, par := range l.Paragraphs {
		para = templateDoc.AddParagraph()
		para.SetStyle("Normal") // style name taken from the doc.Styles.Styles()
		para.AddRun().AddText(par)
		run = para.AddRun()
		run.AddBreak()
	}

	// set footer if it exists.
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
