package main

import (
	"fmt"
	"go-headFirst/ch8/magazine"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

type park struct {
	desc  string
	count int
}
type car struct {
	name     string
	topSpeed float64
}

func applyCount(p park) {
	p.count = 200
}

func double(number *int) {
	*number *= 2
}

func main() {
	fmt.Printf("%#v\n", myStruct)

	var myCar car
	myCar.name = "Porsche 911 R"
	myCar.topSpeed = 323

	var myPark park
	myPark.desc = "Hex park"
	myPark.count = 24
	fmt.Println("Name:", myCar.name)
	fmt.Println("topSpeed:", myCar.topSpeed)

	fmt.Println("Desc:", myPark.desc)
	fmt.Println("Count:", myPark.count)

	var p park
	p.desc = "test"
	p.count = 100
	applyCount(p)
	fmt.Println(p.count)

	fmt.Println(p.count)

	amount := 6
	double(&amount)
	fmt.Println(amount)

	var pp magazine.Park
	pp.Desc = "test"
	pp.Count = 200
	magazine.ApplyCount2(&pp)
	fmt.Println(pp.Count)

	myCar2 := car{name: "car2", topSpeed: 220}
	fmt.Printf("%#v\n", myCar2)

	emp := magazine.Employee{Name: "alvinsu", Age: 30}
	emp.City = "changsha"
	fmt.Println(emp.Name)
	fmt.Println(emp.City)
}
