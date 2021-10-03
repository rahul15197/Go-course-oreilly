package main

import (
	"fmt"
)

func main() {
	a := 5

	if a < 0 {
		fmt.Println("Negative!")
	} else if a < 10 {
		fmt.Println("Value is Single digit")
	} else {
		fmt.Println("Value is with multiple digits")
	}
}
