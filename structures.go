package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p := Person{
		First: "John",
		Last:  "Doe",
		Age:   30,
	}
	q := &Person{"Jane", "Doe", 25}
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(p.Age)
	fmt.Println(q.Age)
}
