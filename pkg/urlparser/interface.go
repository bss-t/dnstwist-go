package urlparser

type Parser interface {
	Parse(urlString string) ParsedUrl
}
