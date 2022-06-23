package main

import "fmt"

func double(number *int) { //函数形参接收一个指针类型
	*number *= 2 //更新传入的指针处的值
}

func main() {
	amount := 6
	double(&amount) //传递一个指针而不是一个变量
	fmt.Println(amount)
}
