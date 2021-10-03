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
	pt := struct {
		x int
		y int
	}{
		x: 10,
		y: 20,
	}
	fmt.Println(pt)
}
