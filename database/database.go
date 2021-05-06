package database

import (
	"fmt"
	"log"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	log.Println("Called Hellow")
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
