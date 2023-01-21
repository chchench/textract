package extract

type document interface {
	extension() string
	trueType() string
	xml2Text(xml string) string
}
