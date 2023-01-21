// package extract
package textract

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type DocumentParser interface {
	extension() string
	trueType() string
	readFile(string) error
	retrieveTextFromFile() (string, error)
}

type Filter func(string) bool

type MemberFileContent struct {
	Identifier string
	Data       []byte
}

func extractArchiveContent(path string, filter Filter) (*[]MemberFileContent, error) {

	ar, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer ar.Close()

	var list []MemberFileContent

	for _, f := range ar.File {

		if filter != nil {
			if !(filter)(f.Name) {
				continue
			}
		}

		mf, err := f.Open()
		if err != nil {
			return nil, err
		}

		buf, err := io.ReadAll(mf)
		if err != nil {
			return nil, err
		}

		mf.Close()

		mfc := &MemberFileContent{
			Identifier: f.Name,
			Data:       buf,
		}

		list = append(list, *mfc)
	}

	return &list, nil
}

func getTrueFileType(fp string) (string, error) {

	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buf), nil
}

func getFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
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
