package main

import "fmt"

func main() {

	favorite_things := map[string][]string{
		`bond_james`:       {`Shaken, not stirred`, `Martinis`, `Women`},
		`monenypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:            {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for name, things := range favorite_things {
		fmt.Println("This record is for:", name)

		for i, thing := range things {
			fmt.Println("\t", i, thing)
		}
	}
}
