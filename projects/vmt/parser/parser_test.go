package parser

import (
	"testing"
	cw "vmt/codewriter"
)

func expectEq(t *testing.T, a, b interface{}){
	if a != b {
		t.Errorf("expected %v, got %v", a, b)
		t.FailNow()
	}
}

func TestParseLine_ExtraSpaces(t *testing.T){
	line := " push  local  3 "
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.PushLocationStatement); !ok {
		t.Errorf("expected: PushLocationStatement, got: %T", stmt)
	}
}

func TestParseLine_Comments(t *testing.T){
	line := "push local 3 // hello world"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.PushLocationStatement); !ok {
		t.Errorf("expected: PushLocationStatement, got: %T", stmt)
	}
}

func TestParseLine_Fail(t *testing.T){
	line := "pish local 3"
	var p Parser
	_, parsed := p.parseLine(line, 1, "")

	if parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}
}

func TestParseLine_Push(t *testing.T){
	line := "push local 3"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.PushLocationStatement); ok {
		expectEq(t, stmt.Location, cw.Location("local"))
		expectEq(t, stmt.Argument, 3)
	} else {
		t.Errorf("expected: PushLocationStatement, got: %T", stmt)
	}
}

func TestParseLine_PushConst(t *testing.T){
	line := "push constant 3"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.PushConstStatement); ok {
		expectEq(t, stmt.Argument, 3)
	} else {
		t.Errorf("expected: PushConstStatement, got: %T", stmt)
	}
}

func TestParseLine_Pop(t *testing.T){
	line := "pop local 3"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.PopStatement); ok {
		expectEq(t, stmt.Location, cw.Location("local"))
		expectEq(t, stmt.Argument, 3)
	} else {
		t.Errorf("expected: PopStatement, got: %T", stmt)
	}
}

func TestParseLine_Add(t *testing.T){
	line := "add"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.AddStatement); !ok {
		t.Errorf("expected: AddStatement, got: %T", stmt)
	}
}

func TestParseLine_Sub(t *testing.T){
	line := "sub"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.SubStatement); !ok {
		t.Errorf("expected: SubStatement, got: %T", stmt)
	}
}

func TestParseLine_Neg(t *testing.T){
	line := "neg"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.NegStatement); !ok {
		t.Errorf("expected: NegStatement, got: %T", stmt)
	}
}

func TestParseLine_Eq(t *testing.T){
	line := "eq"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.EqStatement); ok {
		expectEq(t, stmt.Id, 1)
	} else {
		t.Errorf("expected: EqStatement, got: %T", stmt)
	}
}

func TestParseLine_Gt(t *testing.T){
	line := "gt"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.GtStatement); ok {
		expectEq(t, stmt.Id, 1)
	} else {
		t.Errorf("expected: GtStatement, got: %T", stmt)
	}
}

func TestParseLine_Lt(t *testing.T){
	line := "lt"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.LtStatement); ok {
		expectEq(t, stmt.Id, 1)
	} else {
		t.Errorf("expected: LtStatement, got: %T", stmt)
	}
}

func TestParseLine_And(t *testing.T){
	line := "and"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.AndStatement); !ok {
		t.Errorf("expected: AndStatement, got: %T", stmt)
	}
}

func TestParseLine_Or(t *testing.T){
	line := "or"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.OrStatement); !ok {
		t.Errorf("expected: OrStatement, got: %T", stmt)
	}
}

func TestParseLine_Not(t *testing.T){
	line := "not"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.NotStatement); !ok {
		t.Errorf("expected: NotStatement, got: %T", stmt)
	}
}

func TestParseLine_Label(t *testing.T){
	line := "label hello-world"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.LabelStatement); ok {
		expectEq(t, stmt.Name, "hello-world")
	} else {
		t.Errorf("expected: LabelStatement, got: %T", stmt)
	}
}

func TestParseLine_Goto(t *testing.T){
	line := "goto hello-world"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.GotoStatement); ok {
		expectEq(t, stmt.Name, "hello-world")
	} else {
		t.Errorf("expected: GotoStatement, got: %T", stmt)
	}
}

func TestParseLine_IfGoto(t *testing.T){
	line := "if-goto hello-world"
	var p Parser
	s, parsed := p.parseLine(line, 1, "")

	if !parsed {
		t.Errorf("failed to parse line: %s", line)
		t.FailNow()
	}

	if stmt, ok := s.(*cw.IfGotoStatement); ok {
		expectEq(t, stmt.Name, "hello-world")
		expectEq(t, stmt.Id, 1)
	} else {
		t.Errorf("expected: IfGotoStatement, got: %T", stmt)
	}
}