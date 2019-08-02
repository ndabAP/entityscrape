package main

import (
	"log"

	cli "github.com/ndabAP/entityscrape/internal/entityscrape-cli"
)

func main() {
	err := cli.Do()
	if err != nil {
		log.Fatal(err)
	}
}
