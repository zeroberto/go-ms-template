package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home!")
}

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
