package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xs = append(xs, 52)
	fmt.Println(xs)
	xs = append(xs, 53, 54, 55)
	fmt.Println(xs)
	ys := []int{56, 57, 58, 59, 60}
	xs = append(xs, ys...)
	fmt.Println(xs)
}
