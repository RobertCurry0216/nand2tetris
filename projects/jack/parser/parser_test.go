package parser

import (
	"jack/ast"
	"jack/lexer"
	"jack/token"
	"testing"
)


func assert(t *testing.T, n string, a, b interface{}) {
	if a != b {
		t.Fatalf("%s : expected: %v <%T>   got: %v <%T>", n, a, a, b, b)
	}
}

func TestParseTypeDecelration(t *testing.T) {
	tests := []struct {
		input string
		expDec string
		expType string
		expIdent []string
	}{
		{"var int x;", "var", "int", []string{"x"}},
		{"static boolean x, y, z;", "static", "boolean", []string{"x", "y", "z"}},
		{"field char a;", "field", "char", []string{"a"}},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseTypeDeclaration()

		if err != nil {
			t.Fatalf(err.Error())
		}


		assert(t, "TypeDeclaration", test.expDec, actual.Declaration.Literal)
		assert(t, "TypeDeclaration", test.expType, actual.Type.Literal)
		assert(t, "TypeDecelration", len(test.expIdent), len(actual.Names))

		for i, name := range actual.Names {	
			assert(t, "TypeDecelration", test.expIdent[i], name.TokenLiteral())
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


		ident, ok := actual.Name.(*ast.Identifier)
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


		ident, ok := actual.Name.(*ast.IndexIdentifier)
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

func TestParseParamList(t *testing.T) {
	type exp struct {
		dec string
		name string
	}

	type testType struct {
		input string
		decs []exp
	}


	tests := []testType{
		{"()", []exp{}},
		{"(int a)", []exp{{"int", "a"}}},
		{"(bool x, string y, Vector z)",
			[]exp{{"bool", "x"}, {"string", "y"}, {"Vector", "z"}},
		},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseParameterList()

		if err != nil {
			t.Fatalf(err.Error())
		}

		assert(t, "ParameterList", len(test.decs), len(actual))

		for i, dec := range actual {
			assert(t, "ParameterList", test.decs[i].dec, dec.Type.Literal)
			assert(t, "ParameterList", test.decs[i].name, dec.Name.String())
		}
	}
}


func TestParseSubroutineDeclaration(t *testing.T) {
	test := `
		method void sumValues(int a, int b) {
			let c = 1;
			do add;
			return c;
		}
	`

	lexer := lexer.New(test)
	parser := New(lexer)

	actual, err := parser.parseSubroutineDeclaration()

	if err != nil {
		t.Fatalf(err.Error())
	}

	n := "SubroutineDeclaration"

	assert(t, n, token.Type(token.METHOD), actual.Decelration.Type)
	assert(t, n, token.Type(token.VOID), actual.ReturnType.Type)
	assert(t, n, "sumValues", actual.Name.String())
	assert(t, n, 2, len(actual.Parameters))
	assert(t, n, 3, len(actual.Body))
}


func TestParseClassDeclaration(t *testing.T) {
	test := `
		class Vector {
			field int x, y;

			constructor Vector new(int _x, int _y) {
				let x = _x;
				let y = _y;
				return this;
			}
		}
	`

	lexer := lexer.New(test)
	parser := New(lexer)

	actual, err := parser.parseClassDeclaration()

	if err != nil {
		t.Fatalf(err.Error())
	}

	n := "ClassDeclaration"

	assert(t, n, token.Type(token.CLASS), actual.Token.Type)
	assert(t, n, "Vector", actual.Name.String())
	assert(t, n, 2, len(actual.Body))
}