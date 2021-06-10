package tokenizer

import (
	"encoding/xml"
	"fmt"
	"io"
)

func NewSimpleHTMLXMLTokenizer(r io.Reader) Tokenizer {
	return &SimpleHTMLXMLTokenizer{
		Decoder: xml.NewDecoder(r),
	}
}

type SimpleHTMLXMLTokenizer struct {
	Decoder *xml.Decoder
	prev    xml.StartElement
	buffer  xml.Token
}

func (t *SimpleHTMLXMLTokenizer) Skip() error {
	_, err := t.Token()
	return err
}

func (t *SimpleHTMLXMLTokenizer) Token() (xml.Token, error) {
	if t.buffer != nil {
		defer func() { t.buffer = nil }()
		return t.buffer, nil
	}
	token, err := t.Decoder.Token()
	if err != nil {
		return token, err
	}
	switch parsed := token.(type) {
	case xml.CharData:
		// Make a copy to not interfer with the internal xml decoder buffer
		data := fmt.Sprintf("%s", parsed)
		for {
			token, err := t.Decoder.Token()
			if err != nil {
				return token, err
			}
			switch eval := token.(type) {
			case xml.StartElement:
				data = fmt.Sprintf("%s<%s>", data, eval.Name.Local)
			case xml.EndElement:
				if eval.Name.Local == t.prev.Name.Local {
					t.buffer = token
					return xml.CharData(data), nil
				}
				data = fmt.Sprintf("%s</%s>", data, eval.Name.Local)
			case xml.CharData:
				data = fmt.Sprintf("%s%s", data, eval)
			default:
				t.buffer = token
				return xml.CharData(data), nil
			}
		}
	case xml.StartElement:
		t.prev = parsed
		return token, nil
	default:
		return token, nil
	}
}
