package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
