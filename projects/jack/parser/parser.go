package parser

import (
	"errors"
	"fmt"
	"jack/ast"
	"jack/lexer"
	"jack/token"
	"strconv"
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

func (p *Parser) ParseFile() (string, error) {
	class, err := p.parseClassDeclaration()
	if err == nil {
		return class.String(), nil
	}
	return "", err
}

func (p *Parser) eatToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expect(token token.Type) bool {
	return p.curToken.Type == token
}

func (p *Parser) expectAndEat(token token.Type) bool {
	b := p.curToken.Type == token
	if b {	
		p.eatToken()
	}
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
	return errors.New(fmt.Sprintf("unexpected token, expected: %s   got %s", exp, got))
}

// ---------------------------------------------------------------------------------
// Expression parser function ------------------------------------------------------
// ---------------------------------------------------------------------------------

func (p *Parser) parseExpression() (ast.ExpressionNode, error) {
	var err error
	var exp ast.Expression

	// parse op
	if p.expectOp() {
		exp.Op = p.curToken
		p.eatToken()
	}

	// parse term
	switch p.curToken.Type {
		case token.INT:
			if exp.Term, err = p.parseIntLiteral(); err != nil {
				return nil, err
			}

		case token.STRING:
			if exp.Term, err = p.parseStringLiteral(); err != nil {
				return nil, err
			}

		case token.IDENT:
			switch p.peekToken.Type {
			case token.DOT: fallthrough
			case token.LPAREN:
				if exp.Term, err = p.parseSubroutineCall(); err != nil {
					return nil, err
				}
			case token.LBRACKET:
				if exp.Term, err = p.parseIndexIdentifier(); err != nil {
					return nil, err
				}
			default:
				if exp.Term, err = p.parseIdentifier(); err != nil {
					return nil, err
				}			
			}

		case token.TRUE: fallthrough
		case token.FALSE: fallthrough
		case token.NULL: fallthrough
		case token.THIS:
			if exp.Term, err = p.parseKeywordConstant(); err != nil {
				return nil, err
			}

		case token.LPAREN:
			p.eatToken()
			if exp.Term, err = p.parseExpression(); err != nil {
				return nil, err
			}
			if !p.expectAndEat(token.RPAREN) {
				return nil, tokenError(token.RPAREN, p.curToken.Literal)
			}

		case token.NOT: fallthrough
		case token.MINUS:
			if exp.Term, err = p.parseExpression(); err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("error parsing expression")
	}

	// parse tail
	if p.expectOp() {
		if exp.Tail, err = p.parseExpression(); err != nil {
			return nil, err
		}
	}

	return &exp, nil
}

func (p *Parser) expectOp() bool {
	if 
	p.expect(token.PLUS) ||
	p.expect(token.MINUS) ||
	p.expect(token.ASTERISK) ||
	p.expect(token.SLASH) ||
	p.expect(token.AND) ||
	p.expect(token.OR) ||
	p.expect(token.NOT) ||
	p.expect(token.GT) ||
	p.expect(token.LT) ||
	p.expect(token.EQ) {
		return true
	}
	return false
}

func (p *Parser) parseIdentifier() (*ast.Identifier, error) {
	i := &ast.Identifier{Token: p.curToken, Name: p.curToken.Literal}
	p.eatToken()
	return i, nil
}

func (p *Parser) parseIndexIdentifier() (*ast.IndexIdentifier, error) {
	ii := &ast.IndexIdentifier{Token: p.curToken, Name: p.curToken.Literal}

		p.eatToken()
		p.eatToken()

		index, err := p.parseExpression()

		if err != nil {
			return nil, err
		}

		ii.Index = index

		p.eatToken()

		return ii, nil
}

func (p *Parser) parseStringLiteral() (*ast.StringLiteral, error) {
	sl := &ast.StringLiteral{Token: p.curToken}
	sl.Value = p.curToken.Literal
	p.eatToken()

	return sl, nil
}

func (p *Parser) parseIntLiteral() (*ast.IntLiteral, error) {
	il := &ast.IntLiteral{Token: p.curToken}

	if i, err := strconv.Atoi(p.curToken.Literal); err == nil {
		il.Value = i
	} else {
		return nil, err
	}	
	p.eatToken()
	return il, nil
}

func (p *Parser) parseKeywordConstant() (*ast.KeywordConstant, error) {
	kw := &ast.KeywordConstant{Token: p.curToken}

	if
	p.curToken.Type == token.Type(token.TRUE) ||
	p.curToken.Type == token.Type(token.FALSE) ||
	p.curToken.Type == token.Type(token.NULL) ||
	p.curToken.Type == token.Type(token.THIS) {
		kw.Value = p.curToken.Literal
	} else {
		return nil, tokenError("true | false | null | this", p.curToken.Literal)
	}
	p.eatToken()
	return kw, nil
}

func (p *Parser) parseSubroutineCall() (*ast.SubroutineCall, error) {
	var err error
	sc := &ast.SubroutineCall{Token: p.curToken}

	// parse class
	if p.expectPeek(token.DOT) {
		if sc.Class, err = p.parseIdentifier(); err != nil {
			return nil, err
		}
		p.eatToken()
	}

	// parse name
	if sc.Name, err = p.parseIdentifier(); err != nil {
		return nil, err
	}

	// parse args
	if sc.Arguments, err = p.parseArgumentList(); err != nil {
		return nil, err
	}

	return sc, nil
}

func (p *Parser) parseArgumentList() ([]ast.ExpressionNode, error) {
	var args []ast.ExpressionNode

	if !p.expectAndEat(token.LPAREN) {
		return nil, tokenError(token.LPAREN, p.curToken.Literal)
	}

	for p.curToken.Type != token.RPAREN {
		if exp, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			args = append(args, exp)
		}

		p.expectAndEat(token.COMMA)

		if p.expect(token.RPAREN) {
			break
		}
	}
	p.eatToken()
	return args, nil
}


