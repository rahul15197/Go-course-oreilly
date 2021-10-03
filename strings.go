package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I'm a string variable holding this sentence"
	fmt.Printf("Ends with sentence? %v\n", strings.HasSuffix(s, "sentence"))
}
