package main

import "fmt"

func myFunctionDos() {
	panic("oh no")
	fmt.Println(recover())
}

func main() {
	fmt.Println("a")
	myFunctionDos()
	fmt.Println("b")
}
