package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fs := http.FileServer(http.Dir("./website"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":"+port, nil)
}
