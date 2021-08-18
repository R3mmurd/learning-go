package main

import "fmt"

func main() {
	ss1 := []string{"James", "Bond", "Shaken, not stirred"}
	ss2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	sss := [][]string{ss1, ss2}

	for i, ss := range sss {
		fmt.Println("Record:", i)

		for j, s := range ss {
			fmt.Printf("Value at %d is %q\n", j, s)
		}
	}
}
