package main

import "fmt"

func main() {
	var x interface{}
	x = "Hello, World"
	s, ok := x.(string)
	if !ok {
		panic("Nooooo!")
	}
	fmt.Println(s)
	s1, ok := x.(int)
	if !ok {
		panic("Nooooo!")
	}
	fmt.Println(s1)
}
