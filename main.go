package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("ECS_CONTAINER_METADATA_URI_V4"))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/info", info)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
