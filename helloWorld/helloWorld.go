package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings human: ")
	log.SetFlags(0)

	fmt.Println("Hello, World")

	message, err := greetings.HelloThere("mr. potato head")
	fmt.Println(message)

	message, err = greetings.HelloThere("")

	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
}
