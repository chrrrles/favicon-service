package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/images/{fqdn}", GetFavicon).Methods("GET")

	log.Fatal(http.ListenAndServe(os.Getenv("FAVICON_LISTEN_ADDR"), router))
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fqdn := vars["fqdn"]

	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("FAVICON_REDIS_SERVER"),
	})

	imageBytes, err := rdb.Get(ctx, fqdn).Bytes()
	if err != nil {
		requestURL := fmt.Sprintf("http://localhost:8000/favicon/%s", fqdn)
		_, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Image not found")
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageBytes)
}
