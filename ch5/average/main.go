package main

import (
	"fmt"
	"go-headFirst/ch5/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("d:\\test\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampCount := float64(len(numbers))
	fmt.Printf("Average:%0.2f\n", sum/sampCount)

}
