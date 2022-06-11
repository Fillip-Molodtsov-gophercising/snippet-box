package main

import (
	"log"
	"net/http"
)

const (
	htmlFolder   = "./ui/html/"
	staticFolder = "./ui/static/"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir(staticFolder))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		log.Fatal(err)
	}
}
