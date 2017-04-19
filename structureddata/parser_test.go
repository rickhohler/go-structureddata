package structureddata

import "testing"

// jsonTestData JSON test data
var jsonTestData = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"], "Classes": {"Type": "English"}}`)
var jsonTestArrayData = []byte(`[{"Name":"Wednesday","Age":6},{"Name":"Tuesday","Age":9}]`)

func TestNewNilParser(t *testing.T) {
	b := []byte(`invalid structured data`)
	p := NewParser(b)
	if p != nil {
		t.Fatal("Expected nil")
	}
}

func TestNewJSONParser(t *testing.T) {
	p := NewParser(jsonTestData)
	//	typ := reflect.TypeOf(p)
	_, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
