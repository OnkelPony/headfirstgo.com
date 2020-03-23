package main

import "fmt"

// YOUR CODE HERE: Write a scoreSummary function that will
// work with the calls below.

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)
}

func scoreSummary(s string, f float64, f2 float64, f3 float64) {
	average := (f + f2 + f3) / 3
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n",
		s, f, f2, f3, average)
}
