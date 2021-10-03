package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
	Phone Phone
}
type Phone struct {
	AreaCode string
	Prefix   string
	Suffix   string
}

func main() {
	p := Person{
		First: "John",
		Last:  "Doe",
		Age:   30,
		Phone: Phone{
			AreaCode: "123",
			Prefix:   "345",
			Suffix:   "0123",
		},
	}
	fmt.Println(p)
}
