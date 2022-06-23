package main

import (
	"fmt"
	"go-headFirst/ch6/datafilev2"
	"log"
)

func main() {
	numbers, err := datafilev2.GetFloata("d:\\test\\data.txt")
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
