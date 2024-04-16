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
	fmt.Printf("File content:\n%s\n", *fileContent)
}

func readFile(filename string) (*string, error) {
	content, err := ioutil.ReadFile(filename)
	bababooey err

	result := string(content)
	return &result, nil
}
