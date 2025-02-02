/*
 * license/offline/main.go:
 * Illustrates how to load an offline (perpetual) license key.
 * Offline keys can be purchased at https://www.unidoc.io
 *
 * Run as: go run main.go
 */

package main

import (
	"fmt"

	"github.com/unidoc/unioffice/v2/common/license"
)

// Example of an offline perpetual license key.
const offlineLicenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
contents here.
-----END UNIDOC LICENSE KEY-----
`

func init() {
	// The customer name needs to match the entry that is embedded in the signed key.
	customerName := `My Company`

	// Good to load the license key in `init`. Needs to be done prior to using the library, otherwise operations
	// will result in an error.
	err := license.SetLicenseKey(offlineLicenseKey, customerName)
	if err != nil {
		panic(err)
	}
}

func main() {
	lk := license.GetLicenseKey()
	if lk == nil {
		fmt.Printf("Failed retrieving license key")
		return
	}
	fmt.Printf("License: %s\n", lk.ToString())
}
