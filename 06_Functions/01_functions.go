package main

import "fmt"

func main() {
	n1 := foo()
	n2, s1 := bar()

	fmt.Println(n1)
	fmt.Println(n2, s1)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 25, "Hello"
}
