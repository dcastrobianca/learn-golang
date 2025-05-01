package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	log.SetPrefix("hello: ") // the log entry prefix and a flag to disable printing
	log.SetFlags(0)          // the time, source file, and line number.

	//Request a greeting message
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
