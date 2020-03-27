package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("go/bin/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// YOUR CODE HERE: Build a slice containing all the
	// key strings from the "counts" map.
	// Call the sort.Strings method on your slice.
	// Loop through the names in the sorted slice, and
	// print the name and the corresponding count from
	// the map.
	names := make([]string, len(counts))
	i := 0
	for name := range counts {
		names[i] = name
		i++
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}
