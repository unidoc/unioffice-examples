// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main
// This example demonstrates parsing range references with sheet names containing exclamation marks.

import (
	"fmt"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet/reference"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	parse("A1:A3")
	parse("Sheet1!A1:A3")
	parse("Sheet1!!A1:A3")
	parse("Shee!t1!A1:A3")
	parse("!Sheet1!A1:A3")
}

func parse(s string) {
	from, to, err := reference.ParseRangeReference(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(from.RowIdx, from.Column, from.SheetName)
	fmt.Println(to.RowIdx, to.Column, to.SheetName)
}
