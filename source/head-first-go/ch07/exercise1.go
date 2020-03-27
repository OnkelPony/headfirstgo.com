package main

import "fmt"

func main() {
	// We'll count the number of times each letter occurs
	// within this slice.
	letters := []string{"b", "a", "c", "a", "b", "a",
		"a", "b", "c", "i", "s", "l", "a", "m", "i",
		"s", "b", "u", "l", "l", "s", "h", "i", "t"}
	// YOUR CODE HERE: Process each element of "letters",
	// keeping track of how many times you've seen "a",
	// "b", or "c".
	occurs := make(map[string]int)
	for _, letter := range letters {
		occurs[letter]++
	}
	fmt.Println(occurs)
	// Finally, print out the number of times each letter
	// occurred.
}
