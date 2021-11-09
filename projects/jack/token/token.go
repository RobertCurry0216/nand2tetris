package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

func New(t Type, l byte) Token {
	return Token{Type: t, Literal: string(l)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// operators
	EQ       = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	NOT = "~"
	AND = "&"
	OR  = "|"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// keywords
	CLASS       = "CLASS"
	CONSTRUCTOR = "CONSTRUCTOR"
	FUNCTION    = "FUNCTION"
	METHOD      = "METHOD"
	FIELD       = "FIELD"
	STATIC      = "STATIC"
	VAR         = "VAR"
	CHAR        = "CHAR"
	BOOLEAN     = "BOOLEAN"
	VOID        = "VOID"
	TRUE        = "TRUE"
	FALSE       = "FALSE"
	NULL        = "NULL"
	THIS        = "THIS"
	LET         = "LET"
	DO          = "DO"
	IF          = "IF"
	ELSE        = "ELSE"
	WHILE       = "WHILE"
	RETURN      = "RETURN"
)

var keywords = map[string]Type{
	"class":       CLASS,
	"constructor": CONSTRUCTOR,
	"function":    FUNCTION,
	"method":      METHOD,
	"field":       FIELD,
	"static":      STATIC,
	"var":         VAR,
	"int":         INT,
	"char":        CHAR,
	"boolean":     BOOLEAN,
	"void":        VOID,
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"this":        THIS,
	"let":         LET,
	"do":          DO,
	"if":          IF,
	"else":        ELSE,
	"while":       WHILE,
	"return":      RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}