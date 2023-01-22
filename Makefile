all: clean test

build: build-tools

build-tools:
	(cd tools/src; go build -o ../dumparchive dumparchive.go)
	(cd tools/src; go build -o ../pptx pptx.go)
	(cd tools/src; go build -o ../extract extract.go)

test: build
	@echo "***** UNIT TESTS NOT YET PROVIDED *****"

clean:
	rm -f ./tools/{dumparchive,pptx,extract}

.PHONY: all build test clean