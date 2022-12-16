package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func info(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("ECS_CONTAINER_METADATA_URI_V4")
	fmt.Fprintf(w, url)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "err:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/info", info)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
