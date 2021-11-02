package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"vmt/codewriter"
	"vmt/parser"
)

func main(){

	// check args
	if len(os.Args) != 2 {
		fmt.Println("Error: No file name provided")
		fmt.Println("useage: vmt <path>")
		return
	}

	path := os.Args[1]

	if isFile(path){
		if !checkExt(path) {
			fmt.Printf("Invalid file type, expected: '.vm', got: '%v'\n", filepath.Ext(path))
			return
		}
		
		translateFile(path)

	} else if isDir(path) {

		translateDir(path)

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

func translateFile(path string){
	var parser parser.Parser
	fileOutPath := replaceExt(path, ".asm")

	// translate code
	data := readFile(path)

	parser.Parse(data, "")
	code := codewriter.Write(parser.Statements)

	writeFile(fileOutPath, code)
	fmt.Println("Success!!")
	fmt.Printf("output file: %v\n", fileOutPath)
}

func translateDir(dir string) {
	var parser parser.Parser
	fileOutPath := replaceExt(dir, fmt.Sprintf("%s.asm", filepath.Base(dir)))

	// get .vm files
	files, err := filepath.Glob(filepath.Join(dir, "*.vm"))
	if err != nil {
		log.Fatal(err)
	}

	// move sys.vm to start of slice
	// found := false
	// for i, file := range files {
	// 	if strings.ToLower(filepath.Base(file)) == "sys.vm" {
	// 		found = true
	// 		temp := files[0]
	// 		files[0] = files[i]
	// 		files[i] = temp
	// 		break
	// 	}
	// }

	// if !found {
	// 	panic("sys.vm not found")
	// }

	// translate code

	for _, file := range files {
		data := readFile(file)
		parser.Parse(data, removeExt(filepath.Base(file)))
	}

	code := codewriter.Write(parser.Statements)
	writeFile(fileOutPath, code)

	fmt.Println("Success!!")
	fmt.Printf("output file: %v\n", fileOutPath)
}