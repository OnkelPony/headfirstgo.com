// average calculates the average of several numbers.
package main

import (
	"fmt"
	"headfirstgo.com/source/head-first-go/ch05/go/src/github.com/headfirstgo/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("headfirstgo.com/source/head-first-go/ch05/go/bin/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
