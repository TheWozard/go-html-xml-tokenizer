# HTML XML Tokenizer
Simple wrapper to the XML Decoder Token and Skip functions to support XML with embedded HTML tags.

Provides an interface for either the Simple Tokenizer or using the base XML Decoder
```go
type Tokenizer interface {
	Skip() error
	Token() (xml.Token, error)
}
```

## Simple Tokenizer
Once an xml.CharData token has been started it will continue to read tokens till the matching end token has been found. This will ignore all tags inside of CharData that are not at the start of the data.
```go
func NewSimpleHTMLXMLTokenizer(r io.Reader) Tokenizer
```