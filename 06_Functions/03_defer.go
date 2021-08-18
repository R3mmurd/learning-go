package main

import "fmt"

func main() {
	exec(true, 1)
	exec(false, 2)
}

func exec(err bool, id int) {
	fmt.Println("Enter", id)
	defer secure_exit(id)

	if err {
		fmt.Println("Error")
		return
	}

	fmt.Println("Execution ok!")
}

func secure_exit(id int) {
	fmt.Println("Secure exit", id)
}
