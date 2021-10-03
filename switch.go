package main

import (
	"fmt"
)

func main() {
	a := 20

	switch a {
	case 10:
		fmt.Println("It is 10")
	case 1:
		fmt.Println("It is 1")
	default:
		fmt.Println("Neither 1 nor 10")
	}
}
