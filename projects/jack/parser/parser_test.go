package parser

import (
	"jack/ast"
	"jack/lexer"
	"testing"
)


func assert(t *testing.T, n string, a, b interface{}) {
	if a != b {
		t.Fatalf("%s : expected %v <%T>, got %v <%T>", n, a, a, b, b)
	}
}

func TestParseTypeDecelration(t *testing.T) {
	type dec struct {
		expDec string
		expType string
		expIdent string
	}

	tests := []struct {
		input string
		expDecs []dec
	}{
		{"var int x;", []dec{{"var", "int", "x"}}},
		{"static boolean x, y, z;", []dec{{"static", "boolean", "x"}, {"static", "boolean", "y"}, {"static", "boolean", "z"}}},
		{"field char a;", []dec{{"field", "char", "a"}}},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseTypeDeclaration()

		if err != nil {
			t.Fatalf(err.Error())
		}

		assert(t, "TypeDecelration", len(test.expDecs), len(actual))

		for i, d := range actual {	
			assert(t, "TypeDecelration", test.expDecs[i].expDec, d.Declaration.Literal)
			assert(t, "TypeDecelration", test.expDecs[i].expType, d.Type.Literal)
			assert(t, "TypeDecelration", test.expDecs[i].expIdent, d.Name.Name)
		}
	}
}


func TestParseLetStatement(t *testing.T) {
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


		ident, ok := actual.Identifier.(*ast.Identifier)
		if !ok {
			t.Fatalf("expected *ast.Identifier, got %T", ident)
		}

		assert(t, "LetStatement", test.expName, ident.Name)
		assert(t, "LetStatement", test.expValue, actual.Value.String())
	}
}


func TestParseIndexedLetStatement(t *testing.T) {
	tests := []struct {
		input string
		expName string
		expIndex string
		expValue interface{}
	}{
		{"let x[0] = 8;", "x", "0", "8"},
		{"let y[foo] = true;", "y", "foo", "true"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseLetStatement()

		if err != nil {
			t.Fatalf(err.Error())
		}


		ident, ok := actual.Identifier.(*ast.IndexIdentifier)
		if !ok {
			t.Fatalf("LetIndexedStatement : expected *ast.IndexIdentifier, got %T", ident)
		}

		assert(t, "LetIndexedStatement", test.expName, ident.Name)
		assert(t, "LetIndexedStatement", test.expIndex, ident.Index.String())
		assert(t, "LetIndexedStatement", test.expValue, actual.Value.String())
	}
}


func TestParseReturnStatement(t *testing.T) {
	tests := []struct {
		input string
		expValue interface{}
	}{
		{"return;", nil},
		{"return 3;", "3"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseReturnStatement()

		if err != nil {
			t.Fatalf(err.Error())
		}

		if test.expValue == nil {
			assert(t, "ReturnStatement", test.expValue, actual.Value)
		} else {
			assert(t, "ReturnStatement", test.expValue, actual.Value.String())
		}
	}
}


func TestParseDoStatement(t *testing.T) {
	tests := []struct {
		input string
		expValue interface{}
	}{
		{"do foobar;", "foobar"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseDoStatement()

		if err != nil {
			t.Fatalf(err.Error())
		}

		assert(t, "DoStatement", test.expValue, actual.Expression.String())
	}
}


func TestParseWhileStatement(t *testing.T) {
	test := "while (true) { let x = 3; do foobar; }"

	lexer := lexer.New(test)
	parser := New(lexer)

	actual, err := parser.parseWhileStatement()

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert(t, "WhileStatement", actual.Expression.String(), "true")
	assert(t, "WhileStatement", len(actual.Statements), 2)
	assert(t, "WhileStatement", actual.Statements[0].String(), "let x = 3;")
	assert(t, "WhileStatement", actual.Statements[1].String(), "do foobar;")
}


func TestParseIfStatement(t *testing.T) {
	test := `
	if (bool) {
		let x = 4;
		do foobar;
	} else {
		let y = 4;
	}`

	lexer := lexer.New(test)
	parser := New(lexer)

	actual, err := parser.parseIfStatement()

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert(t, "IfStatement", actual.Expression.String(), "bool")
	assert(t, "IfStatement", len(actual.Statements), 2)
	assert(t, "IfStatement", len(actual.ElseStatements), 1)
	assert(t, "IfStatement", actual.Statements[0].String(), "let x = 4;")
	assert(t, "IfStatement", actual.Statements[1].String(), "do foobar;")
	assert(t, "IfStatement", actual.ElseStatements[0].String(), "let y = 4;")
}