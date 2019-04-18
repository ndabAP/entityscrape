package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fs := http.FileServer(http.Dir("./website"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":"+port, nil)
}
