package main

import "os"

//go:generate go run scripts/includeadjectives.go

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