// ---------------------------------------------------------------------------------
// Statement parser function -------------------------------------------------------
// ---------------------------------------------------------------------------------

func (p *Parser) parseStatement() (ast.StatementNode, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.DO:
		return p.parseDoStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.STATIC: fallthrough
	case token.FIELD: fallthrough
	case token.VAR:
		return p.parseTypeDeclaration()
	case token.FUNCTION: fallthrough
	case token.METHOD: fallthrough
	case token.CONSTRUCTOR:
		return p.parseSubroutineDeclaration()
	case token.CLASS:
		return p.parseClassDeclaration()
	default:
		return nil, errors.New("error reading statement, unexpected token: " + p.curToken.Literal)
	}
}

// parseType => <type>
func (p *Parser) parseType() (token.Token, error) {
	switch p.curToken.Type {
		case token.INT: fallthrough
		case token.CHAR: fallthrough
		case token.BOOLEAN: fallthrough
		case token.VOID: fallthrough
		case token.IDENT:
			t := p.curToken
			p.eatToken()
			return t, nil
		default:
			return token.Token{}, tokenError("int | char | boolean | ident | void", p.curToken.Literal)
	}
}

// parseTypeDeclaration => var <type> <ident>;
func (p *Parser) parseTypeDeclaration() (*ast.TypeDeclaration, error) {
	var err error
	dec := &ast.TypeDeclaration{Token: p.curToken, Declaration: p.curToken}
	p.eatToken()

	if dec.Type, err = p.parseType(); err != nil {
		return nil, err
	}
	
	for {
		if !p.expect(token.IDENT) {
			return nil, tokenError(token.IDENT, p.curToken.Literal)
		}

		if i, err := p.parseIdentifier(); err == nil {
			dec.Names = append(dec.Names, i)
		} else {
			return nil, err
		}

		if p.expectAndEat(token.COMMA) {
			continue
		} else if p.expectAndEat(token.SEMICOLON) {
			break
		} else {
			return nil, tokenError(", or ;", p.curToken.Literal)
		}
	}

	return dec, nil
}


// parseParameterList => ( <type> <ident> {, <type> <ident>} )
func (p *Parser) parseParameterList() ([]*ast.ParamDeclaration, error) {
	var err error
	var params []*ast.ParamDeclaration

	if !p.expectAndEat(token.LPAREN) {
		return nil, tokenError("(", p.peekToken.Literal)
	}

	for p.curToken.Type != token.RPAREN {
		param := &ast.ParamDeclaration{Token: p.curToken}

		if param.Type, err = p.parseType(); err != nil {
			return nil, err
		}

		if param.Name, err = p.parseIdentifier(); err != nil {
			return nil, tokenError(token.IDENT, p.curToken.Literal)
		}

		params = append(params, param)

		if p.expect(token.RPAREN) {
			break
		}

		if !p.expectAndEat(token.COMMA) {
			return nil, tokenError(token.COMMA, p.curToken.Literal)
		}
	}

	p.eatToken()

	return params, nil
}


