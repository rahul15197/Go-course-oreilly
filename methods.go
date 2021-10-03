package main

import (
	"fmt"
	"time"
)

type Person struct {
	First string
	Last  string
	Dob   time.Time
}

func NewPerson(first, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last:  last,
		Dob:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.First)
}

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}
func getHello(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}

func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365)
}
func main() {
	sayHello("Hank")
	s := getHello("Rahul")
	fmt.Print(s)

	p := &Person{First: "John", Last: "Doe"}
	p.SayHello()

	p1 := NewPerson("Jane", "Doe", 1998, 1, 1)
	p1.SayHello()
	fmt.Println(p1.GetAge(), "years old")
}
