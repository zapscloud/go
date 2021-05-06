package main

import (
	"fmt"

	database "github.com/zapscloud/zapsdatabase-go"
)

func main() {
	fmt.Println("Hello Message ", database.Hello("Demo Message"))
}
