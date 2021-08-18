package main

import "fmt"

type person struct {
	first_name              string
	last_name               string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		first_name:              "Ellie",
		last_name:               "Croix",
		favoriteIceCreamFlavors: []string{"Chocolate", "Red fruits"},
	}
	p2 := person{
		first_name:              "Alex",
		last_name:               "R3mmurd",
		favoriteIceCreamFlavors: []string{"Strawberry", "Cranberry", "Cherry"},
	}

	fmt.Println("First:", p1.first_name)
	fmt.Println("Last:", p1.last_name)
	fmt.Println("Favorite ice cream flavors:")
	for i, flavor := range p1.favoriteIceCreamFlavors {
		fmt.Println("\t", i, flavor)
	}

	fmt.Println("First:", p2.first_name)
	fmt.Println("Last:", p2.last_name)
	fmt.Println("Favorite ice cream flavors:")
	for i, flavor := range p2.favoriteIceCreamFlavors {
		fmt.Println("\t", i, flavor)
	}
}
