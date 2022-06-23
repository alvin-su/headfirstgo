package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	//var sum float64 = 0
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, number)
	}
	//sampCount := float64(len(arguments))
	fmt.Printf("Average:%0.2f\n", average(numbers...))

	/*fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))*/
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
