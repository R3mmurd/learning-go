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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, p := range m {
		fmt.Println("First:", p.first_name)
		fmt.Println("Last:", p.last_name)
		fmt.Println("Favorite ice cream flavors:")
		for i, flavor := range p.favoriteIceCreamFlavors {
			fmt.Println("\t", i, flavor)
		}
	}
}
