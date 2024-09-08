package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	initRedis()

	http.Handle("/", RateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Request allowed"))
	})))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
