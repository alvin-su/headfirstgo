package main

import "fmt"

func one() {
	defer fmt.Println("deferred in one()")
	two()
}

func two() {
	defer fmt.Println("deferred in two()")
	panic("Let's see what's been deferred!")
	//three()
}

func three() {
	panic("This call stack's too deep for me!")
}

func main() {
	one()
}
