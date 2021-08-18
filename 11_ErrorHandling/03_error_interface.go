package main

import (
	"fmt"
)

type CustomError struct{}

func (err CustomError) Error() string {
	return "Custom Error!"
}

func foo(err error) {
	fmt.Println(err)
}

func main() {
	var err CustomError
	foo(err)
}
