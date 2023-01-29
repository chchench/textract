all: clean test

build: build-tools

build-tools:
	(cd tool-dumparchive; go build dumparchive.go)
	(cd tool-extract; go build extract.go)
	(cd tool-pptx; go build pptx.go)

test: build
	go test -v

clean:
	rm -f tool-dumparchive/dumparchive
	rm -f tool-extract/extract
	rm -f tool-pptx/pptx

.PHONY: all build test clean