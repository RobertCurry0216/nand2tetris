package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"jack/lexer"
	"jack/parser"
	"log"
	"os"
	"path/filepath"
)

func main(){

	// check args
	if len(os.Args) != 2 {
		fmt.Println("Error: No file name provided")
		fmt.Println("useage: jack <path>")
		return
	}

	path := os.Args[1]

	if isFile(path){
		if !checkExt(path) {
			fmt.Printf("Invalid file type, expected: '.vm', got: '%v'\n", filepath.Ext(path))
			return
		}
		
		if err := translateFile(path); err != nil {
			fmt.Println("Error: translating file")
			fmt.Println(err.Error())
		}

	} else if isDir(path) {

		if err := translateDir(path); err != nil {
			fmt.Println("Error: translating file")
			fmt.Println(err.Error())
		}

	} else {	
		fmt.Printf("Error: could not find file: %v\n", path)
		return
	}
}

func removeExt(file string) string {
	ext := filepath.Ext(file)
	return file[0:len(file) - len(ext)]
}

func replaceExt(file, newExt string) string {
	return removeExt(file) + newExt
}

func checkExt(file string) bool {
	return filepath.Ext(file) == ".jack"
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


func isFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isDir(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func translateFile(path string) error {
	fileOutPath := replaceExt(path, ".xml")

	// translate code
	data := readFile(path)
	lexer := lexer.New(data)
	parser := parser.New(lexer)
	code, err := parser.ParseFile();
	if err != nil {
		return err
	}

	// write file
	writeFile(fileOutPath, code)
	fmt.Println("Success!!")
	fmt.Printf("output file: %v\n", fileOutPath)

	return nil
}

func translateDir(dir string) error {
	// get .vm files
	files, err := filepath.Glob(filepath.Join(dir, "*.jack"))
	if err != nil {
		log.Fatal(err)
	}

	// translate code

	for _, file := range files {
		if err := translateFile(file); err != nil {
			return err
		}
	}
	
	return nil
}