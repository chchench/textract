package textract

import (
	"encoding/xml"
)

// Not sure if the XML structure is really as simple as below, but based on
// limited reverse-engineering done so far, this appears to be the case.
// Will need to enhance testing over time to confirm.

type Doc struct {
	XMLName xml.Name `xml:"document"`
	Bodies  []Body   `xml:"body"`
}

type Body struct {
	XMLName    xml.Name    `xml:"body"`
	Paragraphs []Paragraph `xml:"p"`
}

type Paragraph struct {
	XMLName xml.Name `xml:"p"`
	Runs    []Run    `xml:"r"`
}

type Run struct {
	XMLName xml.Name `xml:"r"`
	Text    string   `xml:"t"`
}
