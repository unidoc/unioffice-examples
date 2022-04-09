// Copyright 2022 UniDoc ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
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

var mapList = []map[string][]string{
	map[string][]string{"Preferred programming language": []string{"Go", "Java", "PHP"}},
	map[string][]string{"Which sport you love to play": []string{"Football", "Basketball", "Diving"}},
	map[string][]string{"Another information": []string{"This A", "This B", "This C"}},
}

func main() {
	doc := document.New()
	defer doc.Close()

	// Create numbering definition.
	nd := doc.Numbering.AddDefinition()

	// Add level to number definition with decimal format.
	lvl := nd.AddLevel()
	lvl.SetFormat(wml.ST_NumberFormatDecimal)
	lvl.SetAlignment(wml.ST_JcLeft)
	lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(1) * measurement.Inch)

	// Sets the numbering level format.
	lvl.SetText("%1.")

	for i := 0; i < len(mapList); i++ {
		for key, val := range mapList[i] {
			para := doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.SetNumberingLevel(0)
			para.AddRun().AddText(key)

			// Create children numbering definition.
			ndChildren := doc.Numbering.AddDefinition()

			// Add level to number definition with lower roman format.
			lvl := ndChildren.AddLevel()
			lvl.SetFormat(wml.ST_NumberFormatLowerRoman)
			lvl.SetAlignment(wml.ST_JcLeft)
			lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(2) * measurement.Inch)

			// Sets the numbering level format.
			lvl.SetText("%1.")

			for i := 0; i < len(val); i++ {
				para := doc.AddParagraph()
				para.SetNumberingDefinition(ndChildren)
				para.SetNumberingLevel(0)
				para.AddRun().AddText(val[i])

				// Add more nested numbering.
				if i == 0 && key == "Another information" {
					// Create children numbering definition.
					ndChild := doc.Numbering.AddDefinition()

					// Add level to number definition with upper letter format.
					lvl := ndChild.AddLevel()
					lvl.SetFormat(wml.ST_NumberFormatUpperLetter)
					lvl.SetAlignment(wml.ST_JcLeft)
					lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(3) * measurement.Inch)

					// Sets the numbering level format.
					lvl.SetText("%1.)")
					for i := 1; i < 5; i++ {
						p := doc.AddParagraph()
						p.SetNumberingLevel(0)
						p.SetNumberingDefinition(ndChild)
						run := p.AddRun()
						run.AddText(fmt.Sprintf("More Level %d", i))
					}
				}
			}
		}
	}

	// Numbering bullet.
	ndBullet := doc.Numbering.Definitions()[0]
	for i := 1; i < 5; i++ {
		p := doc.AddParagraph()
		p.SetNumberingLevel(i - 1)
		p.SetNumberingDefinition(ndBullet)
		run := p.AddRun()
		run.AddText(fmt.Sprintf("Level %d", i))
	}

	doc.SaveToFile("bullet-and-numbering.docx")
}
