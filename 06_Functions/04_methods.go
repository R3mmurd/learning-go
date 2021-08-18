package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %s %s and I'm %d years old!\n", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Alex",
		last:  "Remmurd",
		age:   39,
	}
	p.speak()
}
