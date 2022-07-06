package main

import (
	"fmt"

	sort "github.com/tlgevers/golang-projects/interface-example/pkg/sort"
)

func main() {
	var tracks = []**sort.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, sort.length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	fmt.Println("Hello, playground")
}
