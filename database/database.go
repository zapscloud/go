package database

import (
	"fmt"
)

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! Update on v2", name)
	return message
}
