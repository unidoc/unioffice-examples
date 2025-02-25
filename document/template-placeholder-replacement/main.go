// Copyright 2025 FoxyUtils ehf. All rights reserved.
//
// This example shows how to fill a template document with placeholder strings.
// The placeholders are replaced with the provided strings.
// The placeholders are defined within double curly braces, e.g. {{TITLE}}.
// The replacement strings are provided in a map, where the key is the placeholder and the value is the replacement string.
// The filled document is saved to a new file named template-placeholder-replacement.docx.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv("UNIDOC_LICENSE_API_KEY"))
	if err != nil {
		panic(err)
	}
}

func main() {
	doc, err := document.Open("report.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	replacementStrings := map[string]string{
		"TITLE":         "Annual Report",
		"YEAR":          "2025",
		"DATE":          "1st January 2025",
		"COMPANY NAME":  "UniDoc",
		"AUTHOR":        "John Doe",
		"Title Heading": "Title 1",
		"Subtitle":      "Subtitle 1",
	}

	err = fillTemplate(doc, replacementStrings)
	if err != nil {
		log.Fatalf("error while filling template: %v\n", err.Error())
	}

	// Save new doucment.
	err = doc.SaveToFile("template-placeholder-replacement.docx")
	if err != nil {
		log.Fatalf("error while saving file: %v\n", err.Error())
	}
}

func fillTemplate(doc *document.Document, replacements map[string]string) error {
	for _, para := range doc.Paragraphs() {
		runs := para.Runs()
		if len(runs) == 0 {
			continue
		}

		var buffer strings.Builder
		var runIndex []int
		inPlaceholder := false

		for i, run := range runs {
			text := run.Text()
			buffer.WriteString(text)
			runIndex = append(runIndex, i)

			if strings.Contains(buffer.String(), "{{") && !inPlaceholder {
				inPlaceholder = true
				runIndex = []int{i}
			}

			if inPlaceholder && strings.Contains(buffer.String(), "}}") {
				fullText := buffer.String()
				isBreak := false
				pla := extractPlaceholder(fullText)
				fmt.Println("Extracted placeholder:", pla)

				for placeholder, replacement := range replacements {
					if placeholder == pla {
						updateRuns(runs, runIndex, placeholder, replacement)
						buffer.Reset()
						runIndex = nil
						inPlaceholder = false
						isBreak = true
					}
				}

				if isBreak {
					continue
				}

				buffer.Reset()
				runIndex = nil
				inPlaceholder = false
				isBreak = true
			}
		}
	}

	if err := doc.Validate(); err != nil {
		return fmt.Errorf("validate document failed: %w", err)
	}

	return nil
}

func extractPlaceholder(text string) string {
	// Get text within double curly braces.
	start := strings.Index(text, "{{")
	end := strings.Index(text, "}}")
	if start == -1 || end == -1 {
		return ""
	}

	return text[start+2 : end]
}

func updateRuns(runs []document.Run, runIndex []int, placeholder string, replacement string) {
	for _, idx := range runIndex {
		for _, EgRuninnerContent := range runs[idx].X().EG_RunInnerContent {
			if EgRuninnerContent == nil || EgRuninnerContent.RunInnerContentChoice.T == nil {
				continue
			}

			if strings.Contains(EgRuninnerContent.RunInnerContentChoice.T.Content, "{{") {
				EgRuninnerContent.RunInnerContentChoice.T.Content = strings.ReplaceAll(EgRuninnerContent.RunInnerContentChoice.T.Content, "{{", "")
			}

			if strings.Contains(EgRuninnerContent.RunInnerContentChoice.T.Content, "}}") {
				EgRuninnerContent.RunInnerContentChoice.T.Content = strings.ReplaceAll(EgRuninnerContent.RunInnerContentChoice.T.Content, "}}", "")
			}

			if strings.Contains(EgRuninnerContent.RunInnerContentChoice.T.Content, placeholder) {
				EgRuninnerContent.RunInnerContentChoice.T.Content = strings.ReplaceAll(EgRuninnerContent.RunInnerContentChoice.T.Content, placeholder, replacement)
			}
		}
	}
}
