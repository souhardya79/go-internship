package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	// TODO Implement
	unimap := make(map[string]bool)
	var uniname []string

	for _, dev := range developers {
		if _, exists := unimap[dev.Name]; !exists {
			//unimap[dev.Name] = true
			uniname = append(uniname, dev.Name)
		}
	}

	return uniname
}

func main() {
	developers := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}
	uniqueNames := FilterUnique(developers)
	fmt.Println("Unique names:\n", uniqueNames)
}
