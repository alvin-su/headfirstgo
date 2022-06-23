package main

import (
	"fmt"
	"github.com/alvin-su/keyboard"
	"log"
)

func main() {
	fmt.Print("请输入分数：")

	grade, err := keyboard.GetFloat()
	var status string
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
