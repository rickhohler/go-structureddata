package structureddata

import (
	"fmt"
	"testing"
)

// jsonTestData JSON test data
var jsonTestData = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"], "Classes": {"Type": "English"}}`)
var jsonTestArrayData = []byte(`[{"Name":"Wednesday","Age":6},{"Name":"Tuesday","Age":9}]`)

// xmlTestData XML test data
var xmlTestData = []byte(`<Name>Wednesday</Name>`)

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

func TestNewXMLParser(t *testing.T) {
	p := NewParser(xmlTestData)
	//	typ := reflect.TypeOf(p)
	_, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

// Test level 1 successful search.
func TestSearchKeyTrue1(t *testing.T) {
	p := NewParser(jsonTestData)
	data, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if found, value := SearchKey(data, "Age"); !found {
		fmt.Println("key=" + value[0].key)
		t.Fatal("Expected true search")
	}
}

// Test level 2 successful search.
func TestSearchKeyTrue2(t *testing.T) {
	p := NewParser(jsonTestData)
	data, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if found, _ := SearchKey(data, "Type"); !found {
		t.Fatal("Expected true search")
	}
}

// Test array successful search.
func TestSearchKeyTrue3(t *testing.T) {
	p := NewParser(jsonTestArrayData)
	data, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if found, _ := SearchKey(data, "Name"); !found {
		t.Fatal("Expected true search")
	}
}

// Test unsuccessful search.
func TestSearchKeyFalse1(t *testing.T) {
	p := NewParser(jsonTestData)
	data, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if found, _ := SearchKey(data, "NotFoundKey"); found {
		t.Fatal("Expected false search")
	}
}
