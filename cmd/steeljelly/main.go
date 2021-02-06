package main

import (
	"log"

	"github.com/sapuri/steel-jelly/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
