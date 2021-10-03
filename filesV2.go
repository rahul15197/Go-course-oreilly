package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("output.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer f.Close()
	buf := make([]byte, 64)
	cnt, err := f.Read(buf)
	if err != nil {
		panic("Unable to read")
	}
	fmt.Printf("Read %d bytes\n", cnt)
	fmt.Println(string(buf[:cnt]))
}
