package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))

	var myIntPointer *int //声明一个指向int类型的指针
	myIntPointer = &myInt //给指针变量分配一个指针
	fmt.Println(&myInt)
	fmt.Println(myIntPointer)

	myInt2 := 4
	myIntPointer2 := &myInt2
	fmt.Println(myIntPointer2)  //打印指针本身
	fmt.Println(*myIntPointer2) //打印指针处的值

	*myIntPointer2 = 8
	fmt.Println(*myIntPointer2)
	fmt.Println(myInt2)

	var myFloatPointer *float64 = createPointer() //将函数返回的指针赋给一个变量。
	fmt.Println(*myFloatPointer)                  //打印指针处的值

	var myBool bool = true
	printPointer(&myBool)
}
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) { //为该参数使用一个指针类型
	fmt.Println(*myBoolPointer) //打印传入的指针处的值
}
