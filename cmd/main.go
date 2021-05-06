package main

import (
	"fmt"

	database "github.com/zapscloud/go/database"
)

func main() {
	fmt.Println("Hello Message ", database.Hello("Demo Message"))
}
