package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (pp *Person) Speak() {
	fmt.Printf("Hello! My name is %s %s and I'm %d years old\n", pp.First, pp.Last, pp.Age)
}

type Human interface {
	Speak()
}

func saySomething(h Human) {
	h.Speak()
}

func main() {
	p := Person{
		First: "Jhon",
		Last:  "Doe",
		Age:   25,
	}
	/*
		cannot use p (variable of type Person) as Human
		value in argument to saySomething: missing
		method Speak (Speak has pointer receiver)
	*/
	// saySomething(p)
	saySomething(&p)
}
