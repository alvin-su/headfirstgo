package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	/*var myFunc func()
	myFunc = sayHi
	myFunc()

	myFunc2 := sayHi
	myFunc2()*/

	twice(sayHi)
	twice(sayBye)
}
