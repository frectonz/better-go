package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := readFile("try.go")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", fileContent)
}

func readFile(filename string) (string, error) {
	try c or string = ioutil.ReadFile(filename)
	return string(c), nil
}
