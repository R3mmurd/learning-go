package main

import (
	"fmt"

	"github.com/learning-go/13_TestingAndBenchmarking/01_dog/dog"
)

// Canine is a data type to represent dogs.
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
	fmt.Println(dog.YearsTwo(20))
}
