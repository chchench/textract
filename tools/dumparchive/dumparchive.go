package main

import (
	"flag"
	"fmt"
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
		fatalExit("An input filename/path should be specified using the -file flag\n")
	}

	if getTrueFileType(*target) != "application/zip" {
		fatalExit("This file is not an archive\n")
	}

	createDir(*outputDir)

	list, err := extractArchiveContent(*target, func(string) bool { return true })
	if err != nil {
		fatalExit("Unable to extract file archive content\n")
	}

	dst := *outputDir

	for _, f := range *list {

		dstPath := filepath.Join(dst, f.Identifier)

		dir := filepath.Dir(dstPath)
		if dir != "" {
			createDir(dir)
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			fatalExit(err.Error())
		}

		if _, err := dstFile.Write(f.Data); err != nil {
			fatalExit(err.Error())
		}

		dstFile.Close()
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
