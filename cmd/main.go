package main

import (
	"fmt"

	database "zapscloud/database"
)

func main() {
	fmt.Println("Hello Message ", database.Hello("Demo Message"))
}
