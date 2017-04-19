package structureddata

import (
	"fmt"
	"testing"
)

// Test level 1 successful search.
func TestSearchKeyTrue1(t *testing.T) {
	p := NewParser(jsonTestData)
	data, err := p.Unmarshall()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	if found, value := SearchKey(data, "Age", simpleFilterFunc); !found {
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
	if found, _ := SearchKey(data, "Type", simpleFilterFunc); !found {
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
	if found, _ := SearchKey(data, "Name", simpleFilterFunc); !found {
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
	if found, _ := SearchKey(data, "NotFoundKey", simpleFilterFunc); found {
		t.Fatal("Expected false search")
	}
}
