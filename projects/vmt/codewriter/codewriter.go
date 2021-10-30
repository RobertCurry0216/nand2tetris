package codewriter

import (
	"fmt"
	"strings"
)

type CodeWriter struct {
	strings.Builder
}

func (cw *CodeWriter) Writeln(s string, args ...interface{}){
	if len(args) > 0 {
		cw.WriteString(fmt.Sprintf(s, args...))
	} else {
		cw.WriteString(s)
	}
	cw.WriteString("\n")
}

// Write takes a list of statements and builds the asembly code string.
func Write(statements []Statement) string {
	var cw CodeWriter

	for _, statement := range statements {
		statement.Compile(&cw)
	}

	return cw.String()
}