package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	port := ":4000"
	log.Println("Starting server on ", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving request: ", r.URL.Path)
	w.Write([]byte("Hello from Snippetbox"))
}
