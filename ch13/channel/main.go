package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(chanel chan string) {
	chanel <- "a"
	chanel <- "b"
	chanel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	/*myChannel := make(chan string)
	go greeting(myChannel)
	receiveValue := <-myChannel
	fmt.Println(receiveValue)*/

	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Println(<-channel1) //a
	fmt.Println(<-channel2) //d
	fmt.Println(<-channel1) //b
	fmt.Println(<-channel2) //e
	fmt.Println(<-channel2) //f
	fmt.Println(<-channel1) //c

}
