package main

import "fmt"

func make_greet() func(string) {
	return func(name string) {
		fmt.Println("Hello,", name)
	}
}

func main() {
	f := make_greet()
	f("R3mmurd")
}
