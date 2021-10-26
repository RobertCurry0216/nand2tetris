package codewriter

import (
	"strings"
)

// Write takes a list of statements and builds the asembly code string.
func Write(statements []Statement) string {
	var sb strings.Builder

	for _, statement := range statements {
		for _, line := range statement.ToAsm(){
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	}

	return sb.String()
}