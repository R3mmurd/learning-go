package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Gray",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(s.doors)
	fmt.Println(t.color)
}
