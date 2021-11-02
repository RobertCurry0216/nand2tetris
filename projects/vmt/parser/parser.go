package parser

import (
	"fmt"
	"strconv"
	"strings"
	cw "vmt/codewriter"
)

type Parser struct {
	id int
	Statements []cw.Statement
	Function string
}

func (p *Parser) Parse(bytecode string, file string) {
	for _, line := range strings.Split(bytecode, "\n") {
		stmt, ok := p.parseLine(line, p.id, file)
		p.id++

		if !ok {
			panic(fmt.Sprintf("error parsing line: %s", line))
		}

		if stmt != nil {
			p.Statements = append(p.Statements, stmt)		
		}
	}
}

func (p *Parser) parseLine(bytecode string, n int, file string) (cw.Statement, bool) {
	code := strings.Split(bytecode, "//")[0]
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return nil, true
	}

	words := strings.Fields(code)

	switch words[0] {
		case "push":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }
			switch words[1] {
			case "constant":
				statement := &cw.PushConstStatement{
					Argument: arg,
				}
				return statement, true
			case "static":
				statement := &cw.PushStaticStatement{
					File: file,
					Argument: arg,
				}
				return statement, true
			default:
				statement := &cw.PushLocationStatement{}
				statement.Location = cw.Location(words[1])
				statement.Argument = arg
				return statement, true
			}

		case "pop":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }

			switch words[1] {
			case "static":
				statement := &cw.PopStaticStatement{
					File: file,
					Argument: arg,
				}
				return statement, true
			default:
				statement := &cw.PopStatement{}
				statement.Location = cw.Location(words[1])
				statement.Argument = arg
				return statement, true
			}
		case "add":
			statement := &cw.AddStatement{}
			return statement, true

		case "sub":
			statement := &cw.SubStatement{}
			return statement, true

		case "neg":
			statement := &cw.NegStatement{}
			return statement, true
			
		case "eq":
			statement := &cw.EqStatement{ Id: n }
			return statement, true
			
		case "gt":
			statement := &cw.GtStatement{ Id: n }
			return statement, true

		case "lt":
			statement := &cw.LtStatement{ Id: n }
			return statement, true

		case "and":
			statement := &cw.AndStatement{}
			return statement, true

		case "or":
			statement := &cw.OrStatement{}
			return statement, true

		case "not":
			statement := &cw.NotStatement{}
			return statement, true

		case "label":
			statement := &cw.LabelStatement{ Name: words[1], Function: p.Function }
			return statement, true

		case "goto":
			statement := &cw.GotoStatement{ Name: words[1], Function: p.Function }
			return statement, true

		case "if-goto":
			statement := &cw.IfGotoStatement{ Name: words[1], Id: n, Function: p.Function}
			return statement, true

		case "function":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }
			statement := &cw.FunctionStatement{ Name: words[1], Nvars: arg }
			p.Function = words[1]
			return statement, true

		case "call":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }
			statement := &cw.CallStatement{ Name: words[1], Nargs: arg, Id: n }
			return statement, true

		case "return":
			statement := &cw.ReturnStatement{}
			return statement, true
		
		default: 
			return nil, false

	}
}