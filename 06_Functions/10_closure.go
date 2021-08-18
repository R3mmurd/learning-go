package main

import "fmt"

func build_fibonacci_generator() func() int {
	a := 0
	b := 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	fibonacci_generator_1 := build_fibonacci_generator()

	fmt.Println("Fibonacci generator 1")

	for i := 0; i < 5; i++ {
		fmt.Println(fibonacci_generator_1())
	}

	fibonacci_generator_2 := build_fibonacci_generator()

	fmt.Println("Fibonacci generator 2")

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci_generator_2())
	}

	fmt.Println("Fibonacci generator 1 again")

	for i := 0; i < 2; i++ {
		fmt.Println(fibonacci_generator_1())
	}
}
