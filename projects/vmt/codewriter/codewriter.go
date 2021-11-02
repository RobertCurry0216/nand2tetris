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

	// bootstrap code
	cw.Writeln("// bootstrap")
	cw.Writeln("@261")
	cw.Writeln("D=A")
	cw.Writeln("@SP")
	cw.Writeln("M=D")
	// call Sys.init
	cw.Writeln("@LCL")
	cw.Writeln("M=0")
	cw.Writeln("@ARG")
	cw.Writeln("M=0")
	cw.Writeln("@THIS")
	cw.Writeln("M=0")
	cw.Writeln("@THAT")
	cw.Writeln("M=0")
	cw.Writeln("@Sys.init")
	cw.Writeln("0;JMP")

	for _, statement := range statements {
		statement.Compile(&cw)
	}

	return cw.String()
}