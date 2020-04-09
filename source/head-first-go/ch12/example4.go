package main

import "fmt"

func otherFunction() {
	fmt.Println("c")
	fmt.Println(recover())
}

func myFunctionTres() {
	defer otherFunction()
	panic("oh no")
	fmt.Println("d")
}

func main() {
	fmt.Println("a")
	myFunctionTres()
	fmt.Println("b")
}
