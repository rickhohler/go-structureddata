package structureddata

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

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

// XMLParser XML parser
type XMLParser struct {
	text []byte
	err  error
}

// Unmarshall Unmarshall
func (p XMLParser) Unmarshall() (ParserData, error) {
	var data ParserData
	err := xml.Unmarshal(p.text, &data)
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
			// XML
		} else if startChar == "<" {
			return XMLParser{text: text}
		}
	}
	return nil
}

type SearchResult struct {
	key   string
	value string
	level int
}

// SearchKey Search key
func SearchKey(data ParserData, expression string) (bool, []SearchResult) {
	val := searchMap(data, expression, 0)

	return len(val) > 0, val
}

func searchMap(data ParserData, expression string, level int) []SearchResult {
	ret := []SearchResult{}
	switch data.(type) {
	case map[string]interface{}:
		if m := data.(map[string]interface{}); m != nil {
			if i, ok := m[expression]; ok {
				ret = append(ret, SearchResult{
					key:   expression,
					value: fmt.Sprintf("%v", i),
					level: level,
				})
			} else {
				level++
				for _, value := range m {
					switch v := value.(type) {
					case map[string]interface{}:
						return searchMap(v, expression, level)
					default:
					}
				}
			}
		}
	case []interface{}:
		if a := data.([]interface{}); a != nil {
			level++
			for _, v := range a {
				return searchMap(v, expression, level)
			}
		}

	default:
	}

	return ret
}
