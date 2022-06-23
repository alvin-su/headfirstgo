package main

import "fmt"

func main() {
	var mySlice []string
	mySlice = make([]string, 7)
	mySlice[0] = "do"
	fmt.Println(mySlice[0])

	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	slice2 := []string{"a", "b"}
	fmt.Println(slice2)
	slice2 = append(slice2, "c")
	fmt.Println(slice2, len(slice2))
}
