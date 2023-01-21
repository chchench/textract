package main

import (
	"encoding/xml"
	"flag"
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

var (
	outputPath = flag.String("outfile", "", "Path of file to dump into. If empty, output to standard out.")
	inputPath  = flag.String("infile", "", "Path of file to dump")
)

func main() {

	flag.Parse()

	if *inputPath == "" {
		fatalExit("Must specify an input file path")
	}

	xmlFile, err := os.Open(*inputPath)
	if err != nil {
		fatalExit(err.Error())

	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fatalExit(err.Error())
	}

	doc := Doc{}
	if err = xml.Unmarshal(byteValue, &doc); err != nil {
		fatalExit(err.Error())
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

	if *outputPath == "" {
		fmt.Println(text)
	} else {
		dump(*outputPath, text)
	}
}

func fatalExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func dump(path, content string) {
	f, err := os.Create(path)
	if err != nil {
		fatalExit(err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		fatalExit(err.Error())
	}
}
