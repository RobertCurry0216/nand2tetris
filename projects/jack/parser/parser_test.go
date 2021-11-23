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
		{"let sum = sum + a[i];", "sum", "(sum (+a[(i)]))"},
		{"let foo = bar();", "foo", "(bar())"},
		{"let x = 8;", "x", "(8)"},
		{"let y = ~true;", "y", "(~true)"},
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
		{"let x[0] = 8;", "x", "(0)", "(8)"},
		{"let y[foo] = true;", "y", "(foo)", "(true)"},
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
		{"return 3;", "(3)"},
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
		{"do foobar(1, 2);", "(foobar((1), (2)))"},
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

	assert(t, "WhileStatement", "(true)", actual.Expression.String())
	assert(t, "WhileStatement", 2, len(actual.Statements))
	assert(t, "WhileStatement", "let x = (3);", actual.Statements[0].String())
	assert(t, "WhileStatement", "do (foobar);", actual.Statements[1].String())
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

	assert(t, "IfStatement", "(bool)", actual.Expression.String())
	assert(t, "IfStatement", 2, len(actual.Statements))
	assert(t, "IfStatement", 1, len(actual.ElseStatements))
	assert(t, "IfStatement", "let x = (4);", actual.Statements[0].String())
	assert(t, "IfStatement", "do (foobar);", actual.Statements[1].String())
	assert(t, "IfStatement", "let y = (4);", actual.ElseStatements[0].String())
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
		method void sumValues() {
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
	assert(t, n, 0, len(actual.Parameters))
	assert(t, n, 3, len(actual.Body))
}


func TestParseClassDeclaration(t *testing.T) {
	test := `
		class Vector {
			field int x, y;

			constructor Vector new() {
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


func TestParseExpression(t *testing.T) {
	tests := []struct{
			input string
			exp string
		}{
			{"a[i]", "(a[(i)])"},
			{"1 * 2 + 3", "(1 (*2 (+3)))"},
			{"(4 * 8) - (2 / 3)", "((4 (*8)) (-(2 (/3))))"},
			{"-1", "(-1)"},
			{"Vector.norm()", "(Vector.norm())"},
			{"foo(bar)", "(foo((bar)))"},
			{"true", "(true)"},
			{"this", "(this)"},
			{"~~false", "(~(~false))"},
			{"-1 + (-3)", "(-1 (+(-3)))"},
	}

	for _, test := range tests {
		lexer := lexer.New(test.input)
		parser := New(lexer)

		actual, err := parser.parseExpression()

		if err != nil {
			t.Fatalf(err.Error())
		}

		assert(t, "paresExpression", test.exp, actual.String())
	}
}


func TestLarge(t *testing.T) {
	test := `
		class Main {
			function void main() {
					var Array a;
					var int length;
					var int i, sum;
		
		let length = Keyboard.readInt("HOW MANY NUMBERS? ");
		let a = Array.new(length);
		let i = 0;
		
		while (i < length) {
				let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
				let i = i + 1;
		}
		
		let i = 0;
		let sum = 0;
		
		while (i < length) {
				let sum = sum + a[i];
				let i = i + 1;
		}
		
		do Output.printString("THE AVERAGE IS: ");
		do Output.printInt(sum / length);
		do Output.println();
		
		return;
			}
	}`

	lexer := lexer.New(test)
	parser := New(lexer)

	_, err := parser.ParseFile()

	assert(t, "large", nil, err)
}