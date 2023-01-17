package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {

	xmlFile, err := os.Open("./output/word/document.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened input file")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	doc := Doc{}
	xml.Unmarshal(byteValue, &doc)

	// Each document probably only has one (1) body, but let's still
	// iterate thru for now until we're certain.

	text := ""

	for i := 0; i < len(doc.Bodies); i++ {
		for j := 0; j < len(doc.Bodies[i].Paragraphs); j++ {

			t := ""
			for k := 0; k < len(doc.Bodies[i].Paragraphs[j].Runs); k++ {
				text += doc.Bodies[i].Paragraphs[j].Runs[k].Text
			}

			text += t + "\n"
		}
	}

	fmt.Println(text)
}
