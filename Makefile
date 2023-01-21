all: clean test

build: build-tools

build-tools:
	(cd tools; go build -o dumparchive src/dumparchive.go)
	(cd tools; go build -o pptx src/pptx.go)

test: build
	@echo "***** UNIT TESTS NOT YET PROVIDED *****"

clean:
	rm -f ./tools/{dumparchive,pptx}

.PHONY: all build test clean