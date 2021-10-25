package codewriter

import (
	"testing"
)

func TestPushConstStatement(t *testing.T) {
	s := PushConstStatement{ 4 }
	actual := s.ToAsm()
	expected := []string{
		"// push constant 4",
		"@4",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestPushLocationStatement(t *testing.T) {
	s := PushLocationStatement{ "LCL",  4 }
	actual := s.ToAsm()
	expected := []string{
		"// push LCL 4",
		"@LCL",
		"D=M",
		"@4",
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestPopStatement(t *testing.T) {
	s := PopStatement{ "LCL",  4 }
	actual := s.ToAsm()
	expected := []string{
		"// pop LCL 4",
		"@LCL",
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

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestAddStatement(t *testing.T) {
	s := AddStatement{}
	actual := s.ToAsm()
	expected := []string{
		"// add",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D+M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestSubStatement(t *testing.T) {
	s := SubStatement{}
	actual := s.ToAsm()
	expected := []string{
		"// sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D-M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestNegStatement(t *testing.T) {
	s := NegStatement{}
	actual := s.ToAsm()
	expected := []string{
		"// neg",
		"@SP",
		"A=M-1",
		"M=-M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestEqStatement(t *testing.T) {
	s := EqStatement{ 1234 }
	actual := s.ToAsm()
	expected := []string{
		"// eq",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		"@is-eq-1234",
		"D;JEQ",
		"@end-1234",
		"D=0;JMP",
		"(is-eq-1234)",
		"D=-1",
		"(end-1234)",
		"@SP",
		"A=M-1",
		"M=D",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestGtStatement(t *testing.T) {
	s := GtStatement{ 1234 }
	actual := s.ToAsm()
	expected := []string{
		"// gt",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		"@is-gt-1234",
		"D;JGT",
		"@end-1234",
		"D=0;JMP",
		"(is-gt-1234)",
		"D=-1",
		"(end-1234)",
		"@SP",
		"A=M-1",
		"M=D",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLtStatement(t *testing.T) {
	s := LtStatement{ 1234 }
	actual := s.ToAsm()
	expected := []string{
		"// lt",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		"@is-lt-1234",
		"D;JLT",
		"@end-1234",
		"D=0;JMP",
		"(is-lt-1234)",
		"D=-1",
		"(end-1234)",
		"@SP",
		"A=M-1",
		"M=D",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestAndStatement(t *testing.T) {
	s := AndStatement{ }
	actual := s.ToAsm()
	expected := []string{
		"// and",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D&M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestOrStatement(t *testing.T) {
	s := OrStatement{ }
	actual := s.ToAsm()
	expected := []string{
		"// or",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D|M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestNotStatement(t *testing.T) {
	s := NotStatement{ }
	actual := s.ToAsm()
	expected := []string{
		"// not",
		"@SP",
		"A=M-1",
		"M=!M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}