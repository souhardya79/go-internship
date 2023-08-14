package main

import (
	"fmt"
)

func slices(slice1 []string, slice2 []string) ([]string, []string, []string) {
	unique := make(map[string]bool)
	var onlyinslice1 []string
	var bothslice []string
	var onlyinslice2 []string

	for _, str := range slice1 {
		if !unique[str] {
			unique[str] = true
			onlyinslice1 = append(onlyinslice1, str)
		}
	}
	for _, str := range slice2 {
		if unique[str] {
			unique[str] = false
			bothslice = append(bothslice, str)
		} else {
			unique[str] = true
			onlyinslice2 = append(onlyinslice2, str)
		}
	}
	for str, existsinslice1 := range unique {
		{
			if existsinslice1 {
				onlyinslice1 = append(onlyinslice1, str)
			} else {
				onlyinslice2 = append(onlyinslice2, str)
			}
		}
	}

	return onlyinslice1, bothslice, onlyinslice2
}
func main() {
	slice1 := []string{"Dan", "Rob", "Alan", "James", "Bond"}
	slice2 := []string{"Alan", "David", "William", "Jack", "James"}
	onlyinslice1, bothslice, onlyinslice2 := slices(slice1, slice2)
	fmt.Println("Only in slice 1:", onlyinslice1)
	fmt.Println("In both slices:", bothslice)
	fmt.Println("Only in slice 2:", onlyinslice2)
}
