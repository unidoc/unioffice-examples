// Copyright 2020 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

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

func main() {
	doc, err := document.Open("new_resume_001.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			for _, any := range r.X().Extra {
				if ac, ok := any.(*wml.AlternateContentRun); ok {
					for _, anchor := range ac.Choice.Drawing.Anchor {
						for _, any := range anchor.Graphic.GraphicData.Any {
							if wps, ok := any.(*wml.WdWsp); ok {
								if wps.WChoice != nil {
									fmt.Println("")
									fmt.Println("")
									for _, egcbc := range wps.WChoice.Txbx.TxbxContent.EG_ContentBlockContent {
										for _, p := range egcbc.P {
											fmt.Println("")
											for i, egpc := range p.EG_PContent {
												fmt.Println(i)
												for _, egcrc := range egpc.EG_ContentRunContent {
													run := egcrc.R
													for _, egric := range run.EG_RunInnerContent {
														if egric.T != nil {
															fmt.Println(egric.T.Content)
														}
													}
												}
												if hyperlink := egpc.Hyperlink; hyperlink != nil {
													fmt.Println("Hyperlink:")
													for _, egcrc := range hyperlink.EG_ContentRunContent {
														run := egcrc.R
														for _, egric := range run.EG_RunInnerContent {
															if egric.T != nil {
																fmt.Println(egric.T.Content)
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	doc.SaveToFile("result.docx")
}
