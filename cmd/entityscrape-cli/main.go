package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
