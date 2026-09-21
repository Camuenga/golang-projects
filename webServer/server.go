package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Server failed...", http.StatusBadRequest)
	}
	fmt.Fprintf(w, " Server Root Conncted...", http.StatusOK)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", root)
	error := http.ListenAndServe(":8080", r)
	if error != nil {
		log.Fatal(error)
	}
}
