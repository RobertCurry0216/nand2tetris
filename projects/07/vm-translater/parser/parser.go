package parser

import (
	"strconv"
	"strings"
	cw "vm-translater/codewriter"
)

func Parse(bytecode string) []cw.Statement {
	statements := make([]cw.Statement, 0, 1000)
	
	for i, line := range strings.Split(bytecode, "\n") {
		stmt, match := parseLine(line, i)
		if match {
			statements = append(statements, stmt)		
		}
	}

	return statements
}

func parseLine(bytecode string, n int) (cw.Statement, bool) {
	code := strings.Split(bytecode, "//")[0]
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return nil, false
	}

	words := strings.Split(code, " ")

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
			default:
				statement := &cw.PushLocationStatement{}
				statement.Location = cw.Location(words[1])
				statement.Argument = arg
				return statement, true
			}

		case "pop":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }

			statement := &cw.PopStatement{}
			statement.Location = cw.Location(words[1])
			statement.Argument = arg
			return statement, true

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
		
		default: 
			return nil, false

	}
}