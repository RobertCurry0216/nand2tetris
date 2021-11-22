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

// IndexIdentifier denotes a var name with index ie: foo[bar]
type IndexIdentifier struct {
	Token token.Token
	Name string
	Index Expression
}

func (ii *IndexIdentifier) TokenLiteral() string { return ii.Token.Literal }
func (ii *IndexIdentifier) String() string { 
	return fmt.Sprintf("%s[%s]", ii.Name, ii.Index.String())
}
func (ii *IndexIdentifier) Expression() { }

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

type TypeDeclaration struct {
	Token token.Token
	Declaration token.Token
	Type token.Token
	Names []*Identifier
}

func (td *TypeDeclaration) Statement() {}
func (td *TypeDeclaration) TokenLiteral() string {
	return td.Token.Literal
}

func (td *TypeDeclaration) String() string {
	var sb strings.Builder
	sb.WriteString(td.Declaration.Literal)
	sb.WriteString(" ")
	sb.WriteString(td.Type.Literal)
	sb.WriteString(" ")
	for i, name := range td.Names {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(name.String())
	}
	sb.WriteString(";")
	return sb.String()
}

type ParamDeclaration struct {
	Token token.Token
	Type token.Token
	Name *Identifier
}

func (pd *ParamDeclaration) Statement() {}
func (pd *ParamDeclaration) TokenLiteral() string {
	return pd.Token.Literal
}

func (pd *ParamDeclaration) String() string {
	var sb strings.Builder
	sb.WriteString(pd.Type.Literal)
	sb.WriteString(" ")
	sb.WriteString(pd.Name.String())
	return sb.String()
}

type SubroutineDeclaration struct {
	Token token.Token
	Decelration token.Token
	ReturnType token.Token
	Name Identifier
	Parameters []*ParamDeclaration
	Body []Statement
}

func (sd *SubroutineDeclaration) Statement() {}
func (sd *SubroutineDeclaration) TokenLiteral() string {
	return sd.Token.Literal
}

func (sd *SubroutineDeclaration) String() string {
	var sb strings.Builder
	sb.WriteString(sd.Decelration.Literal)
	sb.WriteString(" ")
	sb.WriteString(sd.ReturnType.Literal)
	sb.WriteString(" ")
	sb.WriteString(sd.Name.String())
	sb.WriteString("(")
	for i, param := range sd.Parameters {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(param.String())
	}

	sb.WriteString(") {")
	for _, stmt := range sd.Body {
		sb.WriteString("\t")
		sb.WriteString(stmt.String())
	}
	sb.WriteString("}")

	return sb.String()
}

type ClassDeclaration struct {
	Token token.Token
	Name Identifier
	Body []Statement
}

func (cd *ClassDeclaration) Statement(){}
func (cd *ClassDeclaration) TokenLiteral() string { 
	return cd.Token.Literal
}

func (cd *ClassDeclaration) String() string {
	var sb strings.Builder

	sb.WriteString(cd.TokenLiteral())
	sb.WriteString(" ")
	sb.WriteString(cd.Name.String())
	sb.WriteString(" {\n")

	for _, stmt := range cd.Body {
		sb.WriteString(stmt.String())
		sb.WriteString("\n")
	}

	sb.WriteString("}")

	return sb.String()
}	

// LetStatement is used when assigning vars
type LetStatement struct {
	Token token.Token
	Identifier Expression
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

// DoStatement => do <expression>;
type DoStatement struct {
	Token token.Token
	Expression Expression
}

func (ds *DoStatement) Statement(){}
func (ds *DoStatement) TokenLiteral() string { return ds.Token.Literal }

func (ds *DoStatement) String() string {
	return fmt.Sprintf("%s %s;", ds.TokenLiteral(), ds.Expression.String())
}


// WhileStatement => while (<expression>) {<statements>}
type WhileStatement struct {
	Token token.Token
	Expression Expression
	Statements []Statement
}

func (ws *WhileStatement) Statement(){}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }

func (ws *WhileStatement) String() string {
	var sb strings.Builder
	sb.WriteString(ws.TokenLiteral())
	sb.WriteString("(")
	sb.WriteString(ws.Expression.String())
	sb.WriteString(") {\n")
	for _, stmt := range ws.Statements {
		sb.WriteString("\t")
		sb.WriteString(stmt.String())
		sb.WriteString("\n")
	}
	sb.WriteString("}")

	return sb.String()
}


// IfStatement => if (<exp>) {<statements>} ?else {<statements>}
type IfStatement struct {
	Token token.Token
	Expression Expression
	Statements []Statement
	ElseStatements []Statement
}

func (is *IfStatement) Statement() {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }

func (is *IfStatement) String() string {
	var sb strings.Builder

	sb.WriteString(is.TokenLiteral())
	sb.WriteString(" (")
	sb.WriteString(is.Expression.String())
	sb.WriteString(") {\n")
	for _, s := range is.Statements {
		sb.WriteString("\t")
		sb.WriteString(s.String())
		sb.WriteString("\n")
	}
	sb.WriteString("}")

	if len(is.ElseStatements) > 0 {
		sb.WriteString(" eles {\n")
		for _, s := range is.ElseStatements {
			sb.WriteString("\t")
			sb.WriteString(s.String())
			sb.WriteString("\n")
		}
		sb.WriteString("}")
	}

	return sb.String()
}