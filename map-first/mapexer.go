package main

import (
	"fmt"
)

type engine struct {
	electric bool
}
type vehicle struct {
	engine
	make   string
	model  string
	doors  int
	colour string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:   "Honda",
		model:  "Tundra",
		doors:  2,
		colour: "red",
	}
	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:   "BMW",
		model:  "series",
		doors:  4,
		colour: "black",
	}
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.make, v1.model)
	fmt.Println(v2.electric, v2.make, v2.model)
}
