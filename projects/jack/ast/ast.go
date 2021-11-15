package ast

import (
	"fmt"
	"jack/token"
	"strings"
)


type Node interface {
	TokenLiteral() string
	String() string
}

// expressions -----------------------------------------------------------------
type Expression interface {
	Node
	Expression()
}

// Identifier denotes the name of a var or function
type Identifier struct {
	Token token.Token
	Name string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Name }
func (i *Identifier) Expression() { }

// StringLiteral denotes a string written in code
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string { return sl.Value }
func (sl *StringLiteral) Expression() { }

// IntLiteral denotes a number value written in the code
type IntLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string { return il.Token.Literal }
func (il *IntLiteral) Expression() { }


// statements -----------------------------------------------------------------

type Statement interface {
	Node
	Statement()
}

// LetStatement is used when assigning vars
type LetStatement struct {
	Token token.Token
	Identifier *Identifier
	Value Expression
}

func (ls *LetStatement) Statement() {}

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


// ReturnStatement => return <exp?>;
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) Statement(){}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	if rs.Value == nil {
		return fmt.Sprintf("%s;", rs.TokenLiteral())
	}
	return fmt.Sprintf("%s %s;", rs.TokenLiteral(), rs.Value.String())
}