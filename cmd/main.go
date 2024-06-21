package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.String("version", "1.0.0", "Version")
	flag.Parse()
	fmt.Println(*version)
	fmt.Println("Hello, World!")
}
