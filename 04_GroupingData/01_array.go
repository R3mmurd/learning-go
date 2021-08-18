package main

import "fmt"

func main() {
	xs := [5]int{10, 12, 25, 38, 42}

	for i, x := range xs {
		fmt.Println(i, x)
	}

	fmt.Printf("%T\n", xs)
}
