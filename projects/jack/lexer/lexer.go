package lexer

import "jack/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) - 1 {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	ok := false

	for !ok {
		l.skipWhitespace()
		switch l.ch {
		case '{':
			ok = true
			tok = token.New(token.LBRACE, l.ch)
		case '}':
			ok = true
			tok = token.New(token.RBRACE, l.ch)
		case '(':
			ok = true
			tok = token.New(token.LPAREN, l.ch)
		case ')':
			ok = true
			tok = token.New(token.RPAREN, l.ch)
		case '[':
			ok = true
			tok = token.New(token.LBRACKET, l.ch)
		case ']':
			ok = true
			tok = token.New(token.RBRACKET, l.ch)
		case '.':
			ok = true
			tok = token.New(token.DOT, l.ch)
		case ',':
			ok = true
			tok = token.New(token.COMMA, l.ch)
		case ';':
			ok = true
			tok = token.New(token.SEMICOLON, l.ch)
		case '+':
			ok = true
			tok = token.New(token.PLUS, l.ch)
		case '-':
			ok = true
			tok = token.New(token.MINUS, l.ch)
		case '*':
			ok = true
			tok = token.New(token.ASTERISK, l.ch)
		case '/':
			if  l.peekChar() == '/' {
				l.skipLine()
			} else if l.peekChar() == '*' {
				l.skipComment()
			} else {
				tok = token.New(token.SLASH, l.ch)
				ok = true
			}
		case '&':
			ok = true
			tok = token.New(token.AND, l.ch)
		case '|':
			ok = true
			tok = token.New(token.OR, l.ch)
		case '<':
			ok = true
			tok = token.New(token.LT, l.ch)
		case '>':
			ok = true
			tok = token.New(token.GT, l.ch)
		case '=':
			ok = true
			tok = token.New(token.EQ, l.ch)
		case '~':
			ok = true
			tok = token.New(token.NOT, l.ch)
			case '"': 
			ok = true
			tok.Type = token.STRING
			tok.Literal = l.readString()
		case 0:
			ok = true
			tok.Literal = ""
			tok.Type = token.EOF
			
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				} else if isDigit(l.ch) {
					tok.Literal = l.readNumber()
					tok.Type = token.INT
				} else {
					tok = token.New(token.ILLEGAL, l.ch)
				}
			return tok
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipLine() {
	for l.ch != '\n' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	prev := l.ch
	for !(prev == '*' && l.ch == '/') || l.ch == 0 {
		prev = l.ch
		l.readChar()
	}
	l.readChar()
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}


// helper functions

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_'
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}