package parser

import (
	"jack/lexer"
	"testing"
)


func TestParseLestStatement(t *testing.T) {
	tests := []struct {
		input string
		expName string
		expValue interface{}
	}{
		{"let x = 8;", "x", "8"},
		{"let y = true;", "y", "true"},
		{"let foo = bar;", "foo", "bar"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseLetStatement()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if actual.Identifier.Name != test.expName {
			t.Fatalf("error parsing let statement name, exp: %s, got %s", test.expName, actual.Identifier.Name)
		}

		if actual.Value.String() != test.expValue {
			t.Fatalf("error parsing let statement value, exp %v, got %v", test.expValue, actual.Value)
		}
	}
}