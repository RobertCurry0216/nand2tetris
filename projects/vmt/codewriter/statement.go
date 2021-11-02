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

func (s *LocationStatement) Address() string {
	return s.Location.Address()
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

type PushStaticStatement struct {
	File string
	Argument int
}

func (s *PushStaticStatement) String() string {
	return fmt.Sprintf("< push static %d >", s.Argument)
}

func (s *PushStaticStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// push static %d", s.Argument)
	cw.Writeln("@%s.%d", s.File, s.Argument)
	cw.Writeln("D=M")
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
	cw.Writeln("// push %v %v", s.Location, s.Argument)
	cw.Writeln(s.Address())

	if s.Argument > 0 {
		base := getBaseLocation(s.Location)
		cw.Writeln("D=%s", base)
		cw.Writeln("@%d", s.Argument)
		cw.Writeln("A=D+A")
	} else {
		cw.Writeln("A=M")
	}

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
	cw.Writeln(s.Address())
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

type PopStaticStatement struct {
	File string
	Argument int
}

func (s *PopStaticStatement) String() string {
	return fmt.Sprintf("< pop static %d>", s.Argument)
}

func (s *PopStaticStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// pop static %v", s.Argument)

	// set A to location + arg
	cw.Writeln("@%s.%d", s.File, s.Argument)
	cw.Writeln("D=A")

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

type AddStatement struct { }
func (s *AddStatement) String() string { return "< add >"}
func (s *AddStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// add")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=D+M")
}

type SubStatement struct { }
func (s *SubStatement) String() string { return "< sub >"}
func (s *SubStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// sub")
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("A=A-1")
	cw.Writeln("M=M-D")
}

type NegStatement struct { }
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

type LabelStatement struct { Name string; Function string }
func (s *LabelStatement) String() string { return fmt.Sprintf("< label %s >", s.Name)}
func (s *LabelStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// label %s.%s", s.Function, s.Name)
	cw.Writeln("(%s.%s)", s.Function, s.Name)
}

type GotoStatement struct { Name string; Function string }
func (s *GotoStatement) String() string { return fmt.Sprintf("< goto %s >", s.Name)}
func (s *GotoStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// goto %s.%s", s.Function, s.Name)
	cw.Writeln("@%s.%s", s.Function, s.Name)
	cw.Writeln("0;JMP")
} 

type IfGotoStatement struct { Name string; Id int; Function string }
func (s *IfGotoStatement) String() string { return fmt.Sprintf("< if-goto %s>", s.Name)}
func (s *IfGotoStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// if-goto %s.%s", s.Function, s.Name)

	// pop top of stack
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")

	// check eq 0
	cw.Writeln("@%s.%s$%d", s.Function, s.Name, s.Id)
	cw.Writeln("D;JEQ")
	cw.Writeln("@%s.%s", s.Function, s.Name)
	cw.Writeln("0;JMP")
	cw.Writeln("(%s.%s$%d)", s.Function, s.Name, s.Id)
}

type FunctionStatement struct { Name string; Nvars int }
func (s *FunctionStatement) String() string { return fmt.Sprintf("< function %s %d >", s.Name, s.Nvars)}
func (s *FunctionStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// function %s %d", s.Name, s.Nvars)
	cw.Writeln("(%s)", s.Name)

	// set local
	cw.Writeln("@SP")
	cw.Writeln("D=M")
	cw.Writeln("@LCL")
	cw.Writeln("AM=D")

	// init nvars
	for i := 0; i < s.Nvars; i++ {
		cw.Writeln("M=0")
		cw.Writeln("A=A+1")
	}

	// set SP
	cw.Writeln("D=A")
	cw.Writeln("@SP")
	cw.Writeln("M=D")
}

type CallStatement struct {
	Name string
	Nargs int
	Id int
}
func (s *CallStatement) String() string { return fmt.Sprintf("< call %s %d >", s.Name, s.Nargs) }
func (s *CallStatement) Compile(cw *CodeWriter) {
	returnId := fmt.Sprintf("%s$ret.%d", s.Name, s.Id)
	nargs := s.Nargs

	cw.Writeln("// call %s %d", s.Name, s.Nargs)

	// if narg == 0, push 0 onto the stack
	// and set narg to 1
	if nargs < 1 {
		cw.Writeln("@0")
		cw.Writeln("D=A")
		cw.Writeln("@SP")
		cw.Writeln("A=M")
		cw.Writeln("M=D")
		cw.Writeln("@SP")
		cw.Writeln("M=M+1")
		nargs = 1
	}

	// save return addr
	cw.Writeln("@%s", returnId)
	cw.Writeln("D=A")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")

	// save local
	cw.Writeln("@LCL")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")

	// save arg
	cw.Writeln("@ARG")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")

	// save this
	cw.Writeln("@THIS")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")

	// save that
	cw.Writeln("@THAT")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("A=M")
	cw.Writeln("M=D")
	cw.Writeln("@SP")
	cw.Writeln("M=M+1")

	// set arg
	cw.Writeln("@%d", nargs + 5)
	cw.Writeln("D=A")
	cw.Writeln("@SP")
	cw.Writeln("D=M-D")
	cw.Writeln("@ARG")
	cw.Writeln("M=D")

	// jump to function
	cw.Writeln("@%s", s.Name)
	cw.Writeln("0;JMP")


	// return addr
	cw.Writeln("(%s)", returnId)
}


type ReturnStatement struct { }
func (s *ReturnStatement) String() string { return "< return >" }
func (s *ReturnStatement) Compile(cw *CodeWriter) {
	cw.Writeln("// return")

	// set arg[0] to return val
	cw.Writeln("@SP")
	cw.Writeln("A=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@ARG")
	cw.Writeln("A=M")
	cw.Writeln("M=D")

	// set SP back to LCL[0]
	cw.Writeln("@LCL")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("M=D")

	// save arg[1] addr so it can be set back to SP
	cw.Writeln("@ARG")
	cw.Writeln("D=M+1")
	cw.Writeln("@R14")
	cw.Writeln("M=D")

	// reset that
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@THAT")
	cw.Writeln("M=D")

	// reset this
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@THIS")
	cw.Writeln("M=D")

	// reset arg
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@ARG")
	cw.Writeln("M=D")

	// reset lcl
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@LCL")
	cw.Writeln("M=D")

	// reset addr and restore SP
	cw.Writeln("@SP")
	cw.Writeln("AM=M-1")
	cw.Writeln("D=M")
	cw.Writeln("@R15")
	cw.Writeln("M=D")

	cw.Writeln("@R14")
	cw.Writeln("D=M")
	cw.Writeln("@SP")
	cw.Writeln("M=D")

	cw.Writeln("@R15")
	cw.Writeln("A=M;JMP")
}