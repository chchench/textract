package textract

import (
	"encoding/xml"
)

/***************************************************************************
    XML structure based on public information and reverse engineering
***************************************************************************/

// Not sure if the XML structure is really as simple as below, but based on
// limited reverse-engineering done so far, this appears to be the case.
// Will need to enhance testing over time to confirm.

type Pptx_Slide struct {
	XMLName xml.Name    `xml:"sld"`
	CSld    Pptx_CSlide `xml:"cSld"`
}

type Pptx_CSlide struct {
	XMLName xml.Name    `xml:"cSld"`
	Tree    Pptx_SPTree `xml:"spTree"`
}

type Pptx_SPTree struct {
	XMLName xml.Name  `xml:"spTree"`
	SPs     []Pptx_SP `xml:"sp"`
}

type Pptx_SP struct {
	XMLName xml.Name      `xml:"sp"`
	Body    Pptx_TextBody `xml:"txBody"`
}

type Pptx_TextBody struct {
	XMLName    xml.Name         `xml:"txBody"`
	Paragraphs []Pptx_Paragraph `xml:"p"`
}

type Pptx_Paragraph struct {
	XMLName xml.Name   `xml:"p"`
	Runs    []Pptx_Run `xml:"r"`
}

type Pptx_Run struct {
	XMLName xml.Name `xml:"r"`
	Text    string   `xml:"t"`
}

/***************************************************************************
      Data structure for various data parsed from this document type
***************************************************************************/

type PptxParser struct {
	Content []MemberFileContent
}

/***************************************************************************
               Functions required for the document interface
***************************************************************************/

func (d *PptxParser) extension() string {
	return ".pptx"
}

func (d *PptxParser) trueType() string {
	return "application/zip"
}

func (d *PptxParser) filter(identifier string) bool {
	return identifier == "word/document.xml"
}

func (d *PptxParser) readFile(path string) error {
	list, err := ExtractArchiveContent(path, d.filter)
	if err != nil {
		return err
	}

	d.Content = *list

	return nil
}

func (d *PptxParser) retrieveTextFromFile() (string, error) {
	overallText := ""

	for _, mfc := range d.Content {
		text, err := d.docXML2Text(mfc.Identifier, mfc.Data)
		if err != nil {
			return "", err
		}
		overallText += text
	}

	return overallText, nil
}

func (d *PptxParser) docXML2Text(identifier string, byteData []byte) (string, error) {

	slide := Pptx_Slide{}

	if err := xml.Unmarshal(byteData, &slide); err != nil {
		return "", err
	}

	var text string

	// Each document probably only has one (1) body, but let's still
	// iterate thru for now until we're certain.

	for i := 0; i < len(slide.CSld.Tree.SPs); i++ {

		for j := 0; j < len(slide.CSld.Tree.SPs[i].Body.Paragraphs); j++ {

			t := ""
			for k := 0; k < len(slide.CSld.Tree.SPs[i].Body.Paragraphs[j].Runs); k++ {
				text += slide.CSld.Tree.SPs[i].Body.Paragraphs[j].Runs[k].Text
			}

			text += t + "\n"
		}
	}

	return text, nil
}
