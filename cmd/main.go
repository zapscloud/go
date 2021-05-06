package main

import (
	"fmt"

	database "zapscloud/go/database"
)

func main() {
	fmt.Println("Hello Message ", database.Hello("Demo Message"))
}
