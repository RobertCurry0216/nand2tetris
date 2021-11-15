package parser

import (
	"fmt"
	"jack/lexer"
	"testing"
)


func TestParseLetStatement(t *testing.T) {
	tests := []struct {
		input string
		expName string
		expValue interface{}
	}{
		{"let x = 8;", "x", "8"},
		{"let y = true;", "y", "true"},
		{"let foo = bar;", "foo", "bar"},
		{"let bar[0] = false", "bar[0]", "false"},
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


func TestParseReturnStatement(t *testing.T) {
	tests := []struct {
		input string
		expValue interface{}
	}{
		{"return;", nil},
		{"return 3;", 3},
		{"return foobar;", "foobar"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseReturnStatement()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if fmt.Sprintf("%v", actual.Value) != fmt.Sprintf("%v", test.expValue) {
			t.Fatalf("error parsing let statement value, exp %v, got %v", test.expValue, actual.Value)
		}
	}
}