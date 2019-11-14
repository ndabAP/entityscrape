package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	srv "github.com/ndabAP/entityscrape/internal/entityscrape-srv"
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
	http.HandleFunc("/api/entities", srv.Entities)
	http.HandleFunc("/api/news", srv.News)

	log.Printf("starting server on port %s", port)
	http.ListenAndServe(":"+port, nil)
}
