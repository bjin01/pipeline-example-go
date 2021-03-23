package main

import (
	"fmt"
	"log"
	"net/http"
)

const webContent = "Bo is giving a peer review. Hope you like it!"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
