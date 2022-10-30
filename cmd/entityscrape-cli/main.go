package main

import (
	"log"
	"os"

	cli "github.com/ndabAP/entityscrape/internal/entityscrape-cli"
)

func main() {
	logger := log.New(os.Stdout, "[CLI] ", log.LstdFlags)

	err := cli.Do(cli.AssocEntities{}, logger)
	if err != nil {
		log.Fatal(err)
	}
}
