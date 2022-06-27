package main

import (
	"fmt"
	"go-headFirst/ch14/prose"
)

func main() {
	pharese := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(pharese))

	pharese = []string{"my parents", "a rodeo clown", "a prize bull"}

	fmt.Println("A photo of", prose.JoinWithCommas(pharese))

	pharese = []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(pharese))
}
