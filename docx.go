// package extract
package main

import (
	"encoding/xml"
)

/***************************************************************************
    XML structure based on public information and reverse engineering
***************************************************************************/

// Not sure if the XML structure is really as simple as below, but based on
// limited reverse-engineering done so far, this appears to be the case.
// Will need to enhance testing over time to confirm.

type Docx_Doc struct {
	XMLName xml.Name    `xml:"document"`
	Bodies  []Docx_Body `xml:"body"`
}

type Docx_Body struct {
	XMLName    xml.Name         `xml:"body"`
	Paragraphs []Docx_Paragraph `xml:"p"`
}

type Docx_Paragraph struct {
	XMLName xml.Name   `xml:"p"`
	Runs    []Docx_Run `xml:"r"`
}

type Docx_Run struct {
	XMLName xml.Name `xml:"r"`
	Text    string   `xml:"t"`
}

/***************************************************************************
      Data structure for various data parsed from this document type
***************************************************************************/

type Docx struct {
	Filepath string
	Content  []MemberFileContent
}

/***************************************************************************
               Functions required for the document interface
***************************************************************************/

func extension() string {
	return "TBD"
}

func xml2Text(identifier string, byteData []byte) (string, error) {

	doc := Docx_Doc{}
	if err := xml.Unmarshal(byteData, &doc); err != nil {
		return "", err
	}

	var text string

	// Each document probably only has one (1) body, but let's still
	// iterate thru for now until we're certain.

	for i := 0; i < len(doc.Bodies); i++ {
		for j := 0; j < len(doc.Bodies[i].Paragraphs); j++ {
			t := ""
			for k := 0; k < len(doc.Bodies[i].Paragraphs[j].Runs); k++ {
				text += doc.Bodies[i].Paragraphs[j].Runs[k].Text
			}
			text += t + "\n"
		}
	}

	return text, nil
}
