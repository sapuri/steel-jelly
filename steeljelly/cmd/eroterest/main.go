package main

import (
	"log"

	"github.com/sapuri/steel-jelly/steeljelly/internal/eroterest"
)

const (
	linksCSV = "./output/eroterest_links.csv"
	blogsCSV = "./output/eroterest_blogs.csv"
)

func main() {
	// TODO: make CLI

	client := eroterest.NewClient()
	//if err := client.GetLinks(linksCSV, 1); err != nil {
	//	log.Fatal(err)
	//}

	if err := client.GetBlogs(linksCSV, blogsCSV); err != nil {
		log.Fatal(err)
	}
}
