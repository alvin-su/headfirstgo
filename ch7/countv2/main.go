package main

import (
	"fmt"
	"go-headFirst/ch5/datafile"
	"log"
)

func main() {
	counts := make(map[string]int)
	lines, err := datafile.GetStrings("d:\\test\\votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		//_, ok := counts[line]
		/*if !ok {
			counts[line] = 1
		} else {
			counts[line]++
		}*/
		counts[line]++
	}
	fmt.Printf("%#v\n", counts)
	for key, value := range counts {
		fmt.Printf("%s:%d\n", key, value)
	}

}
