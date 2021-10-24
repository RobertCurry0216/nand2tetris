package main

import (
	"fmt"
	"vm-translater/codewriter"
	"vm-translater/parser"
)

func main(){
	fmt.Println("hello world")
	parser.Parse()
	codewriter.Write()
}