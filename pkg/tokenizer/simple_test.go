package tokenizer_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-html-xml-tokenizer/pkg/tokenizer"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleHTMLXMLTokenizer(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		output []xml.Token
	}{
		{
			name:  "standard",
			input: "<data>some data</data>",
			output: []xml.Token{
				xml.StartElement{
					Name: xml.Name{Local: "data"},
					Attr: []xml.Attr{},
				},
				xml.CharData("some data"),
				xml.EndElement{
					Name: xml.Name{Local: "data"},
				},
			},
		},
		{
			name:  "unescaped html tags",
			input: "<data>some <b>important</b> data</data>",
			output: []xml.Token{
				xml.StartElement{
					Name: xml.Name{Local: "data"},
					Attr: []xml.Attr{},
				},
				xml.CharData("some <b>important</b> data"),
				xml.EndElement{
					Name: xml.Name{Local: "data"},
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test[%d]:%s", i+1, test.name), func(t *testing.T) {
			simple := tokenizer.NewSimpleHTMLXMLTokenizer(bytes.NewBufferString(test.input))
			index := 0
			for {
				token, err := simple.Token()
				if err == io.EOF {
					break
				}
				require.NoError(t, err)
				if index >= len(test.output) {
					require.Nil(t, token, "Unexpected additional token")
					require.Fail(t, "Both results returned where nil")
				}
				require.Equal(t, test.output[index], token)
				index++
			}

		})
	}

}
