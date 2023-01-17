package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	outputDir = flag.String("outdir", "output", "Directory to dump extracted content into")
	target    = flag.String("file", "", "File to dump")
)

func main() {

	flag.Parse()

	if *target == "" {
		fmt.Fprintf(os.Stderr, "An input filename/path should be specified using the -file flag\n")
		os.Exit(1)
	}

	createDir(*outputDir)

	ar, err := zip.OpenReader(*target)
	if err != nil {
		log.Fatal(err)
	}
	defer ar.Close()

	dst := *outputDir

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

	fmt.Println("Program finished successfully")
}

func createDir(dp string) {
	_, err := os.Stat(dp)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dp, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}
