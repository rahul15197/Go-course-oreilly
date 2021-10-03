package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("output1.txt", []byte("Hello, World"), 0644)
	if err != nil {
		panic("Unable to write file")
	}
	buf, err := ioutil.ReadFile("output1.txt")
	if err != nil {
		panic("Unable to read")
	}
	fmt.Println(string(buf))
}
