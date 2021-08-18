package main

import "fmt"

func main() {
	x := 10
	y := 20

	b1 := x == y
	b2 := x <= y
	b3 := x >= y
	b4 := x != y
	b5 := x < y
	b6 := x > y

	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", b1, b2, b3, b4, b5, b6)

}
