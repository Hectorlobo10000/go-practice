package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined logger, including
	//the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	// a slice of names.
	names := []string{"Gladys", "Smantha", "Darrin"}

	// request gretting messages for the names
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned mp of 
	// messges to the console
	fmt.Println(messages)
}