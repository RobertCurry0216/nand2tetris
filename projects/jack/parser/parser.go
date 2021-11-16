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

func (p *Parser) checkAndEat(token token.Type) bool {
	b := p.curToken.Type == token
	p.eatToken()
	return b
}

func (p *Parser) expectPeek(token token.Type) bool {
	return p.peekToken.Type == token
}

func (p *Parser) peekAndEat(token token.Type) bool {
	b := p.expectPeek(token)
	if b {
		p.eatToken()
	}
	return b
}

// error helpers
func tokenError(exp, got string) error {
	return errors.New(fmt.Sprintf("unexpected token, expected %s, got %s", exp, got))
}

// Expression parser functions

func (p *Parser) parseExpression() (ast.Expression, error) {
	// TODO: parse expressions
	stmt := &ast.IntLiteral{Token: p.curToken}

	p.eatToken()

	return stmt, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	if p.expectPeek(token.LBRACKET) {
		ii := &ast.IndexIdentifier{Token: p.curToken, Name: p.curToken.Literal}

		p.eatToken()
		p.eatToken()

		index, err := p.parseExpression()

		if err != nil {
			return nil, err
		}

		ii.Index = index

		return ii, nil
	} else {
		i := &ast.Identifier{Token: p.curToken, Name: p.curToken.Literal}
		return i, nil
	}
}

// Statement parser function

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.DO:
		return p.parseDoStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	default:
		return nil, errors.New("error reading statement, unexpected token: " + p.curToken.Literal)
	}
}

// parseLetStatement => let <Ident>[<expression>?] = <expression>;
func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.peekAndEat(token.IDENT) {
		return nil, tokenError(token.IDENT, p.peekToken.Literal)
	}

	if ident, err := p.parseIdentifier(); err == nil {
		stmt.Identifier = ident
	} else {
		return nil, err
	}


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

// parseDoStatement => do <expression>;
func (p *Parser) parseDoStatement() (*ast.DoStatement, error) {
	stmt := &ast.DoStatement{Token: p.curToken}

	p.eatToken()

	if exp, err := p.parseExpression(); err == nil {
		stmt.Expression = exp
	} else {
		return nil, err
	}

	if p.curToken.Type != token.SEMICOLON {
		return nil, tokenError(token.SEMICOLON, p.curToken.Literal)
	}

	p.eatToken()

	return stmt, nil
}

// parseCodeBlock => {<statements>}
func (p *Parser) parseCodeBlock() ([]ast.Statement, error) {
	stmts := []ast.Statement{}

	if !p.checkAndEat(token.LBRACE)  {
		return nil, tokenError(token.LBRACE, p.curToken.Literal)
	}

	for p.curToken.Type != token.RBRACE {
		if s, err := p.parseStatement(); err == nil {
			stmts = append(stmts, s)
		} else {
			return nil, err
		}
	}

	if !p.checkAndEat(token.RBRACE)  {
		return nil, tokenError(token.LBRACE, p.curToken.Literal)
	}

	return stmts, nil
}


// parseWhileStatement => while (<exp>) {<statements>}
func (p *Parser) parseWhileStatement() (*ast.WhileStatement, error) {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.peekAndEat(token.LPAREN) {
		return nil, tokenError(token.LPAREN, p.curToken.Literal)
	}
	p.eatToken()

	if exp, err := p.parseExpression(); err == nil {
		stmt.Expression = exp
	} else {
		return nil, err
	}

	if !p.checkAndEat(token.RPAREN) {
		return nil, tokenError(token.RPAREN, p.curToken.Literal)
	}
	
	if stmts, err := p.parseCodeBlock(); err == nil {
		stmt.Statements = stmts
	} else {
		return nil, err
	}

	return stmt, nil
}


// parseIfStatement => if (<exp>) {<statements>} ? else {<statements>}
func (p *Parser) parseIfStatement() (*ast.IfStatement, error) {
	stmt := &ast.IfStatement{Token: p.curToken}

	if !p.peekAndEat(token.LPAREN) {
		return nil, tokenError(token.LPAREN, p.curToken.Literal)
	}
	p.eatToken()

	if exp, err := p.parseExpression(); err == nil {
		stmt.Expression = exp
	} else {
		return nil, err
	}

	if !p.checkAndEat(token.RPAREN) {
		return nil, tokenError(token.RPAREN, p.curToken.Literal)
	}
	
	if stmts, err := p.parseCodeBlock(); err == nil {
		stmt.Statements = stmts
	} else {
		return nil, err
	}

	if p.checkAndEat(token.ELSE) {
		if stmts, err := p.parseCodeBlock(); err == nil {
			stmt.ElseStatements = stmts
		} else {
			return nil, err
		}
	}

	return stmt, nil
}