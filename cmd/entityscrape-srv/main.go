package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	port = os.Getenv("PORT")
)

func init() {
	godotenv.Load()
}

func main() {
	if port == "" {
		log.Fatal("port must be set")
	}

	fs := http.FileServer(http.Dir("./website"))
	http.Handle("/", fs)

	http.ListenAndServe(":"+port, nil)
	log.Println("Listening...")
}
