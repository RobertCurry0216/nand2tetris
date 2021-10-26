package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	fmt.Println("hello world")

	filePath := "./test.txt"

	data := readFile(filePath)
	writeFile("./out.txt", data)

	fmt.Println(data)
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