package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"asciiweb/functions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [PORT]")
		return
	}

	port := ":" + os.Args[1]

	http.HandleFunc("/", functions.Index)
	http.HandleFunc("/ascii-art", functions.Ascii)
	http.HandleFunc("/about", functions.About)
	http.HandleFunc("/how-it-works", functions.HowItWorks)

	staticDir := "./static/style/"
	staticURL := "/static/style/"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle(staticURL, http.StripPrefix(staticURL, fileServer))

	log.Printf("Server started at http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
