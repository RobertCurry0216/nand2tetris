package codewriter

import (
	"strings"
	"testing"
)

func TestCodewriter(t *testing.T) {
	input := []Statement{
		&PushConstStatement{ 8 },
		&PushLocationStatement{ "LCL", 2 },
		&AddStatement{},
		&PopStatement{ "THIS", 4 },
	}

	expected := []string{
		"// push constant 8",
		"@8",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
		"// push LCL 2",
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
		"// pop THIS 4",
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

