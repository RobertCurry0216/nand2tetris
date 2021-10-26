package codewriter

import (
	"fmt"
)

type Location string

const (
	CONSTANT = Location("constant")
	LOCAL = Location("local")
)

type Statement interface {
	ToAsm() []string
}

type PushConstStatement struct {
	Argument int
}

func (s *PushConstStatement) ToAsm() []string	{
	return []string{
		fmt.Sprintf("// push constant %d", s.Argument),
		fmt.Sprintf("@%d", s.Argument),
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

type PushLocationStatement struct {
	Location Location
	Argument int
}

func (s *PushLocationStatement) ToAsm() []string {
	return []string{
		fmt.Sprintf("// push %v %v", s.Location, s.Argument),
		fmt.Sprintf("@%v", s.Location),
		"D=M",
		fmt.Sprintf("@%d", s.Argument),
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}


type PopStatement struct {
	Location Location
	Argument int
}

func (s *PopStatement) ToAsm() []string {
	return []string{
		fmt.Sprintf("// pop %v %v", s.Location, s.Argument),
		fmt.Sprintf("@%v", s.Location),
		"D=M",
		fmt.Sprintf("@%d", s.Argument),
		"D=D+A",
		"@R15",
		"M=D",
		"@SP",
		"AM=M-1",
		"D=M",
		"@R15",
		"A=M",
		"M=D",
	}
}

type AddStatement struct {}
func (s *AddStatement) ToAsm() []string {
	return []string{
		"// add",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D+M",
	}
}

type SubStatement struct {}
func (s *SubStatement) ToAsm() []string {
	return []string{
		"// sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D-M",
	}
}

type NegStatement struct {}
func (s *NegStatement) ToAsm() []string {
	return []string{
		"// neg",
		"@SP",
		"A=M-1",
		"M=-M",
	}
}

type EqStatement struct {
	Id int
}
func (s *EqStatement) ToAsm() []string {
	return []string{
		"// eq",

		// load top of stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		// check
		"A=A-1",
		"D=M-D",

		// in not greater
		fmt.Sprintf("@is-eq-%d", s.Id),
		"D;JEQ",
		fmt.Sprintf("@end-%d", s.Id),
		"D=0;JMP",

		// is greater
		fmt.Sprintf("(is-eq-%d)", s.Id),
		"D=-1",

		fmt.Sprintf("(end-%d)", s.Id),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

type GtStatement struct {
	Id int
}
func (s *GtStatement) ToAsm() []string {
	return []string{
		"// gt",

		// load top of stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		// check
		"A=A-1",
		"D=M-D",

		// in not greater
		fmt.Sprintf("@is-gt-%d", s.Id),
		"D;JGT",
		fmt.Sprintf("@end-%d", s.Id),
		"D=0;JMP",

		// is greater
		fmt.Sprintf("(is-gt-%d)", s.Id),
		"D=-1",

		fmt.Sprintf("(end-%d)", s.Id),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

type LtStatement struct {
	Id int
}
func (s *LtStatement) ToAsm() []string {
	return []string{
		"// lt",

		// load top of stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		// check
		"A=A-1",
		"D=M-D",

		// in not greater
		fmt.Sprintf("@is-lt-%v", s.Id),
		"D;JLT",
		fmt.Sprintf("@end-%v", s.Id),
		"D=0;JMP",

		// is greater
		fmt.Sprintf("(is-lt-%v)", s.Id),
		"D=-1",

		fmt.Sprintf("(end-%v)", s.Id),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

type AndStatement struct { }
func (s *AndStatement) ToAsm() []string {
	return []string{
		"// and",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D&M",
	}
}

type OrStatement struct { }
func (s *OrStatement) ToAsm() []string {
	return []string{
		"// or",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D|M",
	}
}

type NotStatement struct { }
func (s *NotStatement) ToAsm() []string {
	return []string{
		"// not",
		"@SP",
		"A=M-1",
		"M=!M",
	}
}