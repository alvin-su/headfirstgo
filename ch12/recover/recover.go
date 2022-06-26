package main

import "fmt"

func calmDown() {
	//recover()
	//fmt.Println(recover())
	p := recover()
	myErr, ok := p.(error) //类型断言
	if ok {
		fmt.Println(myErr.Error())
	}

}

func freakOut() {
	defer calmDown()
	panic("oh no")
	fmt.Println("I'won't be run!")
}

func main() {
	/*freakOut()
	fmt.Println("Exiting normally")*/
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)

}
