package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello ")
	fmt.Fprintln(os.Stdout, "Hola world")
}
