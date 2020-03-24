package main

import (
	"headfirstgo.com/source/head-first-go/ch04/go/src/github.com/headfirstgo/greeting"
	"headfirstgo.com/source/head-first-go/ch04/go/src/github.com/headfirstgo/greeting/dansk"
	"headfirstgo.com/source/head-first-go/ch04/go/src/github.com/headfirstgo/greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
	dansk.GodMorgen()
}
