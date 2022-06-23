package main

import (
	"fmt"
	"go-headFirst/ch10/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	date.SetYear(2002)
	date.SetMonth(6)
	date.SetDay(23)
	fmt.Println(date)
	fmt.Println(date.Year())

	event := calendar.Event{}
	err := event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(6)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}
	event.SetTitle("今天的日期")
	fmt.Println(event.Title())
	fmt.Println(event.Date)
}
