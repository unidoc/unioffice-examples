// Copyright 2020 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/unidoc/unioffice/v2/document"
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
	doc, err := document.Open("new_resume_001.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			for _, any := range r.X().Extra {
				if ac, ok := any.(*wml.AlternateContentRun); ok {
					if ch := ac.Choice.Drawing; ch != nil {
						for _, dc := range ch.DrawingChoice {
							if anchor := dc.Anchor; anchor != nil {
								for _, any := range anchor.Graphic.GraphicData.Any {
									if wps, ok := any.(*wml.WdWsp); ok {
										if wps.WordprocessingShapeChoice1 != nil {
											fmt.Println("")
											fmt.Println("")
											for _, egcbc := range wps.WordprocessingShapeChoice1.Txbx.TxbxContent.EG_BlockLevelElts {
												if blc := egcbc.BlockLevelEltsChoice; blc != nil {
													for _, cbc := range blc.EG_ContentBlockContent {
														if cbc.ContentBlockContentChoice != nil {
															for _, p := range cbc.ContentBlockContentChoice.P {
																for i, egpc := range p.EG_PContent {
																	fmt.Println(i)
																	if pcc := egpc.PContentChoice; pcc != nil {
																		for _, egcrc := range pcc.EG_ContentRunContent {
																			if egcrc.ContentRunContentChoice != nil {
																				run := egcrc.ContentRunContentChoice.R
																				for _, egric := range run.EG_RunInnerContent {
																					if egric.RunInnerContentChoice != nil {
																						if egric.RunInnerContentChoice.T != nil {
																							fmt.Println(egric.RunInnerContentChoice.T.Content)
																						}
																					}
																				}
																			}
																		}
																		if hyperlink := pcc.Hyperlink; hyperlink != nil {
																			fmt.Println("Hyperlink:")
																			if hyperlink.PContentChoice != nil {
																				for _, egcrc := range hyperlink.PContentChoice.EG_ContentRunContent {
																					if egcrc.ContentRunContentChoice != nil {
																						run := egcrc.ContentRunContentChoice.R
																						for _, egric := range run.EG_RunInnerContent {
																							if egric.RunInnerContentChoice != nil {
																								if egric.RunInnerContentChoice.T != nil {
																									fmt.Println(egric.RunInnerContentChoice.T.Content)
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
