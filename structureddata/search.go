package structureddata

import "fmt"

// SearchResult Search result
type SearchResult struct {
	key   string
	value string
	level int
}

type filterFunc func(expression string, m map[string]interface{}) (interface{}, bool)

var simpleFilterFunc = func(expression string, m map[string]interface{}) (interface{}, bool) {
	i, ok := m[expression]
	return i, ok
}

// SearchKey Search key
func SearchKey(data ParserData, expression string, filterFunc filterFunc) (bool, []SearchResult) {
	val := searchData(data, expression, filterFunc, 0)

	return len(val) > 0, val
}

func searchData(data ParserData, expression string, filterFunc filterFunc, level int) []SearchResult {
	ret := []SearchResult{}
	switch data.(type) {
	case map[string]interface{}:
		if m := data.(map[string]interface{}); m != nil {
			if i, ok := filterFunc(expression, m); ok {
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
						return searchData(v, expression, filterFunc, level)
					default:
					}
				}
			}
		}
	case []interface{}:
		if a := data.([]interface{}); a != nil {
			level++
			for _, v := range a {
				return searchData(v, expression, filterFunc, level)
			}
		}

	default:
	}

	return ret
}
