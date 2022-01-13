package main

import (
	"fmt"
	"log"

 	"example.com/greetings"
)

func main() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Hector")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}