package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	// error handling of "early return" style (idiomatic style in Go)
	if err != nil {
		return "", err
	}
	if len(dat) == 0 {
		//return "", errors.New("Empty content")
		return "", fmt.Errorf("Empty content (filename=%v)", filename)
	}
	// code if no error
	return string(dat), nil // err = nil means err has not been defined
}

func main() {
	dat, err := readFile("test.txt")
	// "early return"
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	// code if no error
	fmt.Println("File content:")
	fmt.Println(dat)
}
