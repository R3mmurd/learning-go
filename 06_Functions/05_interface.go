package main

import "fmt"

const π = 3.1415926535897932384626433

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (c Circle) Area() float64 {
	return π * c.radius * c.radius
}

func info(s Shape) {
	switch s.(type) {
	case Square:
		fmt.Print("Square ")
	case Circle:
		fmt.Print("Circle ")
	}
	fmt.Println("area:", s.Area())
}

func main() {
	c := Circle{
		radius: 3,
	}
	s := Square{
		length: 2,
	}
	info(c)
	info(s)
}
