package main

import (
	"fmt"
	"os"

	"github.com/chchench/textract"
)

func main() {

	text, err := textract.RetrieveTextFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(text)
}
