package main

import "fmt"

func main() {
	year := 1982
	for {
		fmt.Println(year)

		if year == 2021 {
			break
		}

		year++
	}
}
