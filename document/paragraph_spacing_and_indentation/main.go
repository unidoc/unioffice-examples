package main

import (
	"os"

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

	// Create paragraph and set first line indent and line spacing.
	p0 := doc.AddParagraph()
	p0.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p0.SetFirstLineIndent(measurement.Inch * 5)
	p0.SetLineSpacing(measurement.Millimeter*15, wml.ST_LineSpacingRuleExact)

	// Create paragraph and set alignment of paragraph to center,
	// and set above and below spacing to 2 inch.
	// Use wml.ST_JcCenter to set it center,
	// use wml.ST_JcRight to set it right,
	// use wml.ST_JcLeft to set it left,
	// use wml.ST_JcBoth to set it both (justified).
	p1 := doc.AddParagraph()
	p1.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p1.SetAlignment(wml.ST_JcCenter)
	p1.SetBeforeSpacing(measurement.Inch * 2)
	p1.SetAfterSpacing(measurement.Inch * 2)

	// Create paragraph and set left indent and add below spacing to 1 inch.
	p2 := doc.AddParagraph()
	p2.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p2.SetLeftIndent(measurement.Inch * 3)
	p2.SetAfterSpacing(measurement.Inch * 1)

	// Create paragraph and set right indent and set above spacing to 1 inch.
	p3 := doc.AddParagraph()
	p3.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p3.SetRightIndent(measurement.Inch * 3)
	p3.SetBeforeSpacing(measurement.Inch * 1)

	// Create paragraph and set above and below spacing with ignore space between if same style.
	p4 := doc.AddParagraph()
	p4.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p4.SetBeforeSpacing(measurement.Inch * 1)
	p4.SetAfterSpacing(measurement.Inch * 1)
	p4.IgnoreSpaceBetweenParagraphOfSameStyle()

	// Create paragraph and set outline level 3.
	p5 := doc.AddParagraph()
	p5.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p5.SetBeforeSpacing(measurement.Inch * 1)
	p5.SetAfterSpacing(measurement.Inch * 1)
	p5.SetOutlineLvl(3)

	// Create paragraph and set hanging indent 2 inch.
	p6 := doc.AddParagraph()
	p6.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p6.SetBeforeSpacing(measurement.Inch * 2)
	p6.SetAfterSpacing(measurement.Inch * 2)
	p6.SetHangingIndent(measurement.Inch * 2)

	// singlePoint is constant of the text height, it's 12 points.
	const singlePoint measurement.Distance = 12

	// Create paragraph and set line spacing to single.
	p7 := doc.AddParagraph()
	p7.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p7.SetBeforeSpacing(measurement.Inch * 0.5)
	p7.SetAfterSpacing(measurement.Inch * 0.5)
	p7.SetLineSpacing(1*singlePoint*measurement.Point, wml.ST_LineSpacingRuleAuto)

	// Create paragraph and set line spacing to one and half.
	p8 := doc.AddParagraph()
	p8.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p8.SetBeforeSpacing(measurement.Inch * 0.5)
	p8.SetAfterSpacing(measurement.Inch * 0.5)
	p8.SetLineSpacing(1.5*singlePoint*measurement.Point, wml.ST_LineSpacingRuleAuto)

	// Create paragraph and set line spacing to double.
	p9 := doc.AddParagraph()
	p9.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p9.SetBeforeSpacing(measurement.Inch * 0.5)
	p9.SetAfterSpacing(measurement.Inch * 0.5)
	p9.SetLineSpacing(2*singlePoint*measurement.Point, wml.ST_LineSpacingRuleAuto)

	// Create paragraph and set line spacing to muliple 3.
	p10 := doc.AddParagraph()
	p10.AddRun().AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	p10.SetBeforeSpacing(measurement.Inch * 0.5)
	p10.SetAfterSpacing(measurement.Inch * 0.5)
	p10.SetLineSpacing(3*singlePoint*measurement.Point, wml.ST_LineSpacingRuleAuto)

	doc.SaveToFile("out.docx")
}
