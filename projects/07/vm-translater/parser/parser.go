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

	var statement cw.Statement
	match := true

	switch words[0] {
		case "push":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }
			switch words[1] {
			case "constant":
				statement = &cw.PushConstStatement{
					Argument: arg,
				}
			default:
				statement = &cw.PushLocationStatement{
					Location: cw.Location(words[1]),
					Argument: arg,
				}
			}

		case "pop":
			arg, err := strconv.Atoi(words[2])
			if err != nil { panic(err) }

			statement = &cw.PopStatement{
				Location: cw.Location(words[1]),
				Argument: arg,
			}

		case "add":
			statement = &cw.AddStatement{}

		case "sub":
			statement = &cw.SubStatement{}

		case "neg":
			statement = &cw.NegStatement{}

		case "eq":
			statement = &cw.EqStatement{ Id: n }

		case "gt":
			statement = &cw.GtStatement{ Id: n }

		case "lt":
			statement = &cw.LtStatement{ Id: n }

		case "and":
			statement = &cw.AndStatement{}

		case "or":
			statement = &cw.OrStatement{}

		case "not":
			statement = &cw.NotStatement{}
		
		default: 
			match = false
	}


	return statement, match
}