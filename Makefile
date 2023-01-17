all: clean test

build:
	(cd tools; go build -o dumpdocx dumpdocx.go; go build -o x2t-docx x2t-docx.go)

test: build
	@echo "***** UNIT TESTS NOT YET PROVIDED *****"

clean:
	rm -f ./tools/{dumpdocx, x2t-docx}
	rm -rf ./output

.PHONY: all build test clean