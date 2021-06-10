package tokenizer

import "encoding/xml"

type Tokenizer interface {
	Skip() error
	Token() (xml.Token, error)
}
