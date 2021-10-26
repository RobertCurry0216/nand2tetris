package parser

import (
	"strings"
	cw "vm-translater/codewriter"
)

func Parse(bytecode string) []*cw.Statement {
	statements := make([]*cw.Statement, 1000)
	
	for _, line := range strings.Split(bytecode, "\n") {
		stmt := parseLine(line)
		statements = append(statements, &stmt)		
	}

	return statements
}

func parseLine(bytecode string) cw.Statement {
	
	return &cw.AddStatement{}
}