package main

import (
	"fmt"
)

type Developer struct {
	work   string
	salary int
}

func GetDeveloper(work interface{}, salary interface{}) Developer {
	fmt.Println("IS IMPLEMENTED")
	var x Developer
	x.work = work.(string)
	x.salary = salary.(int)
	return x
}
func main() {
	fmt.Println("HELLO")

	var work interface{} = "Software Engineer"
	var salary interface{} = 300

	dynamicDev := GetDeveloper(work, salary)
	fmt.Println(dynamicDev.work)
	fmt.Println(dynamicDev.salary)
}
