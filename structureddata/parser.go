package structureddata

import "encoding/json"

// ParserData Parser data
type ParserData interface{}

// Parser interface
type Parser interface {
	Unmarshall() (data ParserData, err error)
}

// JSONParser JSON parser
type JSONParser struct {
	text []byte
	err  error
}

// Unmarshall Unmarshall
func (p JSONParser) Unmarshall() (ParserData, error) {
	var data ParserData
	err := json.Unmarshal(p.text, &data)
	return data, err
}

// NewParser Determine type of parser based upon input text
func NewParser(text []byte) Parser {
	if len(text) > 0 {
		startChar := string(text[0:1])
		endChar := string(text[len(text)-1])

		// JSON
		if (startChar == "{" || startChar == "[") &&
			(endChar == "}" || endChar == "]") {
			return JSONParser{text: text}
		}
	}
	return nil
}
