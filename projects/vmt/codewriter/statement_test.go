package codewriter

import (
	"strings"
	"testing"
)

func TestPushConstStatement(t *testing.T) {
	s := PushConstStatement{ 4 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestPushLocationStatement(t *testing.T) {
	s := PushLocationStatement{}
	s.Location = "local"
	s.Argument = 4
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// push local 4",
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestPopStatement(t *testing.T) {
	s := PopStatement{ }
	s.Location = "local"
	s.Argument = 4
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// pop local 4",
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestAddStatement(t *testing.T) {
	s := AddStatement{}
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestSubStatement(t *testing.T) {
	s := SubStatement{}
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=M-D",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestNegStatement(t *testing.T) {
	s := NegStatement{}
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// neg",
		"@SP",
		"A=M-1",
		"M=-M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestEqStatement(t *testing.T) {
	s := EqStatement{ 1234 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestGtStatement(t *testing.T) {
	s := GtStatement{ 1234 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLtStatement(t *testing.T) {
	s := LtStatement{ 1234 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestAndStatement(t *testing.T) {
	s := AndStatement{ }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestOrStatement(t *testing.T) {
	s := OrStatement{ }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
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
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestNotStatement(t *testing.T) {
	s := NotStatement{ }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// not",
		"@SP",
		"A=M-1",
		"M=!M",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLabelStatement(t *testing.T) {
	s := LabelStatement{ Name: "test" }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// label test",
		"(test)",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestGotoStatement(t *testing.T) {
	s := GotoStatement{ Name: "test" }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// goto test",
		"@test",
		"0;JMP",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestIfGotoStatement(t *testing.T) {
	s := IfGotoStatement{ Name: "test", Id: 2 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// if-goto test",
		"@SP",
		"AM=M-1",
		"D=M",
		"@test-FALSE-2",
		"D;JEQ",
		"@test",
		"0;JMP",
		"(test-FALSE-2)",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestFunctionStatement(t *testing.T) {
	s := FunctionStatement{ Name: "f.test", Nvars: 2 }
	var cw CodeWriter
	s.Compile(&cw)
	actual := strings.Split(strings.TrimSpace(cw.String()), "\n")
	expected := []string{
		"// function f.test 2",
		"(f.test)",
		"@SP",
		"D=A",
		"@LCL",
		"AM=D",
		"M=0",
		"A=A+1",
		"M=0",
		"A=A+1",
		"D=A",
		"@SP",
		"M=D",
	}

	if len(actual) != len(expected){
		t.Errorf("line count mismatch, expected: %v, got: %v", len(expected), len(actual))
		t.FailNow()
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}