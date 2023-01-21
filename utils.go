package textract

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var UNZIPPED_DIR = ""

func extractArchiveContent(path string) (*[]string, error) {

	ar, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer ar.Close()

	dst := UNZIPPED_DIR

	for _, f := range ar.File {

		dstPath := filepath.Join(dst, f.Name)

		dir := filepath.Dir(dstPath)
		if dir != "" {
			createDir(dir)
		}

		mf, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(dstFile, mf); err != nil {
			log.Fatal(err)
		}

		mf.Close()
	}
}

func getFileType(fp string) (string, error) {

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
