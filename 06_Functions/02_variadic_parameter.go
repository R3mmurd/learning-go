package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xs...))
	fmt.Println(bar(xs))
}

func foo(xs ...int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}

func bar(xs []int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}
