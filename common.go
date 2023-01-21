// package extract
package main

import (
	"archive/zip"
	"io"
)

type document interface {
	extension() string
	xml2Text(string, []byte) (string, error)
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
		defer mf.Close()

		buf, err := io.ReadAll(mf)
		if err != nil {
			return nil, err
		}

		mfc := &MemberFileContent{
			Identifier: f.Name,
			Data:       buf,
		}

		list = append(list, *mfc)
	}

	return &list, nil
}
