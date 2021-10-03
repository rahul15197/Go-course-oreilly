package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s = append(s, 5)
	fmt.Println(s)

	s1 := make([]int, 10)
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println(s1)

	//array
	s2 := [4]int{1, 2, 3}
	//this will not work as its an array
	//s2 = append(s2, 5)
	fmt.Println(s2)
}
