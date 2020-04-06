package main

import "fmt"

type Whistlet string

func (w Whistlet) MakeSound() {
	fmt.Println("Tweet!")
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistlet)
	if ok {
		whistle.MakeSound()
	}
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistlet("Toyco Canary"))
}