// parseSubroutineDeclaration => <dec> <returnType> <ident> ( <parameterList> ) <subroutineBody>
func (p *Parser) parseSubroutineDeclaration() (*ast.SubroutineDeclaration, error) {
	var err error
	dec := &ast.SubroutineDeclaration{Token: p.curToken, Decelration: p.curToken}
	p.eatToken()

	if dec.ReturnType, err = p.parseType(); err != nil {
		return nil, err
	}

	if !p.expect(token.IDENT) {
		return nil, tokenError(token.IDENT, p.curToken.Literal)
	}

	if dec.Name, err = p.parseIdentifier(); err != nil {
		return nil, err
	}

	if dec.Parameters, err = p.parseParameterList(); err != nil {
		return nil, err
	}


	if dec.Body, err = p.parseCodeBlock(); err != nil {
		return nil, err
	}

	return dec, nil
}

// parseClassDeclaration => class <name> { <statements> }
func (p *Parser) parseClassDeclaration() (*ast.ClassDeclaration, error) {
	var err error
	dec := &ast.ClassDeclaration{Token: p.curToken}

	if !p.peekAndEat(token.IDENT) {
		return nil, tokenError(token.IDENT, p.peekToken.Literal)
	}

	if dec.Name , err = p.parseIdentifier(); err != nil {
		return nil, err
	}

	if dec.Body, err = p.parseCodeBlock(); err != nil {
		return nil, err
	}

	return dec, nil
}

// parseLetStatement => let <Ident>[<expression>?] = <expression>;
func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.peekAndEat(token.IDENT) {
		return nil, tokenError(token.IDENT, p.peekToken.Literal)
	}

	if p.expectPeek(token.LBRACKET) {
		if ident, err := p.parseIndexIdentifier(); err == nil {
			stmt.Name = ident
		} else {
			return nil, err
		}
	} else {
		if ident, err := p.parseIdentifier(); err == nil {
			stmt.Name = ident
		} else {
			return nil, err
		}
	}


	if !p.expectAndEat(token.EQ) {
		return nil, tokenError(token.EQ, p.peekToken.Literal)
	}

	if val, err := p.parseExpression(); err == nil {
		stmt.Value = val
	} else {
		return nil, err
	}

	if !p.expectAndEat(token.SEMICOLON) {
		return nil, tokenError(token.SEMICOLON, p.peekToken.Literal)
	}

	return stmt, nil
}

// parseReturnStatement => return <expression>;
func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekAndEat(token.SEMICOLON) {
		p.eatToken()
		return stmt, nil
	}

	p.eatToken()

	if exp, err := p.parseExpression(); err == nil {
		stmt.Value = exp
	} else {
		return nil, err
	}

	if !p.expectAndEat(token.SEMICOLON) {
		return nil, tokenError(token.SEMICOLON, p.curToken.Literal)
	}

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
func (p *Parser) parseCodeBlock() ([]ast.StatementNode, error) {
	stmts := []ast.StatementNode{}

	if !p.expectAndEat(token.LBRACE)  {
		return nil, tokenError(token.LBRACE, p.curToken.Literal)
	}

	for p.curToken.Type != token.RBRACE {
		if s, err := p.parseStatement(); err == nil {
			stmts = append(stmts, s)
		} else {
			return nil, err
		}
	}

	if !p.expectAndEat(token.RBRACE)  {
		return nil, tokenError(token.LBRACE, p.curToken.Literal)
	}

	return stmts, nil
}


// parseWhileStatement => while (<exp>) {<statements>}
func (p *Parser) parseWhileStatement() (*ast.WhileStatement, error) {
	var err error
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.peekAndEat(token.LPAREN) {
		return nil, tokenError(token.LPAREN, p.curToken.Literal)
	}
	p.eatToken()

	if stmt.Expression, err = p.parseExpression(); err != nil {
		return nil, err
	}

	if !p.expectAndEat(token.RPAREN) {
		return nil, tokenError(token.RPAREN, p.curToken.Literal)
	}
	
	if stmt.Statements, err = p.parseCodeBlock(); err != nil {
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

	if !p.expectAndEat(token.RPAREN) {
		return nil, tokenError(token.RPAREN, p.curToken.Literal)
	}
	
	if stmts, err := p.parseCodeBlock(); err == nil {
		stmt.Statements = stmts
	} else {
		return nil, err
	}

	if p.expectAndEat(token.ELSE) {
		if stmts, err := p.parseCodeBlock(); err == nil {
			stmt.ElseStatements = stmts
		} else {
			return nil, err
		}
	}

	return stmt, nil
}