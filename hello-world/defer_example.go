package main

import "fmt"

func myNameIs() {
	defer fmt.Println("Slim Shady")
	fmt.Println("My name is")
}

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("Goodbye")
	myNameIs()
}
