package main

import (
	"log"
	"os"

	cli "github.com/ndabAP/entityscrape/internal/entityscrape-cli"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func main() {
	err := cli.Do(cli.AssocEntities{}, logger)
	if err != nil {
		log.Fatal(err)
	}
}
