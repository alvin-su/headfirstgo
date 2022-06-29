package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*file, err := os.OpenFile("test.txt", os.O_RDONLY, os.FileMode(0600))
	check(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())*/

	/*file, err := os.OpenFile("test.txt", os.O_WRONLY, os.FileMode(0600))
	check(err)
	myString := []byte("amazing!\n")
	//myString := "amazing222!\n"
	//_, err = fmt.Fprintln(file, myString)
	_, err = file.Write(myString)
	check(err)
	err = file.Close()
	check(err)*/

	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))
}
