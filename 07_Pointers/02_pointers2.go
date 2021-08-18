package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func changeMe(p *Person) {
	p.first = "Jane"
	p.age = 27
}

func main() {
	p := Person{
		first: "Jhon",
		last:  "Doe",
		age:   25,
	}
	fmt.Println(p)
	changeMe(&p)

	fmt.Println(p)
}
