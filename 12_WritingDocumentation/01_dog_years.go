package main

import (
	"fmt"

	"github.com/learning-go/12_WritingDocumentation/dog"
)

// Canine defines a data type to represent dogs.
type Canine struct {
	name string
	age  int
}

func main() {
	maggie := Canine{
		name: "Maggie",
		age:  dog.Years(7),
	}
	arya := Canine{
		name: "Arya",
		age:  dog.Years(2),
	}

	fmt.Println(maggie)
	fmt.Println(arya)
}
