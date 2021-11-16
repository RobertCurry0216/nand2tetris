package lexer

import (
	"jack/token"
	"testing"
)

func TestLexer(t *testing.T) {
	const input = `
		class foo {
			function bar() {
				let x = 5;
				do func("hello", "world");
				let y = 1 + 2 * 5;
				return;
			}
		}
	`

	expected := []token.Token{
		{Type: token.CLASS, Literal: "class"},
		{Type: token.IDENT, Literal: "foo"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.FUNCTION, Literal: "function"},
		{Type: token.IDENT, Literal: "bar"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.EQ, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.DO, Literal: "do"},
		{Type: token.IDENT, Literal: "func"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.STRING, Literal: "hello"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "world"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.EQ, Literal: "="},
		{Type: token.INT, Literal: "1"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "2"},
		{Type: token.ASTERISK, Literal: "*"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.EOF, Literal: ""},

	}

	l := New(input)

	for i, tt := range expected {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.Literal, tok.Literal)
		}
	}
}