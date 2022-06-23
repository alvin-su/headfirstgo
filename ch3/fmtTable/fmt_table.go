package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
	fmt.Println("-------------------------------------")

	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5) // 5不够2位用空格填充
}
