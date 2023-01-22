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

type Slide struct {
	XMLName xml.Name `xml:"sld"`
	CSld    CSlide   `xml:"cSld"`
}

type CSlide struct {
	XMLName xml.Name `xml:"cSld"`
	Tree    SPTree   `xml:"spTree"`
}

type SPTree struct {
	XMLName xml.Name `xml:"spTree"`
	SPs     []SP     `xml:"sp"`
}

type SP struct {
	XMLName xml.Name `xml:"sp"`
	Body    TextBody `xml:"txBody"`
}

type TextBody struct {
	XMLName    xml.Name    `xml:"txBody"`
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

	xmlFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened input file")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	slide := Slide{}
	xml.Unmarshal(byteValue, &slide)

	// Each document probably only has one (1) body, but let's still
	// iterate thru for now until we're certain.

	text := ""

	for i := 0; i < len(slide.CSld.Tree.SPs); i++ {

		for j := 0; j < len(slide.CSld.Tree.SPs[i].Body.Paragraphs); j++ {

			t := ""
			for k := 0; k < len(slide.CSld.Tree.SPs[i].Body.Paragraphs[j].Runs); k++ {
				text += slide.CSld.Tree.SPs[i].Body.Paragraphs[j].Runs[k].Text
			}

			text += t + "\n"
		}
	}

	fmt.Println(text)
}
