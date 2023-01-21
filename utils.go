package main

import (
	"fmt"
	"net/http"
	"os"
)

var UNZIPPED_DIR = ""

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
