package main

import "fmt"

func main() {
	s := "My string"
	sptr := &s
	fmt.Println(s)
	fmt.Println(sptr)
	fmt.Println(*sptr)

	sptr2 := new(string)
	fmt.Println(sptr2)
	fmt.Println(*sptr2)

	var sptr3 *string
	fmt.Println(sptr3)

	s2 := "My string"
	var sptr4 *string
	sptr4 = &s2
	fmt.Println(sptr4)
	fmt.Println(*sptr4)

	i := new(int)
	fmt.Println(i)
	fmt.Println(*i)
}
