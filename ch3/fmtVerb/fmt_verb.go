package main

import "fmt"

func main() {
	fmt.Printf("A float:%f\n", 3.1415)                //格式化浮点数
	fmt.Printf("An integer:%d\n", 15)                 //格式化整数
	fmt.Printf("A string:%s\n", "hello")              //格式化字符串
	fmt.Printf("A bollen:%t\n", false)                //格式化布尔值
	fmt.Printf("Values:%v %v %v\n", 1.2, "\t", true)  //格式化任何值，根据提供的值的类型
	fmt.Printf("Values:%#v %#v %#v", 1.2, "\t", true) //格式化任何值，根据在Go语言中显示的格式进行格式化
	fmt.Printf("Types:%T %T %T\n", 1.2, "\t", true)   //格式化所提供的值得类型 int、string等
	fmt.Printf("Percent sign:%%\n")                   //格式化一个完完全全的%号

}
