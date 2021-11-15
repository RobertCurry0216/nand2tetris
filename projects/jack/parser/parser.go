package parser

import (
	"errors"
	"fmt"
	"jack/ast"
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
	p.eatToken()
	p.eatToken()

	return p
}

func (p *Parser) eatToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectPeek(token token.Type) bool {
	if p.peekToken.Type == token {
		return true
	}
	return false
}

func (p *Parser) peekAndEat(token token.Type) bool {
	b := p.expectPeek(token)
	p.eatToken()
	return b
}

// error helpers
func tokenError(exp, got string) error {
	return errors.New(fmt.Sprintf("unexpected token, expected %s, got %s", exp, got))
}

// parser functions

func (p *Parser) parseExpression() (ast.Expression, error) {
	// TODO: parse expressions
	stmt := &ast.IntLiteral{Token: p.curToken}

	p.eatToken()

	return stmt, nil
}

// parseLetStatement => let <Ident> = <expression>;
func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.peekAndEat(token.IDENT) {
		return nil, tokenError(token.IDENT, p.peekToken.Literal)
	}

	stmt.Identifier = &ast.Identifier{Token: p.curToken, Name: p.curToken.Literal}

	if !p.peekAndEat(token.EQ) {
		return nil, tokenError(token.EQ, p.peekToken.Literal)
	}

	p.eatToken()

	if val, err := p.parseExpression(); err == nil {
		stmt.Value = val
	} else {
		return nil, err
	}

	if p.curToken.Type != token.SEMICOLON {
		return nil, tokenError(token.SEMICOLON, p.peekToken.Literal)
	}

	p.eatToken()

	return stmt, nil
}


// parseReturnStatement => return <expression>;
func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekAndEat(token.SEMICOLON) {
		p.eatToken()
		return stmt, nil
	}

	if exp, err := p.parseExpression(); err == nil {
		stmt.Value = exp
	} else {
		return nil, err
	}

	if p.curToken.Type != token.SEMICOLON {
		return nil, tokenError(token.SEMICOLON, p.curToken.Literal)
	}

	p.eatToken()

	return stmt, nil
}