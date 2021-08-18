package main

import "fmt"

const (
	year0 = 2021 + iota
	year1
	year2
	year3
)

func main() {
	fmt.Println(year0, year1, year2, year3)
}
