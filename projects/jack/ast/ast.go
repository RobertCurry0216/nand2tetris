package ast

import (
	"jack/token"
	"strings"
)


type Node interface {
	TokenLiteral() string
	String() string
}


// Identifier denotes the name of a var or function
type Identifier struct {
	Token token.Token
	Name string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Name }

// StringLiteral denotes a string written in code
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string { return sl.Value }

// IntLiteral denotes a number value written in the code
type IntLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string { return il.Token.Literal }


// LetStatement is used when assigning vars
type LetStatement struct {
	Token token.Token
	Identifier *Identifier
	Value Node
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var sb strings.Builder
	sb.WriteString(ls.Token.Literal)
	sb.WriteString(" ")
	sb.WriteString(ls.Identifier.String())
	sb.WriteString(" = ")
	sb.WriteString(ls.Value.String())
	sb.WriteString(";")
	return sb.String()
}
