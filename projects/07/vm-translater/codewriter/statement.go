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
	ToAsm() []string
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

func (s *PushLocationStatement) ToAsm() []string {
	base := getBaseLocation(s.Location)
	return []string{
		fmt.Sprintf("// push %v %v", s.Location, s.Argument),
		s.Location.Address(),
		fmt.Sprintf("D=%s", base),
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
	LocationStatement
}

func (s *PopStatement) String() string {
	return fmt.Sprintf("< pop %v %d >", s.Location, s.Argument)
}

func (s *PopStatement) ToAsm() []string {
	base := getBaseLocation(s.Location)

	return []string{
		fmt.Sprintf("// pop %v %v", s.Location, s.Argument),

		// set A to location + arg
		s.Location.Address(),
		fmt.Sprintf("D=%s", base),
		fmt.Sprintf("@%d", s.Argument),
		"D=D+A",

		// store addr in R15
		"@R15",
		"M=D",

		// dec stack pointer
		"@SP",
		"AM=M-1",

		// load value
		"D=M",
		"@R15",
		"A=M",
		"M=D",
	}
}

type AddStatement struct {}
func (s *AddStatement) String() string { return "< add >"}
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
func (s *SubStatement) String() string { return "< sub >"}
func (s *SubStatement) ToAsm() []string {
	return []string{
		"// sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=M-D",
	}
}

type NegStatement struct {}
func (s *NegStatement) String() string { return "< neg >"}
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
func (s *EqStatement) String() string { return fmt.Sprintf("< Eq %d >", s.Id) }
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
func (s *GtStatement) String() string { return fmt.Sprintf("< gt %d >", s.Id)}
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
func (s *LtStatement) String() string { return fmt.Sprintf("< lt %d >", s.Id)}
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
func (s *AndStatement) String() string { return "< and >"}
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
func (s *OrStatement) String() string { return "< or >"}
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
func (s *NotStatement) String() string { return "< not >"}
func (s *NotStatement) ToAsm() []string {
	return []string{
		"// not",
		"@SP",
		"A=M-1",
		"M=!M",
	}
}