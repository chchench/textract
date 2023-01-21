package extract

type document interface {
	extension() string
	trueType() string
	doc2Text(xml string) (string, error)
}
