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
type ExpressionNode interface {
	Node
	Expression()
}

// Expression is a tree like structure to denote any combination of expressions
type Expression struct {
	Term ExpressionNode
	Tail ExpressionNode
	Op *token.Token
}

func (e *Expression) TokenLiteral() string { return e.Term.TokenLiteral() }
func (e *Expression) Expression(){}
func (e *Expression) String() string {
	var sb strings.Builder

	if e.Op != nil {
		sb.WriteString(e.Op.Literal)
	}
	sb.WriteString(e.Term.String())

	if e.Tail != nil {
		sb.WriteString(e.Tail.String())
	}

	return sb.String()
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
	Index ExpressionNode
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
	Value int
}

func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string { return il.Token.Literal }
func (il *IntLiteral) Expression() { }


// KeywordConstant => true | false | null | this
type KeywordConstant struct {
	Token token.Token
	Value string
}

func (kc *KeywordConstant) TokenLiteral() string { return kc.Token.Literal }
func (kc *KeywordConstant) Expression(){}
func (kc *KeywordConstant) String() string { return kc.Value }

// SubroutinCall => <class name?>.<sub name>(<expression list>)
type SubroutineCall struct {
	Token token.Token
	Class *Identifier
	Name *Identifier
	Arguments []ExpressionNode
}

func (sc *SubroutineCall) TokenLiteral() string { return sc.Token.Literal }
func (sc *SubroutineCall) Expression(){}
func (sc *SubroutineCall) String() string {
	var sb strings.Builder

	if sc.Class != nil {
		sb.WriteString(sc.Class.String())
		sb.WriteString(".")
	}

	sb.WriteString(sc.Name.String())

	sb.WriteString("(")
	for i, arg := range sc.Arguments {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(arg.String())
	}
	sb.WriteString(")")

	return sb.String()
}

type UnaryExpression struct {
	Token token.Token
	Prefix token.Token
	Term ExpressionNode
}

func (ue *UnaryExpression) TokenLiteral() string { return ue.Token.Literal }
func (ue *UnaryExpression) Expression() {}
func (ue *UnaryExpression) String() string {
	return ue.Prefix.Literal + ue.Term.String()
}

type ParenExpression struct {
	Token token.Token
	Term ExpressionNode
}

func (pe *ParenExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *ParenExpression) Expression() {}
func (pe *ParenExpression) String() string { return "(" + pe.Term.String() + ")"}	


// ----------------------------------------------------------------------------
// statements -----------------------------------------------------------------
// ----------------------------------------------------------------------------

type StatementNode interface {
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
	Name *Identifier
	Parameters []*ParamDeclaration
	Body []StatementNode
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
	Name *Identifier
	Body []StatementNode
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
	Name ExpressionNode
	Value ExpressionNode
}

func (ls *LetStatement) Statement() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var sb strings.Builder
	sb.WriteString(ls.Token.Literal)
	sb.WriteString(" ")
	sb.WriteString(ls.Name.String())
	sb.WriteString(" = ")
	sb.WriteString(ls.Value.String())
	sb.WriteString(";")
	return sb.String()
}


// ReturnStatement => return <exp?>;
type ReturnStatement struct {
	Token token.Token
	Value ExpressionNode
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
	Expression ExpressionNode
}

func (ds *DoStatement) Statement(){}
func (ds *DoStatement) TokenLiteral() string { return ds.Token.Literal }

func (ds *DoStatement) String() string {
	return fmt.Sprintf("%s %s;", ds.TokenLiteral(), ds.Expression.String())
}


// WhileStatement => while (<expression>) {<statements>}
type WhileStatement struct {
	Token token.Token
	Expression ExpressionNode
	Statements []StatementNode
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
	Expression ExpressionNode
	Statements []StatementNode
	ElseStatements []StatementNode
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