package main

import "fmt"

type Appliancette interface {
	TurnOn()
}
type EleFan string

func (f EleFan) TurnOn() {
	fmt.Println("Spinning")
}

type SmallCoffeePot string

func (c SmallCoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c SmallCoffeePot) Brew() {
	fmt.Println("Heating Up")
}
func main() {
	var device Appliancette
	device = EleFan("Windco Breeze")
	device.TurnOn()
	device = SmallCoffeePot("LuxBrew")
	device.TurnOn()

}
