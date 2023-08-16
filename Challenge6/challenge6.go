package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	// TODO Implement
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is empty")

	} else {
		lastindex := len(s.Items) - 1
		flight := s.Items[lastindex]
		s.Items = s.Items[:lastindex]
		return flight, nil
	}
}
func (s *Stack) Push(Flight Flight) {
	// TODO Implement
	s.Items = append(s.Items, Flight)
}

func (s *Stack) Peek() (Flight, error) {
	// TODO Implement
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is empty")

	} else {
		lastindex := len(s.Items) - 1
		flight := s.Items[lastindex]
		return flight, nil
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func main() {
	fmt.Println("Go Stack Implementation")
}
