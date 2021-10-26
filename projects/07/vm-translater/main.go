package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"vm-translater/codewriter"
	"vm-translater/parser"
)

func main(){

	// check args
	if len(os.Args) != 2 {
		fmt.Println("Error: No file name provided")
		fmt.Println("useage: vm-translater <filepath>")
		return
	}

	filePath := os.Args[1]

	if !fileExists(filePath){
		fmt.Printf("Error: could not find file: %v\n", filePath)
		return
	}

	if !checkExt(filePath) {
		fmt.Printf("Invalid file type, expected: '.vm', got: '%v'\n", filepath.Ext(filePath))
		return
	}

	fileOutPath := replaceExt(filePath, ".asm")

	if fileExists(fileOutPath){
		fmt.Printf("Error: output file already exists: %v\n", fileOutPath)
		return
	}


	// translate code
	data := readFile(filePath)

	statements := parser.Parse(data)
	code := codewriter.Write(statements)

	writeFile(fileOutPath, code)
	fmt.Println("Success!!")
	fmt.Printf("output file: %v\n", fileOutPath)
}

func replaceExt(file, newExt string) string {
	ext := filepath.Ext(file)
	return file[0:len(file) - len(ext)] + newExt
}

func checkExt(file string) bool {
	return filepath.Ext(file) == ".vm"
}


func readFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}


func writeFile(filePath, data string) {
	message := []byte(data)
	err := ioutil.WriteFile(filePath, message, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}


func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}