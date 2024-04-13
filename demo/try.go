package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := readFile("sample.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", fileContent)
}

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	try err or string

	return string(content), nil
}
