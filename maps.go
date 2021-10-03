package main

import "fmt"

func main() {
	m := map[string]string{}
	m["first"] = "John"
	m["last"] = "Doe"
	fmt.Println(m)

	m1 := map[string]string{
		"first": "Jane",
		"last":  "Doe",
	}
	fmt.Println(m1)

	m2 := make(map[string]string)
	m2["first"] = "Rahul"
	m2["last"] = "Maheshwari"
	fmt.Println(m2["first"])
}
