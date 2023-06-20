package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the properties of the predefined Logger, including the log entry prefix
	// and a flag to disable printing the time, source line, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Create a slice of names
	names := []string{"Beyonce", "Taylor", "Eleni"}

	// Request a greeting message for the names
	messages, err := greetings.Hellos(names)

	// If an error is returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console
	fmt.Println(messages)
}
