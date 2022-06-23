package main

import (
	"fmt"
	"go-headFirst/ch5/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("d:\\test\\votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s:%d\n", name, counts[i])
	}
}
