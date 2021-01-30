package main

import (
	"log"

	"github.com/sapuri/steel-jelly/steeljelly/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
