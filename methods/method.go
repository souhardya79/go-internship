package main

import (
	"fmt"
	"math"
)

type circle struct { //structure by the name circle
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 { //calculating the area of circle
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("area", s.area())
}
func main() {
	c := circle{5}
	fmt.Println(c.area)
}
