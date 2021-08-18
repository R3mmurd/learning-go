package main

import "fmt"

func exec_n_times(n int, f func(string)) func(string) {
	return func(name string) {
		for i := 0; i < n; i++ {
			f(name)
		}
	}
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	exec_n_times(10, greet)("R3mmurd")
}
