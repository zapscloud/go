package main

import (
	"fmt"

	database "github.com/zapscloud/golib/database"
)

func main() {
	fmt.Println("Hello Message ", database.Hello("GO Lib Sample Message"))
}
