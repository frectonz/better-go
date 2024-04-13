package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := readFile("/etc/locale.conf")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", fileContent)
}

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return make([]string, 1)[0], err
	}
	return string(content), nil
}
