package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 43, 98, 56, 11, 8, 0}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr.No"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
