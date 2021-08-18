package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	ys := append(xs[:3], xs[6:]...)
	fmt.Println(ys)
}
