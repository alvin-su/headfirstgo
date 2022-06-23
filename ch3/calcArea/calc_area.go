package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0 //返回area 而不是打印它。
}

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total:%.2f liters\n", total)
}
