all: clean test

build: build-tools

build-tools:
	(cd tools; go build -o dumpdocx dumpdocx.go)
	(cd tools; go build -o x2t-docx x2t-docx.go)
	(cd tools; go build -o x2t-pptx x2t-pptx.go)

test: build
	@echo "***** UNIT TESTS NOT YET PROVIDED *****"

clean:
	rm -f ./tools/{dumpdocx,x2t-docx,x2t-pptx}
	rm -rf ./output

.PHONY: all build test clean