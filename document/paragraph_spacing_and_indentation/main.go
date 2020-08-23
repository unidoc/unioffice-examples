package main

import (
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func main() {
	doc := document.New()

	p0 := doc.AddParagraph()
	p0.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p0.SetFirstLineIndent(measurement.Inch * 5)
	p0.SetLineSpacing(measurement.Millimeter * 15, wml.ST_LineSpacingRuleExact)

	doc.SaveToFile("out.docx")
}
