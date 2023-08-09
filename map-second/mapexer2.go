package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{

		first: "Name",
		friends: map[string]int{
			"Jenny": 24,
			"Q":     87,
			"Ian":   47,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	for _, v := range p1.friends {
		fmt.Println(p1.first, "-friends-", v)

	}
}
