package textract

import (
	"errors"
)

func RetrieveTextFromFile(path string) (string, error) {

	ft, err := GetTrueFileType(path)
	if err != nil {
		return "", err
	}

	ext := GetFileExtension(path)

	p1 := DocxParser{}
	parsers := []DocumentParser{&p1}

	for _, p := range parsers {
		if ft == p.trueType() && ext == p.extension() {
			err := p.readFile(path)
			if err != nil {
				return "", nil
			}

			return p.retrieveTextFromFile()
		}
	}

	return "", errors.New("unsupported file format")
}
