package main

import (
	"fmt"
	"log"
)

// YOUR CODE HERE: Write a "perimeter" function.
func perimeter(width, length float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if length < 0 {
		return 0, fmt.Errorf("a length of %.2f is invalid", length)
	}
	return 2 * (width + length), nil
}

func main() {
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// together in the "together" variable.
	perim, err := perimeter(8.2, 10)
	if err != nil {
		log.Fatal(err)
	}
	together := perim

	perim, err = perimeter(5, 5.2)
	if err != nil {
		log.Fatal(err)
	}
	together += perim

	perim, err = perimeter(6.2, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	together += perim

	perim, err = perimeter(-10, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	together += perim

	total := fmt.Sprintf("%.1f", together)
	fmt.Println("You'll need", total, "meters of fencing")
}
