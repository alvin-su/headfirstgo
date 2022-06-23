package main

import "fmt"

func main() {
	fmt.Printf("%%7.3f:%7.3f\n", 12.3456) //四舍五入到三位
	fmt.Printf("%%7.2f:%7.2f\n", 12.3456) //四舍五入到两位

	fmt.Printf("%%.1f:%.1f\n", 12.3456) //四舍五入到1位 无空格填充
}
