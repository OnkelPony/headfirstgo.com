package main

import "fmt"

// YOUR CODE HERE: Declare a "rectangle" struct type.
type rectangle struct {
	length float64
	width  float64
}

// YOUR CODE HERE: Define a rectangleInfo function.
func rectangleInfo(r *rectangle) {
	fmt.Printf("Length: %0.1f\n", r.length)
	fmt.Printf("Width: %0.1f\n", r.width)
}

func makeSquare(r *rectangle) {
	if r.width < r.length {
		r.length = r.width
	} else {
		r.width = r.length
	}
}
func main() {
	// YOUR CODE HERE: Create a new rectangle, and set its
	// length and width fields. Pass it to rectangleInfo.
	smallRect := rectangle{
		width:  2.321,
		length: 4.198,
	}
	rectangleInfo(&smallRect)
	makeSquare(&smallRect)
	rectangleInfo(&smallRect)
	bigRect := rectangle{
		length: 108,
		width:  666,
	}
	rectangleInfo(&bigRect)
	makeSquare(&bigRect)
	rectangleInfo(&bigRect)
}
