package codewriter

import (
	"fmt"
)

type Location string	

func (l *Location) Address() string {
	switch string(*l) {
	case "local":
		return "@LCL"
	case "argument":
		return "@ARG"
	case "this":
		return "@THIS"
	case "that":
		return "@THAT"
	case "static":
		return "@16"
	case "pointer":
		return "@THIS"
	case "temp":
		return "@5"
	default:
		panic(fmt.Sprintf("invalid location: %s", *l))
	}	
}

type Statement interface {
	Compile(*CodeWriter)
}

type LocationStatement struct {
	Location Location
	Argument int
}


type PushConstStatement struct {
	Argument int
}

func (s *PushConstStatement) String() string {
	return fmt.Sprintf("< push constant %d >", s.Argument)
}

func (s *PushConstStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// push constant %d", s.Argument)
	cw.Writeln("@%d", s.Argument)
	cw.Writeln("D=A")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")
}

func getBaseLocation(location Location) string {
	// use base location A if using a virtual memory segment
	// otherwise use M
	baseLocation := "M"
	virtual := []string{
		"static",
		"pointer",
		"temp",
	}

	for _, l := range virtual {
		if l == string(location) {
			baseLocation = "A"
			break
		}
	}

	return baseLocation
}

type PushLocationStatement struct {
	LocationStatement
}

func (s *PushLocationStatement) String() string {
	return fmt.Sprintf("< push %v %d >", s.Location, s.Argument)
}

func (s *PushLocationStatement) Compile(cw *CodeWriter) {
	base := getBaseLocation(s.Location)
	cw.Writeln("// push %v %v", s.Location, s.Argument)
	cw.Writeln(s.Location.Address())
	cw.Writeln("D=%s", base)
	cw.Writeln("@%d", s.Argument)
	cw.Writeln("A=D+A")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")
}


type PopStatement struct {
	LocationStatement
}

func (s *PopStatement) String() string {
	return fmt.Sprintf("< pop %v %d >", s.Location, s.Argument)
}

func (s *PopStatement) Compile(cw *CodeWriter) {
	base := getBaseLocation(s.Location)

	cw.Writeln("// pop %v %v", s.Location, s.Argument)

	// set A to location + arg
	cw.Writeln(s.Location.Address())
	cw.Writeln("D=%s", base)
	cw.Writeln("@%d", s.Argument)
	cw.Writeln("D=D+A")

	// store addr in R15
	cw.Writeln("@R15")
	cw.Writeln("M=D")

	// dec stack pointer
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")

	// load value
	cw.Writeln("D=M")
	cw.Writeln("@R15")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
}

type AddStatement struct {}
func (s *AddStatement) String() string { return "< add >"}
func (s *AddStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// add")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=D+M")
}

type SubStatement struct {}
func (s *SubStatement) String() string { return "< sub >"}
func (s *SubStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// sub")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=M-D")
}

type NegStatement struct {}
func (s *NegStatement) String() string { return "< neg >"}
func (s *NegStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// neg")
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("M=-M")
}

type EqStatement struct {
	Id int
}
func (s *EqStatement) String() string { return fmt.Sprintf("< Eq %d >", s.Id) }
func (s *EqStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// eq")

		// load top of stack into D
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")

		// check
	cw.Writeln("A=A-1")
	cw.Writeln("D=M-D")

		// in not greater
	cw.Writeln("@is-eq-%d", s.Id)
	cw.Writeln("D;JEQ")
	cw.Writeln("@end-%d", s.Id)
	cw.Writeln("D=0;JMP")

		// is greater
	cw.Writeln("(is-eq-%d)", s.Id)
	cw.Writeln("D=-1")

	cw.Writeln("(end-%d)", s.Id)
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("M=D")
}

type GtStatement struct {
	Id int
}
func (s *GtStatement) String() string { return fmt.Sprintf("< gt %d >", s.Id)}
func (s *GtStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// gt")

		// load top of stack into D
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")

		// check
	cw.Writeln("A=A-1")
	cw.Writeln("D=M-D")

		// in not greater
	cw.Writeln("@is-gt-%d", s.Id)
	cw.Writeln("D;JGT")
	cw.Writeln("@end-%d", s.Id)
	cw.Writeln("D=0;JMP")

		// is greater
	cw.Writeln("(is-gt-%d)", s.Id)
	cw.Writeln("D=-1")

	cw.Writeln("(end-%d)", s.Id)
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("M=D")
}

type LtStatement struct {
	Id int
}
func (s *LtStatement) String() string { return fmt.Sprintf("< lt %d >", s.Id)}
func (s *LtStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// lt")

		// load top of stack into D
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")

		// check
	cw.Writeln("A=A-1")
	cw.Writeln("D=M-D")

		// in not greater
	cw.Writeln("@is-lt-%v", s.Id)
	cw.Writeln("D;JLT")
	cw.Writeln("@end-%v", s.Id)
	cw.Writeln("D=0;JMP")

		// is greater
	cw.Writeln("(is-lt-%v)", s.Id)
	cw.Writeln("D=-1")

	cw.Writeln("(end-%v)", s.Id)
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("M=D")
}

type AndStatement struct { }
func (s *AndStatement) String() string { return "< and >"}
func (s *AndStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// and")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=D&M")
}

type OrStatement struct { }
func (s *OrStatement) String() string { return "< or >"}
func (s *OrStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// or")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=D|M")
}

type NotStatement struct { }
func (s *NotStatement) String() string { return "< not >"}
func (s *NotStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// not")
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("M=!M")
}

type LabelStatement struct { name string }
func (s *LabelStatement) String() string { return fmt.Sprintf("< label %s >", s.name)}
func (s *LabelStatement) Compile(cw *CodeWriter) {
	cw.Writeln("//label %s", s.name)
	cw.Writeln("(%s)", s.name)
}

type GotoStatement struct { name string }
func (s *GotoStatement) String() string { return fmt.Sprintf("< goto %s >", s.name)}
func (s *GotoStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// goto %s", s.name)
	cw.Writeln("@%s", s.name)
	cw.Writeln("0;JMP")
} 

