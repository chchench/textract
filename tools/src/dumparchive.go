package dumparchive

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chchench/textract"
)

var (
	outputDir = flag.String("outdir", "output", "Directory to dump extracted content into")
	target    = flag.String("file", "", "File to dump")
)

func main() {

	flag.Parse()

	if *target == "" {
		log.Fatal("An input filename/path should be specified using the -file flag\n")
	}

	ft, _ := textract.GetTrueFileType(*target)
	if ft != "application/zip" {
		log.Fatal("This file is not an archive\n")
	}

	createDir(*outputDir)

	list, err := textract.ExtractArchiveContent(*target, func(string) bool { return true })
	if err != nil {
		log.Fatal("Unable to extract file archive content\n")
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
			log.Fatal(err.Error())
		}

		if _, err := dstFile.Write(f.Data); err != nil {
			log.Fatal(err.Error())
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
