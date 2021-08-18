package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, x := range xs {
		fmt.Println(i, x)
	}

	fmt.Printf("%T\n", xs)
}
