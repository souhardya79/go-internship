package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
	fmt.Print(nf)
}
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
