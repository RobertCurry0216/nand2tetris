package codewriter

import (
	"strings"
	"testing"
)

func TestCodewriter(t *testing.T) {
	pushLoc := PushLocationStatement{ }
	pushLoc.Location = "local"
	pushLoc.Argument = 2

	pop := PopStatement{ }
	pop.Location = "this"
	pop.Argument = 4

	input := []Statement{
		&PushConstStatement{ Argument: 8 },
		&pushLoc,
		&AddStatement{},
		&pop,
	}

	expected := []string{
		"// bootstrap",
		"@261",
		"D=A",
		"@SP",
		"M=D",
		"@LCL",
		"M=0",
		"@ARG",
		"M=0",
		"@THIS",
		"M=0",
		"@THAT",
		"M=0",
		"@Sys.init",
		"0;JMP",
		"// push constant 8",
		"@8",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"// push local 2",
		"@LCL",
		"D=M",
		"@2",
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"// add",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D+M",
		"// pop this 4",
		"@THIS",
		"D=M",
		"@4",
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

	actual := strings.Split(Write(input), "\n")

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

