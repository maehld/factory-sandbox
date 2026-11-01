package main

import (
	"flag"
	"fmt"

	"github.com/maehld/factory-sandbox/greeting"
)

func main() {
	name := flag.String("name", "", "who to greet")
	flag.Parse()
	fmt.Println(greeting.Greet(*name))
}
