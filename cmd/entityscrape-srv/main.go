package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	srv "github.com/ndabAP/entityscrape/internal/entityscrape-srv"
)

func init() {
	// .env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("port must be set")
	}

	fs := http.FileServer(http.Dir("./website"))
	http.Handle("/", fs)

	http.HandleFunc("/api/entities", srv.Entities)
	http.HandleFunc("/api/news", srv.News)
	http.HandleFunc("/api/list", srv.List)
	http.HandleFunc("/api/associations", srv.Associations)

	log.Printf("starting server on port %s", port)
	http.ListenAndServe(":"+port, nil)
}
