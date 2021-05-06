package database

import (
	"fmt"
)

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome! GOLANG database Lib", name)
	return message
}
