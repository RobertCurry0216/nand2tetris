package parser

import (
	"jack/lexer"
	"jack/token"
)

type Parser struct {
	lexer *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{ lexer: l }
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
